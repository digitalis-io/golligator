package golligator

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

type Event struct {
	Type    string
	Message proto.Message
}

func parseEvent(eventType string, eventValue []byte) (*Event, error) {
	message, err := eventFactory(eventType)
	if err != nil {
		log.Println(err)
	}
	err = proto.Unmarshal(eventValue, message)
	if err != nil {
		log.Println(err)
	}
	return &Event{Type: eventType, Message: message}, err
}

func eventFactory(eventType string) (proto.Message, error) {
	var event proto.Message
	var err error
	switch eventType {
	case "AddFramework":
		event = &AddFramework{}
	case "RemoveFramework":
		event = &RemoveFramework{}
	case "ActivateFramework":
		event = &DeactivateFramework{}
	case "DeactivateFramework":
		event = &DeactivateFramework{}
	case "AddSlave":
		event = &AddSlave{}
	case "RemoveSlave":
		event = &RemoveSlave{}
	case "UpdateSlave":
		event = &UpdateSlave{}
	case "ActivateSlave":
		event = &ActivateSlave{}
	case "DeactivateSlave":
		event = &DeactivateSlave{}
	case "UpdateWhitelist":
		event = &UpdateWhitelist{}
	case "RequestResources":
		event = &RequestResources{}
	case "UpdateAllocation":
		event = &UpdateAllocation{}
	case "RecoverResources":
		event = &RecoverResources{}
	case "ReviveOffers":
		event = &ReviveOffers{}
	default:
		event = nil
		err = errors.New(fmt.Sprintf("Unexpected event type: %s", eventType))
	}
	return event, err
}
