package reqCafe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"log"
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

func RtCafeInfo(calltime time.Time)string{
	
	fg := 0
	file, err := ioutil.ReadFile("config.json")
	var datasets []Dataset
	log.Print(datasets)
	json_err := json.Unmarshal(file, &datasets)
	if err != nil{
		fmt.Println("Format Error: ", json_err)
	}
	
	for k := range datasets{
		var timeformat = "2006-01-02"
		t,err := time.Parse(timeformat,datasets[k].ID)
		log.Print(datasets)
		if err != nil{
			panic(err)
		}
		if t.Day() == calltime.Day(){
			//menu := []string{datasets[k].Text,datasets[k].Spa,datasets[k].Fish,datasets[k].Salad,datasets[k].Dessert,datasets[k].One,datasets[k].Noodle,datasets[k].Supper}
			log.Print(datasets)
			return datasets[k].Salad
			fg += 1
		}
	}
	//a := []string{"err","end"}
	if fg == 0{
		return "err"//a
	}else{
		return "end"//a
	}

}
