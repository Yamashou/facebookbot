package randomword

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Word express each word
type Word struct {
	Word string `json:"word"`
}

// RandomWord return word the user sent in past
func RandomWord(InWord string) string {
	fmt.Println("In RandomWord...")
	dbJSONPath := "./json/RWord.json"
	var T [1000]Word
	fmt.Println("set")
	rand.Seed(time.Now().UnixNano())
	file, err := ioutil.ReadFile(dbJSONPath)
	var datasets []Word
	jsonerr := json.Unmarshal(file, &datasets)
	if err != nil {
		fmt.Println("Format Error: ", jsonerr)
	}
	i := 0
	k := 0
	fg := 0
	for i = range datasets {
		if datasets[i].Word == InWord {
			fg++
			break
		}

	}
	for j := range datasets {
		T[j] = datasets[j]

		if datasets[j].Word == "" {
			break
		}
		k++
	}
	if fg != 1 {
		T[k].Word = InWord
		bytes, err := json.Marshal(T)
		if err != nil {
			fmt.Println(err)
		}
		ioutil.WriteFile(dbJSONPath, bytes, os.ModePerm)
	}

	r := rand.Intn(k)
	return datasets[r].Word
}
