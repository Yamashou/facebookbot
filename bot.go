package main

import (
	"fmt"
	"os"

	"github.com/m2mtu/facebookbot/endpoint"
	"github.com/m2mtu/facebookbot/reply"
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
