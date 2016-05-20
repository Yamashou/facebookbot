package shiritori

import (
	"fmt"
	"regexp"
	"math/rand"

	"github.com/m2mtu/facebookbot/endpoint"
	"github.com/m2mtu/facebookbot/types"
)

// TempState 構造体に入れられた情報は、ユーザーが一度その話題(この場合ではShiritori)を離れると削除されます。
type TempState struct {
	LastNoun types.Noun
}

// IsProper 関数は、ユーザーに関する情報を受けとり、この話題(この場合はShiritori)に入るべきかどうかを返します。
func IsProper(static types.StaticState) bool {
	// static.ReceivedContentにはユーザーから送られた情報が入っています。interface{}型なので、内容を見るには型アサーションが必要です。テキストデータの場合、endpoint.TextContent型となっています。
	if content, ok := static.ReceivedContent.(endpoint.TextContent); ok {
		return regexp.MustCompile(`しりとり`).MatchString(content.Text)
	}
	return false
}

// Talk 関数はこの話題に入っている時に呼ばれます。相手のユーザーID、送ってきた内容などを含むStaticState型を引数として取ります。引数の2、3番目と同じ型を返さなければなりません。
func Talk(static types.StaticState, temp TempState, perm types.PermState) (TempState, types.PermState, bool) {
	fmt.Println("shiritori!")
	// tempが空だったら(Shiritoriに入って初めて実行されたら)
	if temp == (TempState{}) {
		temp = TempState{
			LastNoun: types.Noun{Text: "しりとり"},
		}
	}
	if textContent, ok := static.ReceivedContent.(endpoint.TextContent); ok {
		text := textContent.Text
		textRunes := []rune(text)
		if len(textRunes) >= 1 {
			// endpoint.SendText関数でテキストデータを相手に送ることが出来ます。
			endpoint.SendText(temp.LastNoun.Text, static.OpponentID)
			if doesEndWith(textRunes, 'ん') {
				endpoint.SendText("んひひぃ", static.OpponentID)
				endpoint.SendText("しりとりは決着がついたね!", static.OpponentID)
				return temp, perm, false
			}
			if answer, isFound := nounsStartsWith(perm.LearnedNouns, textRunes[len(textRunes) - 1]); isFound {
				endpoint.SendText(answer.Text, static.OpponentID)
				temp.LastNoun = answer
			}
		}
	}
	return temp, perm, true
}

func doesEndWith(text []rune, lastCharacter rune) bool {
	if len(text) < 1 {
		return false
	}
	return text[len(text) - 1] == lastCharacter
}

func nounsStartsWith(dictionary []types.Noun, firstCharacter rune) (types.Noun, bool) {
	candidates := []types.Noun{}
	for _, noun := range dictionary {
		if firstCharacter == []rune(noun.Text)[0] {
			candidates = append(candidates, noun)
		}
	}
	if len(candidates) >= 1 {
		return candidates[rand.Int63n(int64(len(candidates)))], true
	}
	return types.Noun{}, false
}
