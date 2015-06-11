// Package golligator provides server and client for mesos protobuf events
package golligator

type Config struct {
	Port      int    // port to listen
	ServerUrl string // server url to send events
}

type Golligator struct {
	config      *Config
	inputEvents chan *Event
	client      *client
	server      *server
}

// NewGolligator creates new server with specified config
func NewGolligator(config *Config) *Golligator {
	g := &Golligator{config: config}
	g.inputEvents = make(chan *Event)
	g.client = newClient(config.ServerUrl)
	return g
}

// ListenAndServe creates new listener and returns channel of incoming events
func (g *Golligator) ListenAndServe() chan *Event {
	g.server = newServer(g.config.Port, g.inputEvents)
	return g.inputEvents
}

// Send event back to the server
func (g *Golligator) Send(event *Event) {
	g.client.sendEvent(event)
}
