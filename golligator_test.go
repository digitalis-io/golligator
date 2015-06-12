package golligator

import (
	"os"
	"testing"
)

var events chan *Event
var g *Golligator

func TestMain(m *testing.M) {
	config := &Config{Port: 8081, ServerUrl: "http://localhost:8081"}
	g = NewGolligator(config)
	events = g.Listen()
	os.Exit(m.Run())
}

func TestNewGolligator(t *testing.T) {
	config := &Config{Port: 8081, ServerUrl: "http://localhost:8081"}
	g := NewGolligator(config)
	if g.config.Port != config.Port {
		t.Fail()
	}
	if g.client.serverUrl != config.ServerUrl {
		t.Fail()
	}
	if g.inputEvents == nil {
		t.Fail()
	}
}

func TestListenAndSend(t *testing.T) {
	message := createAddFrameworkEvent("1", "user", "framework")
	g.Send(&Event{Type: "AddFramework", Message: message})
	event := <-events
	if event.Message.(*AddFramework).GetFrameworkId().GetValue() != "1" {
		t.Fail()
	}
}

func BenchmarkSendAndReceive(b *testing.B) {
	message := createAddFrameworkEvent("1", "user", "framework")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Send(&Event{Type: "AddFramework", Message: message})
		<-events
	}
}
