package reqCafe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Dataset struct {
	ID string `json:"id"`
	Text string `json:"text"`
}

func RtCafeInfo(calltime time.Time)string{
	// Loading jsonfile
	fg := 0
	file, err := ioutil.ReadFile("./config.json")
	// 指定したDataset構造体が中身になるSliceで宣言する
	var datasets []Dataset
	json_err := json.Unmarshal(file, &datasets)
	if err!=nil{
		fmt.Println("Format Error: ", json_err)
	}

	for k:=range datasets{
		var timeformat = "2006-01-02"
		t,err := time.Parse(timeformat,datasets[k].ID)
		if err != nil{
			panic(err)
		}
		if t.Day() == calltime.Day(){
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
