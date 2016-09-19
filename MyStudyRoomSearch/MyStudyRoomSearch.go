package MyStudyRoomSearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type person struct {
	Std   string    `json:"std"`
	M     [6]string `json:"M"`
	Tu    [6]string `json:"Tu"`
	W     [6]string `json:"W"`
	T     [6]string `json:"T"`
	F     [6]string `json:"F"`
	ather string    `json:"ather"`
}

func RtRoom(menber string) [6]string {
	Mon := time.Date(2016, 5, 9, 0, 0, 0, 0, time.Local)
	Tus := time.Date(2016, 5, 10, 0, 0, 0, 0, time.Local)
	Wen := time.Date(2016, 5, 11, 0, 0, 0, 0, time.Local)
	Thu := time.Date(2016, 5, 12, 0, 0, 0, 0, time.Local)
	Fre := time.Date(2016, 5, 13, 0, 0, 0, 0, time.Local)

	file, err := ioutil.ReadFile("./json/room.json")
	var datasets []person
	json_err := json.Unmarshal(file, &datasets)
	if err != nil {
		fmt.Println("Format Error: ", json_err)
	}
	var T [6]string

	for k := range datasets {
		if datasets[k].Std == menber {
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

	for f := range T {
		if T[f] == "" {
			T[f] = "[あき]"
		}
	}
	return T
}
