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

func RtCafeInfo(calltime time.Time)string{
	
	fg := 0
	file, err := ioutil.ReadFile("cafe")
	var datasets []Dataset
	json_err := json.Unmarshal(file, &datasets)
	if err != nil{
		fmt.Println("Format Error: ", json_err)
	}
	
	for k := range datasets{
		var timeformat = "2006-01-02"
		t,err := time.Parse(timeformat,datasets[k].ID)
		if err != nil{
			panic(err)
		}
		if t.Day() == calltime.Day(){
			/*b:= make([]byte, 0, 2048)
			b = append (b,datasets[k].Text...)
			b = append (b,datasets[k].Spa...)
			b = append (b,datasets[k].Fish...)
			b = append (b,datasets[k].Salad...)
			b = append (b,datasets[k].Dessert...)
			b = append (b,datasets[k].One...)
			b = append (b,datasets[k].Noodle...)
			b = append (b,datasets[k].Supper...)*/
			return datasets[k].Text
			fg += 1
		}
	}
	if fg == 0{
		return "err"
	}else{
		return "end"
	}

}
