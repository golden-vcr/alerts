package main

import (
	"encoding/json"
	"os"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/errgroup"

	ebroadcast "github.com/golden-vcr/schemas/broadcast-events"
	eonscreen "github.com/golden-vcr/schemas/onscreen-events"
	"github.com/golden-vcr/server-common/entry"
	"github.com/golden-vcr/server-common/rmq"
)

type Config struct {
	RmqHost     string `env:"RMQ_HOST" required:"true"`
	RmqPort     int    `env:"RMQ_PORT" required:"true"`
	RmqVhost    string `env:"RMQ_VHOST" required:"true"`
	RmqUser     string `env:"RMQ_USER" required:"true"`
	RmqPassword string `env:"RMQ_PASSWORD" required:"true"`
}

func main() {
	app, ctx := entry.NewApplication("alerts-consumer")
	defer app.Stop()

	// Parse config from environment variables
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		app.Fail("Failed to load .env file", err)
	}
	config := Config{}
	if err := env.Set(&config); err != nil {
		app.Fail("Failed to load config", err)
	}

	// Initialize an AMQP client
	amqpConn, err := amqp.Dial(rmq.FormatConnectionString(config.RmqHost, config.RmqPort, config.RmqVhost, config.RmqUser, config.RmqPassword))
	if err != nil {
		app.Fail("Failed to connect to AMQP server", err)
	}
	defer amqpConn.Close()

	// Prepare a producer that we can use to send messages to the onscreen-events queue,
	// whenenver we determine that a new alert needs to be displayed
	onscreenEventsProducer, err := rmq.NewProducer(amqpConn, "onscreen-events")
	if err != nil {
		app.Fail("Failed to initialize AMQP producer for onscreen-events", err)
	}

	// Prepare a consumer and start receiving incoming messages from the
	// broadcast-events exchange, so we can produce onscreen events whenenver we start
	// screening a new tape
	broadcastEventsConsumer, err := rmq.NewConsumer(amqpConn, "broadcast-events")
	if err != nil {
		app.Fail("Failed to initialize AMQP consumer for broadcast-events", err)
	}
	broadcastEvents, err := broadcastEventsConsumer.Recv(ctx)
	if err != nil {
		app.Fail("Failed to init recv channel on broadcast-events consumer", err)
	}

	// Each time we read a message from the queue, spin up a new goroutine for that
	// message, parse it according to our broadcast-events schema, then handle it
	wg, ctx := errgroup.WithContext(ctx)
	done := false
	for !done {
		select {
		case <-ctx.Done():
			app.Log().Info("Consumer context canceled; exiting main loop")
			done = true
		case d, ok := <-broadcastEvents:
			if ok {
				wg.Go(func() error {
					var ev ebroadcast.Event
					if err := json.Unmarshal(d.Body, &ev); err != nil {
						return err
					}
					logger := app.Log().With("broadcastEvent", ev)
					logger.Info("Consumed from broadcast-events")

					// If we've just started screening a new tape, generate an onscreen
					// event to display the details of that tape
					if ev.Type == ebroadcast.EventTypeScreeningStarted {
						data, err := json.Marshal(eonscreen.Event{
							Type: eonscreen.EventTypeStatus,
							Payload: eonscreen.Payload{
								Status: &eonscreen.PayloadStatus{
									CurrentTapeId: ev.Screening.TapeId,
								},
							},
						})
						if err != nil {
							return err
						}
						return onscreenEventsProducer.Send(ctx, data)
					}
					return nil
				})
			} else {
				app.Log().Info("Channel is closed; exiting main loop")
				done = true
			}
		}
	}

	if err := wg.Wait(); err != nil {
		app.Fail("Encountered an error during message handling", err)
	}
}
