package learnword

import (
	"fmt"
	"math/rand"
	"strings"

	cabocha "github.com/ledyba/go-cabocha"
	"github.com/Yamashou/facebookbot/talk"
	"github.com/Yamashou/facebookbot/types"
)

const _init, _asked, _complete = 0, 1, 2

// TempState .
type TempState struct {
	Stage     int64
	AskedNoun cabocha.Token
}

// InitialTempState returns initial instance of TempState.
func InitialTempState() TempState {
	return TempState{
		Stage: _init,
	}
}

// IsProper returns the judgment should endter this topic.
func IsProper(static types.StaticState) bool {
	return rand.Int63n(2) == 1
}

// Talk method talk with user.
func Talk(static types.StaticState, temp TempState, perm types.PermState) (TempState, types.PermState, bool) {
	fmt.Println("learnword")
	if _content, ok := static.ReceivedContent.(talk.TextContent); ok {
		content, err := talk.AddDependentInfo(&_content)
		if err != nil {
			fmt.Println("Error in parsing the message: "+err.Error())
		}
		if temp.Stage == _init {
			sentence := content.Dependent()
			nounTokens := filterNouns(sentence)
			if len(nounTokens) >= 1 {
				nounToken := nounTokens[rand.Int63n(int64(len(nounTokens)))]
				if len(nounToken.Features) >= 8 {
					talk.SendText(nounToken.Body+"ってどういう意味?", static.OpponentID)
					temp.AskedNoun = nounToken
					temp.Stage = _asked
				}
			}
			return temp, perm, true
		} else if temp.Stage == _asked {
			fmt.Println("asked")
			sentence := content.Dependent()
			for _, chunk := range sentence.Chunks {
				fmt.Println(chunk.ToString())
			}
			if len(sentence.Chunks) >= 1 {
				lastChunk := sentence.Chunks[len(sentence.Chunks)-1]
				description := getStringAsNoun(sentence.Chunks, lastChunk.ID, temp.AskedNoun)
				talk.SendText("なるほど! "+temp.AskedNoun.Body+"は"+description+"なんだね!", static.OpponentID)
				talk.SendText("賢くなったかも!", static.OpponentID)
				perm.LearnedNouns = append(perm.LearnedNouns, types.Noun{
					Text:        temp.AskedNoun.Body,
					Yomi:        temp.AskedNoun.Features[7],
					Description: description,
				})
			}
		}
	}
	return temp, perm, false
}

func getStr(chunk cabocha.Chunk) string {
	str := ""
	for _, token := range chunk.Tokens {
		str += token.Body
	}
	return str
}

func getStringAsNoun(wholeChunks []cabocha.Chunk, id int, askedNoun cabocha.Token) string {
	chunk, isFound := findChunkByID(wholeChunks, id)
	if !isFound {
		fmt.Println("error in getStringAsNoun")
	}
	strBefore := ""
	strBeforeNoun := ""
	for _, token := range chunk.Tokens {
		strBefore += token.Body
		if token.Features[0] == "名詞" && token.Body != "ん" {
			strBeforeNoun = strBefore
		}
	}
	chunks := getChunksConnectTo(wholeChunks, id)
	wholeStr := ""
	for _, chunk := range chunks {
		if !strings.Contains(chunk.ToString(), askedNoun.Body) {
			wholeStr += getStr(chunk)
		}
	}
	return wholeStr + strBeforeNoun
}

func getChunksConnectTo(chunks []cabocha.Chunk, lastID int) []cabocha.Chunk {
	firstID, _ := getFirstChunkIDConnectTo(chunks, lastID)
	result := []cabocha.Chunk{}
	for iID := firstID; iID <= lastID-1; iID++ {
		chunk, _ := findChunkByID(chunks, iID)
		result = append(result, chunk)
	}
	return result
}

func findChunkByID(chunks []cabocha.Chunk, id int) (cabocha.Chunk, bool) {
	for _, chunk := range chunks {
		if chunk.ID == id {
			return chunk, true
		}
	}
	return cabocha.Chunk{}, false
}

func getFirstChunkIDConnectTo(chunks []cabocha.Chunk, id int) (int, bool) {
	for _, chunk := range chunks {
		if chunk.Link == id {
			if parentID, isFound := getFirstChunkIDConnectTo(chunks, chunk.ID); isFound {
				return parentID, true
			}
			return chunk.ID, true
		}
	}
	return id, false
}

func filterNouns(sentence cabocha.Sentence) []cabocha.Token {
	nouns := []cabocha.Token{}
	for _, chunk := range sentence.Chunks {
		for _, tok := range chunk.Tokens {
			if tok.Features[0] == "名詞" {
				nouns = append(nouns, tok)
			}
		}
	}
	return nouns
}
