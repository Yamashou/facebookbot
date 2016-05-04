package main

import (
	"github.com/facebookbot/reqCafe"
	"github.com/facebookbot/fbmessenger"
	"regexp"
	"time"
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
	foods.Judgment = []string{"kondate", "献立", "学食", "メニュー"}
	foods.Jf = false

	computers := new(DistributeMenu)
	computers.Judgment = []string{"演習室", "パソコン", "pc"}
	computers.Jf = false

	eves := new(DistributeMenu)
	eves.Judgment = []string{"hoge"}
	eves.Jf = false

	for i := 0; i < len(foods.Judgment); i++ {
		r := regexp.MustCompile(foods.Judgment[i])
		if r.MatchString(txt) {
			foods.Jf = true
		}
	}
	if foods.Jf {
		foods.Jf = false
		return "foods"
	} else {
		return txt
	}

	//for i:=0;i<len(Fncs);i++{
	//	if Fncs[i].Jf {
	//		r := regexp.MustCompile("*main")
	//		Fncs[i].Jf = false
	//		return r.ReplaceAllString(reflect.TypeOf(Fncs[i]),"")
	//	}
	//}
}

func getMessageText(receivedText string) string {
	if selectMenu(receivedText) == "foods" {
		//menu := reqCafe.RtCafeInfo(time.Now())
		/*b := make([]byte,0,1024)
		record := "\n"
		for _, line := range menu {
			b = append(b,line...)
			b = append(b,"\n")
		}
		m.Message.Text = string(b) */
		//log.Print(menu[0])
		return reqCafe.RtCafeInfo(time.Now())
	}
	return ""
}
