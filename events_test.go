package golligator

import (
	"testing"

	"github.com/gogo/protobuf/proto"
)

func TestParseEvent(t *testing.T) {
	id := "1"
	user := "Jonh Snow"
	name := "Wall"
	protoMessage := createAddFrameworkEvent(id, user, name)
	bytes, _ := proto.Marshal(protoMessage)
	event, _ := parseEvent("AddFramework", bytes)
	addFramework := event.Message.(*AddFramework)
	if addFramework.GetFrameworkId().GetValue() != id {
		t.Fail()
	}
	if addFramework.GetFrameworkInfo().GetUser() != user {
		t.Fail()
	}
	if addFramework.GetFrameworkInfo().GetName() != name {
		t.Fail()
	}
}

func createAddFrameworkEvent(id string, user string, name string) proto.Message {
	frameworkId := &FrameworkID{Value: &id}
	frameworkInfo := &FrameworkInfo{User: &user, Name: &name}
	slaveResources := []*SlaveResources{}
	event := &AddFramework{FrameworkId: frameworkId, FrameworkInfo: frameworkInfo, SlaveResources: slaveResources}
	return event
}
