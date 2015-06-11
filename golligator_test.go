package golligator

import "testing"

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

func TestListenAndServe(t *testing.T) {
	config := &Config{Port: 8081, ServerUrl: "http://localhost:8081"}
	g := NewGolligator(config)
	events := g.ListenAndServe()
	if events == nil {
		t.Fail()
	}
}

func TestSend(t *testing.T) {
	config := &Config{Port: 8081, ServerUrl: "http://localhost:8081"}
	g := NewGolligator(config)
	message := createAddFrameworkEvent("1", "user", "framework")
	g.Send(&Event{Type: "AddFramework", Message: message})
}
