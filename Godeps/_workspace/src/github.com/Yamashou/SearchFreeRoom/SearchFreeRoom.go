package SearchFreeRoom

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

func Serect(o int) [15]string {
	Mon := time.Date(2016, 5, 9, 0, 0, 0, 0, time.Local)
	Tus := time.Date(2016, 5, 10, 0, 0, 0, 0, time.Local)
	Wen := time.Date(2016, 5, 11, 0, 0, 0, 0, time.Local)
	Thu := time.Date(2016, 5, 12, 0, 0, 0, 0, time.Local)
	Fre := time.Date(2016, 5, 13, 0, 0, 0, 0, time.Local)

	file, err := ioutil.ReadFile("./json/room2.json")
	var datasets []person
	json_err := json.Unmarshal(file, &datasets)
	if err != nil {
		fmt.Println("Format Error: ", json_err)
	}
	var T [15]string
	Num := 0
	now := time.Now()
  o = o -1
	if now.Weekday() == Mon.Weekday() {
		for k := range datasets {
			if datasets[k].M[o] == "" {
				T[Num] = "( "+datasets[k].Std+" )"
				Num++

			}
		}
	} else if Tus.Weekday() == now.Weekday() {
		for k := range datasets {
			if datasets[k].Tu[o] == "" {
				T[Num] = "( "+datasets[k].Std+" )"
				Num++

			}
		}
	} else if Wen.Weekday() == now.Weekday() {
		for k := range datasets {
			if datasets[k].W[o] == "" {
				T[Num] = "( "+datasets[k].Std+" )"
				Num++

			}
		}
	} else if Thu.Weekday() == now.Weekday() {
		for k := range datasets {
			if datasets[k].T[o] == "" {
				T[Num] = "( "+datasets[k].Std+" )"
				Num++

			}
		}
	} else if Fre.Weekday() == now.Weekday() {
		for k := range datasets {
			if datasets[k].F[o] == "" {
				T[Num] = "( "+datasets[k].Std+" )"
				Num++

			}

		}
	}

	return T
}

