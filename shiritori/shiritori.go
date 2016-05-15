package shiritori

import (
	"fmt"
	"regexp"

	"github.com/m2mtu/facebookbot/endpoint"
	"github.com/m2mtu/facebookbot/types"
)

// TempState 構造体に入れられた情報は、ユーザーが一度その話題(この場合ではShiritori)を離れると削除されます。
type TempState struct {
	LastWord string
}

// PermState 構造体に入れられた情報は、恒久的に保持されます。(まだ未実装)
type PermState struct{}

// IsProper 関数は、ユーザーに関する情報を受けとり、この話題(この場合はShiritori)に入るべきかどうかを返します。
func IsProper(static types.StaticState) bool {
	// static.ReceivedContentにはユーザーから送られた情報が入っています。interface{}型なので、内容を見るには型アサーションが必要です。テキストデータの場合、endpoint.TextContent型となっています。
	if content, ok := static.ReceivedContent.(endpoint.TextContent); ok {
		return regexp.MustCompile(`しりとり`).MatchString(content.Text)
	}
	return false
}

// Talk 関数はこの話題に入っている時に呼ばれます。相手のユーザーID、送ってきた内容などを含むStaticState型を引数として取ります。引数の2、3番目と同じ型を返さなければなりません。
func Talk(static types.StaticState, temp TempState, perm PermState) (TempState, PermState) {
	fmt.Println("shiritori!")
	// tempが空だったら(Shiritoriに入って初めて実行されたら)
	if temp == (TempState{}) {
		temp = TempState{
			LastWord: "しりとり",
		}
	}
	if textContent, ok := static.ReceivedContent.(endpoint.TextContent); ok {
		text := textContent.Text
		textRunes := []rune(text)
		// endpoint.SendText関数でテキストデータを相手に送ることが出来ます。
		endpoint.SendText(temp.LastWord, static.OpponentID)
		if answer, isFound := wordsStartsWith(textRunes[len(textRunes)-1]); isFound {
			endpoint.SendText(string(answer), static.OpponentID)
			temp.LastWord = string(answer)
		}
	}
	return temp, perm
}

func wordsStartsWith(firstCharacter rune) ([]rune, bool) {
	if firstCharacter == []rune("ご")[0] {
		return []rune("ごりら"), true
	}
	return []rune{}, false
}
