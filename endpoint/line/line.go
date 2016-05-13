package line

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// notification express whole JSON data
type notification struct {
	Result []*ReceiveEvent `json:"result"`
}

// ReceiveEvent 一つの投稿をあわらす
type ReceiveEvent struct {
	Content     *ReceiveContent `json:"content"`
	CreatedTime *JSONTime       `json:"createdTime"`
	EventType   string          `json:"eventType"`
	From        string          `json:"from"`
	FromChannel int             `json:"fromChannel"`
	ID          string          `json:"id"`
	To          []string        `json:"to"`
	ToChannel   int             `json:"toChannel"`
}

// ReceiveContent 中身
type ReceiveContent struct {
	ContentMetadata *ReceiveContentMetadata `json:"contentMetadata"`
	ContentType     int                     `json:"contentType"`
	CreatedTime     *JSONTime               `json:"createdTime"`
	DeliveredTime   int                     `json:"deliveredTime"`
	From            string                  `json:"from"`
	ID              string                  `json:"id"`
	Location        *Location               `json:"location"`
	Seq             string                  `json:"seq"`
	Text            string                  `json:"text"`
	To              []string                `json:"to"`
}

// ReceiveContentMetadata メタデータ
type ReceiveContentMetadata struct {
	RecvMode string `json:"AT_RECV_MODE"`
	Emtver   string `json:"EMTVER"`
}

// JSONTime parse用の時間
type JSONTime struct {
	time.Time
}

// SendEvent 実際に送信される構造体
type SendEvent struct {
	To        []string    `json:"to"`
	ToChannel int         `json:"toChannel"`
	EventType string      `json:"eventType"`
	Content   interface{} `json:"content"`
}

// SendContent 全てのコンテンツのベース
type SendContent struct {
	ContentType int `json:"contentType"`
	ToType      int `json:"toType"`
}

// SendTextContent テキスト送信系のコンテンツ
type SendTextContent struct {
	*SendContent
	Text string `json:"text"`
}

// sendImageContent イメージ送信用のコンテンツ
type sendImageContent struct {
	*SendContent
	OriginalContentURL string `json:"originalContentUrl"`
	PreviewImageURL    string `json:"previewImageUrl"`
}

// UnmarshalJSON JSONのタイムスタンプから変換する用
func (j *JSONTime) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	j.Time = time.Unix(i, 0)
	return nil
}

var _ json.Unmarshaler = (*JSONTime)(nil)

var channelID = os.Getenv("LINE_CHANNEL_ID")
var channelSecret = os.Getenv("LINE_CHANNEL_SECRET")
var mid = os.Getenv("LINE_MID")

// Location 位置情報
type Location struct {
	Title     string  `json:"title"`
	Address   string  `json:"address"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

var handleReceiveMessage func(ReceiveEvent)

const apiEndpoint string = "https://trialbot-api.line.me/v1/events"

// Listen webhook action from LINE server. You should pass callback function.
func Listen(callback func(ReceiveEvent)) {
	fmt.Println("Starting Listening LINE...")
	handleReceiveMessage = callback
	http.HandleFunc("/", webhookHandler)
	http.HandleFunc("/webhook", webhookHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// using go routine to return responce to LINE rapidly
		go bridgeForCallback(body)
	} else {
		fmt.Println(err)
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte("OK"))
}

func bridgeForCallback(body []byte) {
	var notification notification
	fmt.Println(string(body))
	err := json.Unmarshal(body, &notification)
	if err != nil {
		fmt.Println(err)
	}
	for _, receiveEvent := range notification.Result {
		handleReceiveMessage(*receiveEvent)
	}
}

// SendTextMessage 送信用ルーチン
func SendTextMessage(event *SendEvent) {
	go func(e *SendEvent) {
		// じゃあAPI叩いて送るべさ
		request(jsonEncode(e))
	}(event)
}

// 送信用のJSONにして返すよ
func jsonEncode(event *SendEvent) string {
	j, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(j)
}

// Request APIたたくーよ
func request(body string) error {
	fmt.Println(body)
	client := &http.Client{}
	//body := io.Reader
	req, err := http.NewRequest("POST", apiEndpoint, strings.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-type", "application/json; charset=UTF-8")
	req.Header.Add("X-Line-ChannelID", channelID)
	req.Header.Add("X-Line-ChannelSecret", channelSecret)
	req.Header.Add("X-Line-Trusted-User-With-ACL", mid)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if b, err := ioutil.ReadAll(resp.Body); err == nil {
		log.Print(string(b))
	} else {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return nil
}
