package main

import (
	"fmt"
	"os"

	"github.com/m2mtu/facebookbot/endpoint"
	"github.com/m2mtu/facebookbot/reply"
	"github.com/m2mtu/facebookbot/types"
)

func main() {
	os.Setenv("HTTP_PROXY", os.Getenv("FIXIE_URL"))
	os.Setenv("HTTPS_PROXY", os.Getenv("FIXIE_URL"))
	fmt.Println("starting...")
	endpoint.Listen(handleReceiveMessage)
}

func handleReceiveMessage(receivedEvent endpoint.Event) {
	sendEvent := endpoint.Event{
		SenderID:    receivedEvent.RecepientID,
		RecepientID: receivedEvent.SenderID,
	}
	reply.Talk(receivedEvent)
	endpoint.Send(sendEvent)
}

func createStaticState(receivedEvent endpoint.Event) types.StaticState {
	_state := types.StaticState{
		PossibleTopics:  []types.Topic{},
		OpponentID:      receivedEvent.SenderID,
		EndPointName:    endpoint.GetEndPointName(),
		ReceivedContent: []interface{}{receivedEvent.Content},
	}
	return _state
}
