package main

import (
	"github.com/facebookbot/reqCafe"
	"github.com/facebookbot/fbmessenger"
	"regexp"
	"time"
//	"bytes"
)


type DistributeMenu struct {
	Judgment []string
	Jf       bool
}

func main() {
	fbmessenger.Listen(handleRecieveMessage)
}

func handleRecieveMessage(event fbmessenger.Messaging) {
	recipient := new(fbmessenger.Recipient)
	recipient.ID = event.Sender.ID
	fbmessenger.SendTextMessage(*recipient, getMessageText(event.Message.Text))
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
	
	for i := range allEvents {
		if allEvents[i].Jf {
			allEvents[i].Jf = false
			return stringnames[i]
		}
	}
	return "notthing"
}

func getMessageText(receivedText string) string {
	if selectMenu(receivedText) == "foods" {
 		var res []string
		res = reqCafe.RtCafeInfo(time.Now())

		/*
		var b []byte
		for v := range res {
			b = append(b,v...)
			b = append(b,'\n')
		}*/
		return res[0]+res[1]
		
	}else if selectMenu(receivedText) == "tandai"{
		var res []string
		res = reqCafe.RtTnCafeInfo(time.Now())

		/*var b []byte
		for v := range res {
			b = append(b,v...)
			b = append(b,'\n')
		}*/
		return res[0]+res[1]
		
	}
	
	return receivedText
}
