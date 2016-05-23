package main

import (
	"fmt"
	"os"

	"github.com/m2mtu/facebookbot/reply"
	"github.com/m2mtu/facebookbot/talk"
)

func main() {
	os.Setenv("HTTP_PROXY", os.Getenv("FIXIE_URL"))
	os.Setenv("HTTPS_PROXY", os.Getenv("FIXIE_URL"))
	fmt.Println("starting...")
	talk.Listen(handleReceiveMessage)
}

func handleReceiveMessage(receivedEvent talk.Event) {
	sendEvent := talk.Event{
		SenderID:    receivedEvent.RecepientID,
		RecepientID: receivedEvent.SenderID,
	}
	reply.Talk(receivedEvent)
	talk.Send(sendEvent)
}
