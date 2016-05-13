package main

import (
	"fmt"
	"os"

	"github.com/m2mtu/facebookbot/endpoint"
	"github.com/m2mtu/facebookbot/reply"
	"github.com/m2mtu/facebookbot/state"
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
	_state := createState(receivedEvent)
	sendEvent.Content = reply.Get(_state)
	endpoint.Send(sendEvent)
}

func createState(receivedEvent endpoint.Event) state.State {
	_state := state.State{
		PossibleTopics:   []state.Topic{state.Topic{ID: 0}},
		OpponentID:       receivedEvent.SenderID,
		EndPointName:     endpoint.GetEndPointName(),
		ReceivedContents: []interface{}{receivedEvent.Content},
	}
	return _state
}
