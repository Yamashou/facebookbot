package fbmessenger

import (
	//"reflect"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

var accessToken = os.Getenv("ACCESS_TOKEN")
var verifyToken = os.Getenv("VERIFY_TOKEN")

const (
	endPoint = "https://graph.facebook.com/v2.6/me/messages"
)

// receivedMessage ...
type receivedMessage struct {
	Object string  `json:"object"`
	Entry  []entry `json:"entry"`
}

// entry ...
type entry struct {
	ID        int64        `json:"id"`
	Time      int64        `json:"time"`
	Messagings []Messaging `json:"messaging"`
}

// Messaging ...
type Messaging struct {
	Sender    Sender    `json:"sender"`
	Recepient Recepient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   Message   `json:"message"`
}

// Sender ...
type Sender struct {
	ID int64 `json:"id"`
}

// Recepient ...
type Recepient struct {
	ID int64 `json:"id"`
}

// Message ...
type Message struct {
	MID  string `json:"mid"`
	Seq  int64  `json:"seq"`
	Text string `json:"text"`
}

// sendMessage ...
type sendMessage struct {
	Recepient Recepient `json:"recipient"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}

// Listen call callback function given when requested from Facebook service.
func Listen(callback func(Messaging)) {
	handleReceiveMessage = callback
	http.HandleFunc("/", webhookHandler)
	http.HandleFunc("/webhook", webhookHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

var handleReceiveMessage func(Messaging)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Query().Get("hub.verify_token") == verifyToken {
			fmt.Fprintf(w, r.URL.Query().Get("hub.challenge"))
		} else {
			fmt.Fprintf(w, "Error, wrong validation token")
		}
	}
	if r.Method == "POST" {
		var receivedMessage receivedMessage
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		if err = json.Unmarshal(b, &receivedMessage); err != nil {
			fmt.Println(err)
		}
		messagings := receivedMessage.Entry[0].Messagings
		for _, m := range messagings {
			if &m.Message != nil && m.Message.Text != "" {
				handleReceiveMessage(m)
			}
		}
		fmt.Fprintf(w, "Success")
	}
}

// SendTextMessage send message to Facebook endpoint.
func SendTextMessage(recepient Recepient, sendText string) {
	m := new(sendMessage)
	m.Recepient = recepient

	fmt.Println("------------------------------------------------------------")
	fmt.Println(m.Message.Text)
	fmt.Println("------------------------------------------------------------")

	m.Message.Text = sendText

	fmt.Println(m.Message.Text)

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	values := url.Values{}
	values.Add("access_token", accessToken)
	req.URL.RawQuery = values.Encode()
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	var result map[string]interface{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
