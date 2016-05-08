package reqCafe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	//"log"
)

type Dataset struct {
	ID string `json:"id"`
	Text string `json:"text"`
	Don string `json:"don"`
	Spa string `json:"spaghetti"`
	Fish string `json:"fish"`
	Salad string `json:"salad"`
	Dessert string `json:"dessert"`
	One string `json:"one"`
	Noodle string `json:"noodle"`
	Supper string `json:"supper"`
}

type TDataset struct {
	ID string `json:"id"`
	Text string `json:"text"`
	Don string `json:"don"`
	Salad string `json:"salad"`
}

func RtCafeInfo(calltime time.Time)[]string{
	fg := 0
	file, err := ioutil.ReadFile("json/config.json")
	var datasets []Dataset
	json_err := json.Unmarshal(file, &datasets)
	if err != nil{
		fmt.Println("Format Error: ", json_err)
	}

	for k := range datasets {
		var timeformat = "2006-01-02"
		t, err := time.Parse(timeformat,datasets[k].ID)
		if err != nil {
			panic(err)
		}
		if t.Day() == calltime.Day() {
			menues := []string{"日替わり　　: "+datasets[k].Text,"丼物　　　　: "+datasets[k].Don,"スパゲッティ: "+datasets[k].Spa,"魚　　　　　: "+datasets[k].Fish,"サラダ　　　: "+datasets[k].Salad,"デザート　　: "+datasets[k].Dessert,"単品　　　　: "+datasets[k].One,"麺類　　　　: "+datasets[k].Noodle,"夕飯　　　　: "+datasets[k].Supper}
			return menues
			fg += 1
		}
	}
	a := []string{"err","end"}
	if fg == 0 {
		return a
	}else{
		return a
	}
}


func RtTnCafeInfo(calltime time.Time)[]string{
	fg := 0
	file, err := ioutil.ReadFile("json/tandai2.json")
	var datasets []TDataset
	json_err := json.Unmarshal(file, &datasets)
	if err != nil{
		fmt.Println("Format Error: ", json_err)
	}

	for k := range datasets {
		var timeformat = "2006-01-02"
		t, err := time.Parse(timeformat,datasets[k].ID)
		if err != nil {
			panic(err)
		}
		if t.Day() == calltime.Day() {
			menues := []string{"日替わり　　: "+datasets[k].Text,"丼物 　　　　: "+datasets[k].Don,"サラダ　　　: "+datasets[k].Salad}
			return menues
			fg += 1
		}
	}
	a := []string{"err","end"}
	if fg == 0 {
		return a
	}else{
		return a
	}
}
