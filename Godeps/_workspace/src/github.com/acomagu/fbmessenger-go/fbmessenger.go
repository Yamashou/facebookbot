package fbmessenger

import (
	//"reflect"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
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
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}

func Listen(callback func(Messaging)) {
	handleReceiveMessage = callback
	http.HandleFunc("/", webhookHandler)
	http.HandleFunc("/webhook", webhookHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

var handleReceiveMessage func(Messaging)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
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
			if &event.Message != nil && event.Message.Text != "" {
				handleReceiveMessage(event)
			}
		}
		fmt.Fprintf(w, "Success")
	}
}

func SendTextMessage(recipient Recipient, sendText string) {
	m := new(SendMessage)
	m.Recipient = recipient

	log.Print("------------------------------------------------------------")
	log.Print(m.Message.Text)
	log.Print("------------------------------------------------------------")

	m.Message.Text = sendText

	log.Print(m.Message.Text)

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
