package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
	"github.com/Yamashou/MyClassSearch"
	"github.com/Yamashou/MyStudyRoomSearch"
	"github.com/acomagu/fbmessenger-go"
	"github.com/acomagu/linebot-go"
	"github.com/kurouw/infoSub"
	"github.com/kurouw/reqCafe"
	"github.com/Yamashou/RandomWord"
)

var endPointName = os.Getenv("ENDPOINT_NAME")

// DistributeMenu express functions of bot
type DistributeMenu struct {
	Judgment []string
	Jf       bool
}

func main() {
	os.Setenv("HTTP_PROXY", os.Getenv("FIXIE_URL"))
	os.Setenv("HTTPS_PROXY", os.Getenv("FIXIE_URL"))
	fmt.Println("starting...")
	if endPointName == "facebook" {
		fbmessenger.Listen(handleReceiveMessage)
	} else if endPointName == "line" {
		linebot.Listen(handleReceiveLINEMessage)
	}
}

func handleReceiveMessage(event fbmessenger.Messaging) {
	recipient := new(fbmessenger.Recipient)
	recipient.ID = event.Sender.ID
	fbmessenger.SendTextMessage(*recipient, getMessageText(event.Message.Text))
}

func handleReceiveLINEMessage(receiveEvent linebot.ReceiveEvent) {
	sendEvent := &linebot.SendEvent{}
	sendTextContent := &linebot.SendTextContent{SendContent: &linebot.SendContent{}}
	sendEvent.To = []string{receiveEvent.Content.SenderID}
	sendEvent.ToChannel = 1383378250
	sendEvent.EventType = "138311608800106203"
	sendTextContent.ContentType = 1
	sendTextContent.ToType = 1
	sendTextContent.Text = getMessageText(receiveEvent.Content.Text)
	sendEvent.Content = sendTextContent
	linebot.SendTextMessage(sendEvent)
}

func selectMenu(txt string) string {
	foods := new(DistributeMenu)
	foods.Judgment = []string{"kondate", "こんだて", "献立", "学食", "めにゅー", "メニュー"}
	foods.Jf = false

	tandai := new(DistributeMenu)
	tandai.Judgment = []string{"tandai", "短大", "たんだい"}
	tandai.Jf = false

	computers := new(DistributeMenu)
	computers.Judgment = []string{"演習室", "パソコン", "pc"}
	computers.Jf = false

	eves := new(DistributeMenu)
	eves.Judgment = []string{"hoge"}
	eves.Jf = false

	rooms := new(DistributeMenu)
	rooms.Judgment = []string{"std1", "std2", "std3", "std4", "std5", "std6", "hdw1", "hdw2", "hdw3", "hdw4", "CALL1", "CALL2", "iLab1", "iLab2"}
	rooms.Jf = false

	stringnames := []string{"foods", "tandai", "computers", "eves", "rooms"}
	allEvents := []DistributeMenu{*foods, *tandai, *computers, *eves, *rooms}

	for i := range allEvents {
		for j := 0; j < len(allEvents[i].Judgment); j++ {
			r := regexp.MustCompile(allEvents[i].Judgment[j])
			if r.MatchString(txt) {
				allEvents[i].Jf = true
			}
		}
	}
	flag := false
	for i := range allEvents {
		if allEvents[i].Jf {
			allEvents[i].Jf = false
			flag = true
			return stringnames[i]
		}
	}
	if !flag {
		cflag := false
		name := txt
		name = string([]rune(name)[:1])
		if name == "s" || name == "m" {
			cflag = true
			return "classes"
		}
		if !cflag {
			return "Subject!"
		}
	}
	return "notthing"
}

func getMessageText(receivedText string) string {
	dir, _ := os.Getwd()
	jsondir := dir + "/json/"
	selectRes := selectMenu(receivedText)
	if selectRes == "foods" {
		var res []string
		res = reqCafe.RtCafeInfo(jsondir)

		b := make([]byte, 0, 30)
		for v := 0; v < len(res); v++ {
			b = append(b, res[v]...)
			b = append(b, '\n')
		}
		return string(b)

	} else if selectRes == "tandai" {
		var res []string
		res = reqCafe.RtTnCafeInfo(jsondir)

		b := make([]byte, 0, 30)
		for v := 0; v < len(res); v++ {
			b = append(b, res[v]...)
			b = append(b, '\n')
		}
		return string(b)

	} else if selectRes == "rooms" {
		room := MyStudyRoomSearch.RtRoom(receivedText)
		b := make([]byte, 0, 30)
		for v := 0; v < len(room); v++ {
			b = append(b, strconv.Itoa(v+1)+"限: "...)
			b = append(b, room[v]...)
			b = append(b, '\n')
		}
		return string(b)
	}

	if selectRes == "Subject!" {
		return infoSub.ReturnSubInfo(receivedText)
	}

	if selectRes == "classes" {
		stdClass := MyClassSearch.RtClass(receivedText)

		b := make([]byte, 0, 30)
		for v := 0; v < len(stdClass); v++ {
			b = append(b, strconv.Itoa(v+1)+"限: "...)
			b = append(b, stdClass[v]...)
			b = append(b, '\n')
		}
		return string(b)

	}
	return RandomWord.ReturnWord(receivedText)
}
