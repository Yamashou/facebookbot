package randomword

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// Word express each word
type Word struct {
	Word string `json:"word"`
}

// RandomWord return word the user sent in past
func RandomWord(InWord string) string {
	var T [1000]Word
	fmt.Println("set")
	rand.Seed(time.Now().UnixNano())
	file, err := ioutil.ReadFile("./json/RWord.json")
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

	r := rand.Intn(k)
	return datasets[r].Word
}
