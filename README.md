# golligator

Serving and sending protobufs from/to mesos via HTTP.

## Usage

```go
package main

import "github.com/stealthly/golligator"

func main() {
  // Create config first, specify port to serve and url of remote server to respond
	config := &Config{Port: 8081, ServerUrl: "http://localhost:8081"}
  // Create Golligator from config
	g := NewGolligator(config)
  // Start server and get channel with incoming events
  events := g.ListenAndServe()
  for {
    event := <-events // Get event from channel
    // ...
    // Do actual work here
    // ...
    g.Send(event) // Send event back to the server
  }
}
```

