package main

import (
	"regexp"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	"math/rand"
)

var accessToken = os.Getenv("ACCESS_TOKEN")
var verifyToken = os.Getenv("VERIFY_TOKEN")

// const ...
const (
	EndPoint = "https://graph.facebook.com/v2.6/me/messages"
)

// ReceivedMessage ...
type ReceivedMessage struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

// Entry ...
type Entry struct {
	ID        int64       `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

// Messaging ...
type Messaging struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   Message   `json:"message"`
}

// Sender ...
type Sender struct {
	ID int64 `json:"id"`
}

// Recipient ...
type Recipient struct {
	ID int64 `json:"id"`
}

// Message ...
type Message struct {
	MID  string `json:"mid"`
	Seq  int64  `json:"seq"`
	Text string `json:"text"`
}

// SendMessage ...
type SendMessage struct {
	Recipient Recipient `json:"recipient"`
	Message   struct {
	        Text string `json:"text"`
	} `json:"message"`
}

func main() {
	http.HandleFunc("/", webhookHandler)
	http.HandleFunc("/webhook", webhookHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Facebook Bot")
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
	        if r.URL.Query().Get("hub.verify_token") == verifyToken {
			fmt.Fprintf(w, r.URL.Query().Get("hub.challenge"))
		} else {
			fmt.Fprintf(w, "Error, wrong validation token")
		}
	}
	if r.Method == "POST" {
	        var receivedMessage ReceivedMessage
	        b, err := ioutil.ReadAll(r.Body)
	        if err != nil {
			log.Print(err)
		}
	        if err = json.Unmarshal(b, &receivedMessage); err != nil {
			log.Print(err)
		}
	        messagingEvents := receivedMessage.Entry[0].Messaging
	        for _, event := range messagingEvents {
			senderID := event.Sender.ID
			if &event.Message != nil && event.Message.Text != "" {
				sentTextMessage(senderID, event.Message.Text)
			}
		}
		fmt.Fprintf(w, "Success")
	}
}

func rtFoods(rtext string) (f bool){
	foods := []string{"kondate","献立","学食","メニュー"}

	flag := false
	for i:=0;i<len(foods);i++ {
		r := regexp.MustCompile(foods[i])
		if r.MatchString(rtext){
			flag = true
			break
		}
		
	}
	return
}

func sentTextMessage(senderID int64, text string) {
	recipient := new(Recipient)
	recipient.ID = senderID
	m := new(SendMessage)
	m.Recipient = *recipient
	m.Message.Text = text
	
	log.Print("------------------------------------------------------------")
	log.Print(m.Message.Text)
	log.Print("------------------------------------------------------------")
	
	if rtFoods(m.Message.Text){
		m.Message.Text = "からあげカレー"
	}



	
	b, err := json.Marshal(m)
	if err != nil {
	        log.Print(err)
	}
	req, err := http.NewRequest("POST", EndPoint, bytes.NewBuffer(b))
	if err != nil {
	        log.Print(err)
	}
	values := url.Values{}
	values.Add("access_token", accessToken)
	req.URL.RawQuery = values.Encode()
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
	        log.Print(err)
	}
	defer res.Body.Close()
	var result map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	        log.Print(err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
	        log.Print(err)
	}
	log.Print(result)
}
