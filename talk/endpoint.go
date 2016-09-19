package talk

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Yamashou/facebookbot/talk/fbmessenger"
	"github.com/Yamashou/facebookbot/talk/line"
	"github.com/Yamashou/facebookbot/types"
)

// Event express one messaging
type Event struct {
	SenderID    types.UserID
	RecepientID types.UserID
	Content     types.EndPointContent
}

type texter interface {
	Text() string
}

var endPointName = os.Getenv("ENDPOINT_NAME")

func init() {
	if endPointName != "facebook" && endPointName != "line" {
		fmt.Println("Warning: ENDPOINT_NAME environment variable is not set or invalid. Please set as \"line\" or \"facebook\".")
		endPointName = "facebook"
	}
}

// GetEndPointName は今使われているチャットサービスの名前を返します。"line"か"facebook"です。
func GetEndPointName() string {
	return endPointName
}

// Listen 関数はコールバック関数を受け取り、チャットサービスからメッセージが届いた時に関数を呼び出します。
func Listen(receiveMessageCallback func(Event)) {
	switch endPointName {
	case "facebook":
		fbmessenger.Listen(handleReceiveFacebookMessage(receiveMessageCallback))
	case "line":
		line.Listen(handleReceiveLINEMessage(receiveMessageCallback))
	}
}

func handleReceiveFacebookMessage(handleReceiveMessage func(Event)) func(fbmessenger.Messaging) {
	return func(messaging fbmessenger.Messaging) {
		handleReceiveMessage(fbmessagingToEvent(messaging, fbmessenger.Recepient{0}))
	}
}

func handleReceiveLINEMessage(handleReceiveMessage func(Event)) func(line.ReceiveEvent) {
	return func(receiveEvent line.ReceiveEvent) {
		handleReceiveMessage(lineReceiveEventToEvent(receiveEvent))
	}
}

func fbmessagingToEvent(_messaging fbmessenger.Messaging, _recepient fbmessenger.Recepient) Event {
	e := Event{}
	e.SenderID = types.UserID(strconv.FormatInt(_messaging.Sender.ID, 10))
	e.RecepientID = types.UserID(strconv.FormatInt(_recepient.ID, 10))
	tc := TextContent{}
	tc.SetText(_messaging.Message.Text)
	e.Content = tc
	return e
}

func lineReceiveEventToEvent(receiveEvent line.ReceiveEvent) Event {
	e := Event{}
	e.SenderID = types.UserID(receiveEvent.Content.From)
	e.RecepientID = types.UserID(receiveEvent.To[0])
	tc := TextContent{}
	tc.SetText(receiveEvent.Content.Text)
	e.Content = tc
	return e
}

// SendText はテキストと受信者IDを受け取り、メッセージを送ります。Sendメソッドのラッパー関数です。
func SendText(text string, recepientID types.UserID) error {
	event := Event{}
	event.RecepientID = recepientID
	content := TextContent{}
	content.SetText(text)
	event.Content = content
	return Send(event)
}

// Send はEvent型の構造体を受け取り、そのEvent通りにメッセージを送信します。
func Send(event Event) error {
	switch endPointName {
	case "facebook":
		switch content := event.Content.(type) {
		case texter:
			fmt.Println("<<<", content.Text())
			intRecepientID, err := strconv.ParseInt(event.RecepientID.String(), 10, 64)
			if err != nil {
				return errors.New("cannot parse RecepientID to int64")
			}
			fbmessenger.SendTextMessage(fbmessenger.Recepient{intRecepientID}, content.Text())
		default:
			return errors.New("Event.Content type is invalid")
		}
	case "line":
		switch content := event.Content.(type) {
		case texter:
			fmt.Println(content.Text)
			sendEvent := &line.SendEvent{}
			sendTextContent := &line.SendTextContent{SendContent: &line.SendContent{}}
			sendEvent.To = []string{event.RecepientID.String()}
			sendEvent.ToChannel = 1383378250
			sendEvent.EventType = "138311608800106203"
			sendTextContent.ContentType = 1
			sendTextContent.ToType = 1
			sendTextContent.Text = content.Text()
			sendEvent.Content = sendTextContent
			line.SendTextMessage(sendEvent)
		default:
			return errors.New("Event.Content type is invalid")
		}
	}
	return nil
}
