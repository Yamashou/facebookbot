package reqCafe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	//"log"
)
const calltim = '
{
"id":2016-05-04,
"text":"tanakaarou"
}



'
type Dataset struct {
	ID string `json:"id"`
	Text string `json:"text"`
}

func RtCafeInfo(calltime time.Time)string{
	// Loading jsonfile
	fg := 0
	file, err := ioutil.ReadFile("./config.json")
	var datasets []Dataset
	json_err := json.Unmarshal(file, &datasets)
	if err != nil{
		fmt.Println("Format Error: ", json_err)
	}
	
	for k := range datasets{
		var timeformat = "2006-01-02"
		t,err := time.Parse(timeformat,calltim)
		if err != nil{
			//panic(err)
		}
		if t.Day() == calltime.Day(){
			return datasets[k].Text
			fg += 1
		}
	}
	if fg == 0{
		return "error"
	}else{
		return "end"
	}

}
