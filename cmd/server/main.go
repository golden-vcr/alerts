package main

import (
	"encoding/json"
	"os"

	"github.com/codingconcepts/env"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"

	eonscreen "github.com/golden-vcr/schemas/onscreen-events"
	"github.com/golden-vcr/server-common/entry"
	"github.com/golden-vcr/server-common/rmq"
	"github.com/golden-vcr/server-common/sse"
)

type Config struct {
	BindAddr   string `env:"BIND_ADDR"`
	ListenPort uint16 `env:"LISTEN_PORT" default:"5009"`

	RmqHost     string `env:"RMQ_HOST" required:"true"`
	RmqPort     int    `env:"RMQ_PORT" required:"true"`
	RmqVhost    string `env:"RMQ_VHOST" required:"true"`
	RmqUser     string `env:"RMQ_USER" required:"true"`
	RmqPassword string `env:"RMQ_PASSWORD" required:"true"`
}

func main() {
	app, ctx := entry.NewApplication("alerts")
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

	// Prepare a consumer and start receiving incoming messages from the onscreen-events
	// exchange: as we receive events, we'll simply fan them out to all connected SSE
	// clients
	onscreenEventsConsumer, err := rmq.NewConsumer(amqpConn, "onscreen-events")
	if err != nil {
		app.Fail("Failed to initialize AMQP consumer for onscreen-events", err)
	}
	onscreenEvents, err := onscreenEventsConsumer.Recv(ctx)
	if err != nil {
		app.Fail("Failed to init recv channel on onscreen-events consumer", err)
	}

	eventsChan := make(chan eonscreen.Event)
	go func() {
		done := false
		for !done {
			select {
			case <-ctx.Done():
				app.Log().Info("Consumer context canceled; exiting consumer loop")
				done = true
			case d, ok := <-onscreenEvents:
				if ok {
					var ev eonscreen.Event
					if err := json.Unmarshal(d.Body, &ev); err != nil {
						app.Fail("Failed to unmarshal message from onscreen-events", err)
					}
					eventsChan <- ev
				} else {
					app.Log().Info("Channel is closed; exiting consumer loop")
					done = true
				}
			}
		}
	}()

	// Start setting up our HTTP handlers, using gorilla/mux for routing
	r := mux.NewRouter()

	// Clients can use GET /events to get a stream of events that should be displayed
	// onscreen during broadcasts
	h := sse.NewHandler[eonscreen.Event](ctx, eventsChan)
	r.Path("/events").Methods("GET").Handler(h)

	// Handle incoming HTTP connections until our top-level context is canceled, at
	// which point shut down cleanly
	entry.RunServer(ctx, app.Log(), r, config.BindAddr, config.ListenPort)
}
