package golligator

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	Events chan *Event
}

func (s *server) protoHandler(w http.ResponseWriter, r *http.Request) {
	var protoType string
	var protoValue []byte
	reader, err := r.MultipartReader()
	if err != nil {
		log.Println("Can't get reader")
	}
	for {
		part, err := reader.NextPart()
		if err != nil {
			break
		}
		partBytes, err := ioutil.ReadAll(part)
		if err != nil {
			log.Println(err)
			log.Println("Can't read bytes")
		}
		if part.FormName() == "type" {
			protoType = string(partBytes)
		}
		if part.FormName() == "value" {
			protoValue = partBytes
		}
	}
	event, err := parseEvent(protoType, protoValue)
	if err == nil {
		s.Events <- event
	}
}

func newServer(port int, events chan *Event) *server {
	server := &server{Events: events}
	http.HandleFunc("/", server.protoHandler)
	host := fmt.Sprintf(":%d", port)
	go http.ListenAndServe(host, nil)
	return server
}
