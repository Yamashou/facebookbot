package shiritori

import (
	"fmt"
	"regexp"
	"github.com/m2mtu/facebookbot/endpoint"
	"github.com/m2mtu/facebookbot/types"
)

// TempState of this topic
type TempState struct {
	LastWord string
}

// PermState of this topic
type PermState struct {}

// Shiritori is topic module of Shiritori
type Shiritori struct {}

// IsProper defines when this topic should be called.
func (p Shiritori) IsProper(content types.EndPointContent) bool {
	fmt.Println("IsProper")
	switch _content := content.(type) {
	case endpoint.TextContent:
		return regexp.MustCompile(`しりとり`).MatchString(_content.Text)
	default:
		return false
	}
}

// InitialTempState returns initial object of TempState
func (p Shiritori) InitialTempState() types.TempState {
	return TempState{LastWord: "しりとり"}
}

// InitialPermState returns initial object of PermState
func (p Shiritori) InitialPermState() types.PermState {
	return PermState{}
}

// Talk method define the logic of Shiritori topic.
func (p Shiritori) Talk(static types.StaticState, temp types.TempState, perm types.PermState) (types.TempState, types.PermState) {
	fmt.Println("shiritori!")
	textContent, ok1 := static.ReceivedContent.(endpoint.TextContent)
	_temp, ok2 := temp.(TempState)
	if ok1 && ok2 {
		text := textContent.Text
		textRunes := []rune(text)
		endpoint.SendText(_temp.LastWord, static.OpponentID)
		if answer, isFound := wordsStartsWith(textRunes[len(textRunes) - 1]); isFound {
			endpoint.SendText(string(answer), static.OpponentID)
			_temp.LastWord = string(answer)
		}
		return types.TempState(_temp), perm
	}
	return temp, perm
}

func wordsStartsWith(firstCharacter rune) ([]rune, bool) {
	if firstCharacter == []rune("ご")[0] {
		return []rune("ごりら"), true
	}
	return []rune{}, false
}

// Init regist this topic
func Init(regist func(types.Topic)) {
	fmt.Println("init")
	shiritori := Shiritori{}
	regist(shiritori)
}
