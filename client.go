package golligator

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gogo/protobuf/proto"
)

type client struct {
	serverUrl string
}

func newClient(serverUrl string) *client {
	return &client{serverUrl: serverUrl}
}

func (c *client) sendRequest(eventType string, data []byte) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	typeField, _ := writer.CreateFormField("type")
	typeField.Write([]byte(eventType))
	valueField, _ := writer.CreateFormField("value")
	valueField.Write(data)
	writer.Close()
	req, _ := http.NewRequest("POST", c.serverUrl, &buffer)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, _ := client.Do(req)
	resp.Body.Close()
}

func (c *client) sendEvent(event *Event) {
	go func() {
		bytes, err := proto.Marshal(event.Message)
		if err != nil {
			log.Println(err)
		} else {
			c.sendRequest(event.Type, bytes)
		}
	}()
}
