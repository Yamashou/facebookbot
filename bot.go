package main

import (
	"github.com/kurouw/reqCafe"
	"github.com/acomagu/fbmessenger-go"
	"github.com/kurouw/infoSub"
	"github.com/Yamashou/MyClassSearch"
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
	foods.Judgment = []string{"kondate","こんだて","献立", "学食","めにゅー", "メニュー"}
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

	
	name := txt
	name = string([]rune(name)[:1])
	if name == "s" || name  == "m" {
	}	
	
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
		for v := 0;v < len(res) ; v++{
			b = append(b,res[v]...)
			b = append(b,'\n')
		}
		return string(b)
		
	}else if selectRes == "tandai"{
		var res []string
		res = reqCafe.RtTnCafeInfo(time.Now())

		b := make([]byte,0,30)
		for v := 0;v < len(res) ; v++{
			b = append(b,res[v]...)
			b = append(b,'\n')
		}
		return string(b)
		
	}

	if selectRes == "Subject!" {
		return infoSub.ReturnSubInfo(receivedText)
	}

	if selectRes == "classes" {
		stdClass := MyClassSearch.RtClass(receivedText)

		b := make([]byte,0,100)
		for v := 0;v < len(stdClass) ; v++ {
			b = append(b,stdClass[v]...)
			b = append(b,'\n')
		}
		return string(b)
		
	}

	
	return receivedText
}
