# alerts

The **alerts** service is responsible for serving real-time data about alerts that need
to appear onscreen during streams. The primary client application for the alerts server
is [**graphics**](https://github.com/golden-vcr/graphics) frontend which runs in an OBS
browser source.

At present, this service consists of a single, very simple HTTP server application. It
maintains a RabbitMQ consumer which reads incoming events from the
[**onscreen-events**](https://github.com/golden-vcr/schemas/tree/main/onscreen-events)
queue. Meanwhile, it accepts HTTP client connections via `GET /events`, opening an SSE
stream for each connected client. Each time an event is received from the queue, it is
simply fanned out to all connected clients.
