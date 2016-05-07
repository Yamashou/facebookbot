package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Word struct {
	Word string `json:"word"`
}

func randam_word(InWord string){
  var T [1000]Word
  fmt.Println("set")
	rand.Seed(time.Now().UnixNano())
	file, err := ioutil.ReadFile("./RWord.json")
	var datasets []Word
	json_err := json.Unmarshal(file, &datasets)
	if err != nil {
		fmt.Println("Format Error: ", json_err)
	}
	i := 0
  k := 0
  fg := 0
	for i = range datasets {
		if datasets[i].Word == InWord {
      fg +=1
			break
		}

	}
  for j := range datasets{
    T[j] = datasets[j]

    if datasets[j].Word ==""{
      break
    }
    k++
  }

	if fg != 1 {
		T[k].Word = InWord
		bytes, _ := json.Marshal(T)
		ioutil.WriteFile("./RWord.json", bytes, os.ModePerm)
	} else {
    ky := rand.Intn(10)
    fmt.Println(ky)
	fmt.Println(datasets[ky].Word)
	}
}

func main() {
	randam_word("米うま〜〜〜〜〜〜")
}
