package shiritori

import (
	"fmt"
	"math/rand"
	"regexp"
	"errors"

	"github.com/m2mtu/facebookbot/talk"
	"github.com/m2mtu/facebookbot/types"
	cabocha "github.com/ledyba/go-cabocha"
)

type stage int

const (
	_Initial stage = iota
	_Going
	_Finish
)

// TempState 構造体に入れられた情報は、ユーザーが一度その話題(この場合ではShiritori)を離れると削除されます。
type TempState struct {
	LastNoun types.Noun
	Stage    stage
}

// InitialTempState returns TempState object initialized.
func InitialTempState() TempState {
	return TempState{
		LastNoun: types.Noun{
			Text:        "しりとり",
			Description: "しりとり",
			Yomi:        "シリトリ",
		},
		Stage: _Initial,
	}
}

// IsProper 関数は、ユーザーに関する情報を受けとり、この話題(この場合はShiritori)に入るべきかどうかを返します。
func IsProper(static types.StaticState) bool {
	// static.ReceivedContentにはユーザーから送られた情報が入っています。interface{}型なので、内容を見るには型アサーションが必要です。テキストデータの場合、talk.TextContent型となっています。
	if content, ok := static.ReceivedContent.(talk.TextContent); ok {
		return regexp.MustCompile(`しりとり`).MatchString(content.Text())
	}
	return false
}

// Talk 関数はこの話題に入っている時に呼ばれます。相手のユーザーID、送ってきた内容などを含むStaticState型を引数として取ります。引数の2、3番目と同じ型を返さなければなりません。
func Talk(static types.StaticState, temp TempState, perm types.PermState) (TempState, types.PermState, bool) {
	fmt.Println("shiritori!")
	if textContent, ok := static.ReceivedContent.(talk.TextContent); ok {
		content, err := talk.AddDependentInfo(&textContent)
		if err != nil {
			fmt.Println(err)
		}

		switch temp.Stage {
		case _Initial:
			// talk.SendText関数でテキストデータを相手に送ることが出来ます。
			talk.SendText("いいよ、しりとりしよう!", static.OpponentID)
			temp.Stage = _Going
			return temp, perm, true

		case _Going:
			text := content.Text()
			textRunes := []rune(text)
			if len(textRunes) == 0 {
				return temp, perm, true
			}
			ok, msg, err := validAnswer(temp.LastNoun, content)
			if err != nil {
				fmt.Println(err)
				return temp, perm, false
			}
			if !ok {
				talk.SendText(msg, static.OpponentID)
				return temp, perm, true
			}
			if doesWin(content) {
				talk.SendText("んひひぃ", static.OpponentID)
				talk.SendText("しりとりは決着がついたね!", static.OpponentID)
				return temp, perm, false
			}
			last, err := lastRune(theNoun(content.Dependent()).Features[7])
			if err != nil {
				fmt.Println(err)
			}
			if answer, isFound := nounsStartsWith(perm.LearnedNouns, last); isFound {
				talk.SendText(answer.Text, static.OpponentID)
				temp.LastNoun = answer
			} else {
				talk.SendText("ま、まけた...", static.OpponentID)
				talk.SendText("うーん、難しすぎるね!", static.OpponentID)
				return temp, perm, false
			}
		}
	}
	return temp, perm, true
}

func doesWin(content talk.TextContentWithDependent) bool {
	yomi := theNoun(content.Dependent()).Features[7]
	c, err := lastRune(yomi)
	return c == rune('ン') && err == nil
}

func validAnswer(lastNoun types.Noun, content talk.TextContentWithDependent) (bool, string, error) {
	lastYomi := lastNoun.Yomi
	if !hasOnlyOneNoun(content.Dependent()) {
		return false, "ずるいよ!", nil
	}
	yomi := theNoun(content.Dependent()).Features[7]
	ok, err := areConnectable(lastYomi, yomi)
	if err != nil {
		return false, "", err
	}
	if !ok {
		return false, "つながんないよ!", nil
	}
	return true, "", nil
}

func areConnectable(a string, b string) (bool, error) {
	fmt.Println(a, b)
	lr, err := lastRune(a)
	if err != nil {
		return false, err
	}
	fr, err := firstRune(b)
	if err != nil {
		return false, err
	}
	return (lr == fr), nil
}

func lastRune(text string) (rune, error) {
	runes := []rune(text)
	if len(runes) == 0 {
		return rune(0), errors.New("Cannot get last character because the length of the string is 0.")
	}
	return runes[len(runes)-1], nil
}

func firstRune(text string) (rune, error) {
	runes := []rune(text)
	if len(runes) == 0 {
		return rune(0), errors.New("Cannot get first character because the length of the string is 0.")
	}
	return runes[0], nil
}

func hasOnlyOneNoun(sentence cabocha.Sentence) bool {
	return len(sentence.Chunks) == 1 && len(sentence.Chunks[0].Tokens) == 1 && sentence.Chunks[0].Tokens[0].Features[0] == "名詞" && len(sentence.Chunks[0].Tokens[0].Features) >= 8
}

func theNoun(sentence cabocha.Sentence) cabocha.Token {
	return sentence.Chunks[0].Tokens[0]
}

func doesEndWith(text []rune, lastCharacter rune) bool {
	if len(text) < 1 {
		return false
	}
	return text[len(text)-1] == lastCharacter
}

func nounsStartsWith(dictionary []types.Noun, firstCharacter rune) (types.Noun, bool) {
	candidates := []types.Noun{}
	for _, noun := range dictionary {
		if firstYomi, _ := firstRune(noun.Yomi); firstCharacter == firstYomi {
			candidates = append(candidates, noun)
		}
	}
	if len(candidates) >= 1 {
		return candidates[rand.Int63n(int64(len(candidates)))], true
	}
	return types.Noun{}, false
}
