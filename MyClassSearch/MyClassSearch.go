package MyClassSearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type namegetter struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type person struct {
	No    string    `json:"No"`
	M     [6]string `json:"M"`
	Tu    [6]string `json:"Tu"`
	W     [6]string `json:"W"`
	T     [6]string `json:"T"`
	F     [6]string `json:"F"`
	Ather string    `json:"ather"`
}

//Rtclass ...
func RtClass(menber string) [6]string {
	Mon := time.Date(2016, 5, 9, 0, 0, 0, 0, time.Local)
	Tus := time.Date(2016, 5, 10, 0, 0, 0, 0, time.Local)
	Wen := time.Date(2016, 5, 11, 0, 0, 0, 0, time.Local)
	Thu := time.Date(2016, 5, 12, 0, 0, 0, 0, time.Local)
	Fre := time.Date(2016, 5, 13, 0, 0, 0, 0, time.Local)

	file, err := ioutil.ReadFile("./json/subjects2.json")
	var datasets []person
	json_err := json.Unmarshal(file, &datasets)
	if err != nil {
		fmt.Println("Format Error: ", json_err)
	}
	var T [6]string

	for k := range datasets {
		if datasets[k].No == menber {
			now := time.Now()
			if now.Weekday() == Mon.Weekday() {
				T = datasets[k].M
				break
			} else if Tus.Weekday() == now.Weekday() {
				T = datasets[k].Tu
				break
			} else if Wen.Weekday() == now.Weekday() {
				T = datasets[k].W
				break
			} else if Thu.Weekday() == now.Weekday() {
				T = datasets[k].T
				break
			} else if Fre.Weekday() == now.Weekday() {
				T = datasets[k].F
				break
			} else {
				break
			}
		}
	}
	T = chName(T)
	for f := range T {
		if T[f] == "" {
			T[f] = "[あき]"
		}
	}
	return T
}
func chName(code [6]string) [6]string {
	file, err := ioutil.ReadFile("./json/subjects.json")
	var datasets []namegetter
	json_err := json.Unmarshal(file, &datasets)
	if err != nil {
		log.Print("Format Error: ", json_err)
	}

	for l := range datasets {
		for i := 0; i < 6; i++ {
			if code[i] == datasets[l].Code {
				code[i] = datasets[l].Name
			}
		}
	}
	fmt.Println(code)
	return code
}

func getMyClassSerch(chatroom chan string) {
	text := <-chatroom
	if (text[0] == 's') || (text[0] == 'm') {
		m := RtClass(text)
		t := strings.Join(m[:], "\n")
		chatroom <- t
	}
}
