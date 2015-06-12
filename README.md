# golligator

Serving and sending protobufs from/to mesos via HTTP.

## Usage

```go
package main

import "github.com/stealthly/golligator"

func main() {
  // Create config first, specify port to serve and url of remote server to respond
  config := &golligator.Config{Port: 8081, ServerUrl: "http://remote.server.com"}
  // Create Golligator from config
  g := golligator.NewGolligator(config)
  // Start server and get channel with incoming events
  events := g.Listen()
  for {
    event := <-events // Get event from channel
    // ...
    // Do actual work here
    // ...
    g.Send(event) // Send event back to the server
  }
}
```

## Reference

#### type Config

```go
type Config struct {
	Port      int    // port to listen
	ServerUrl string // server url to send events
}
```


#### type Event

```go
type Event struct {
	Type    string
	Message proto.Message
}
```


#### type Golligator

```go
type Golligator struct {
}
```


#### func  NewGolligator

```go
func NewGolligator(config *Config) *Golligator
```
NewGolligator creates new server with specified config

#### func (*Golligator) Listen

```go
func (g *Golligator) Listen() chan *Event
```
Listen creates new http listener and returns channel of incoming events

#### func (*Golligator) Send

```go
func (g *Golligator) Send(event *Event)
```
Send event back to the server
