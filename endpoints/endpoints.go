package endpoints

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/acomagu/fbmessenger-go"
	"github.com/m2mtu/facebookbot/line"
)

// Event express one messaging
type Event struct {
	SenderID    string
	RecepientID string
	Content     interface{}
}

// TextContent express content of one message
type TextContent struct {
	Text string
}

var endPointName = os.Getenv("ENDPOINT_NAME")
var handleReceiveMessage func(Event)

func init() {
	if endPointName != "facebook" && endPointName != "line" {
		fmt.Println("Warning: ENDPOINT_NAME environment variable is not set or invalid. Please set as \"line\" or \"facebook\".")
		endPointName = "facebook"
	}
}

// Listen start listening at the endpoint. The arguments must be callback function.
func Listen(receiveMessageCallback func(Event)) {
	handleReceiveMessage = receiveMessageCallback
	switch endPointName {
	case "facebook":
		fbmessenger.Listen(handleReceiveFacebookMessage)
	case "line":
		line.Listen(handleReceiveLINEMessage)
	}
}

func handleReceiveFacebookMessage(messaging fbmessenger.Messaging) {
	handleReceiveMessage(fbmessagingToEvent(messaging, fbmessenger.Recepient{0}))
}

func handleReceiveLINEMessage(receiveEvent line.ReceiveEvent) {
	handleReceiveMessage(lineReceiveEventToEvent(receiveEvent))
}

func fbmessagingToEvent(_messaging fbmessenger.Messaging, _recepient fbmessenger.Recepient) Event {
	e := Event{}
	e.SenderID = strconv.FormatInt(_messaging.Sender.ID, 10)
	e.RecepientID = strconv.FormatInt(_recepient.ID, 10)
	e.Content = TextContent{_messaging.Message.Text}
	return e
}

func lineReceiveEventToEvent(receiveEvent line.ReceiveEvent) Event {
	e := Event{}
	e.SenderID = receiveEvent.From
	e.RecepientID = receiveEvent.To[0]
	e.Content = TextContent{receiveEvent.Content.Text}
	return e
}

// Send method send messaging
func Send(event Event) error {
	switch endPointName {
	case "facebook":
		switch content := event.Content.(type) {
		case TextContent:
			intRecepientID, err := strconv.ParseInt(event.RecepientID, 10, 64)
			if err != nil {
				return errors.New("cannot parse RecepientID to int64")
			}
			fbmessenger.SendTextMessage(fbmessenger.Recepient{intRecepientID}, content.Text)
		default:
			return errors.New("Event.Content type is invalid")
		}
	case "line":
		switch content := event.Content.(type) {
		case TextContent:
			sendEvent := &line.SendEvent{}
			sendTextContent := &line.SendTextContent{SendContent: &line.SendContent{}}
			sendEvent.To = []string{event.RecepientID}
			sendEvent.ToChannel = 1383378250
			sendEvent.EventType = "138311608800106203"
			sendTextContent.ContentType = 1
			sendTextContent.ToType = 1
			sendTextContent.Text = content.Text
			sendEvent.Content = sendTextContent
			line.SendTextMessage(sendEvent)
		default:
			return errors.New("Event.Content type is invalid")
		}
	}
	return nil
}
