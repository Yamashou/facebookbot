package main

import (
	"github.com/kurouw/reqCafe"
	"github.com/acomagu/fbmessenger-go"
	"github.com/kurouw/infoSub"
	"github.com/acomagu/linebot-go"
	"regexp"
	"time"
	"fmt"
	"os"
//	"bytes"
)

var endPointName = "line"

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
		fbmessenger.Listen(handleReceiveFacebookMessage)
	} else if endPointName == "line" {
		linebot.Listen(handleReceiveLINEMessage)
	}
}

func handleReceiveFacebookMessage(event fbmessenger.Messaging) {
	recipient := new(fbmessenger.Recipient)
	recipient.ID = event.Sender.ID
	fbmessenger.SendTextMessage(*recipient, getMessageText(event.Message.Text))
}

func handleReceiveLINEMessage(receiveEvent linebot.ReceiveEvent) {
	sendEvent := &linebot.SendEvent{}
	sendTextContent := &linebot.SendTextContent{}
	sendEvent.To = []string{receiveEvent.From}
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
	foods.Judgment = []string{"kondate","こんだて","献立", "学食","めにゅー", "メニュー","menu"}
	foods.Jf = false

	tandai := new(DistributeMenu)
	tandai.Judgment = []string{"tandai","短大","たんだい"}
	tandai.Jf = false
	
	computers := new(DistributeMenu)
	computers.Judgment = []string{"演習室", "パソコン", "pc"}
	computers.Jf = false

	eves := new(DistributeMenu)
	eves.Judgment = []string{"hoge"}
	eves.Jf = false

	stringnames := []string{"foods","tandai","computers","eves"}
	allEvents := []DistributeMenu{*foods,*tandai,*computers,*eves}

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
		return "Subject!"
	}
	return "notthing"
}

func getMessageText(receivedText string) string {
	selectRes := selectMenu(receivedText)
	if selectRes == "foods" {
		var res []string
		res = reqCafe.RtCafeInfo(time.Now())

	
		b := make([]byte,0,30)
		for v := 0;v < len(res) ; v ++{
			b = append(b,res[v]...)
			b = append(b,'\n')
		}
		return string(b)
		
	}else if selectRes == "tandai"{
		var res []string
		res = reqCafe.RtTnCafeInfo(time.Now())

		b := make([]byte,0,30)
		for v := 0;v < len(res) ; v ++{
			b = append(b,res[v]...)
			b = append(b,'\n')
		}
		return string(b)
		
	}

	if selectRes == "Subject!" {
		return infoSub.ReturnSubInfo(receivedText)
	}
	
	return receivedText
}
