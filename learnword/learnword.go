package learnword

import (
	"fmt"
	cabocha "github.com/ledyba/go-cabocha"
	"github.com/m2mtu/facebookbot/types"
	"math/rand"
	"github.com/m2mtu/facebookbot/endpoint"
)

// TempState .
type TempState struct {}

// IsProper returns the judgment should endter this topic.
func IsProper(static types.StaticState) bool {
	return rand.Int63n(2) == 1
}

// Talk method talk with user.
func Talk(static types.StaticState, temp TempState, perm types.PermState) (TempState, types.PermState, bool) {
	if content, ok := static.ReceivedContent.(endpoint.TextContent); ok {
		_cabocha := cabocha.MakeCabocha()
		sentence, err := _cabocha.Parse(content.Text)
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			for _, chunk := range sentence.Chunks {
				for _, tok := range chunk.Tokens {
					endpoint.SendText(tok.Body + ": " + tok.Features[0], static.OpponentID)
					if tok.Features[0] == "名詞" {}
				}
			}
		}
	}
	return temp, perm, false
}
