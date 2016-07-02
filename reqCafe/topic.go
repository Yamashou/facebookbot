package reqCafe

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/m2mtu/facebookbot/talk"
	"github.com/m2mtu/facebookbot/types"
)

// TempState is local scope state.
type TempState struct {}

// InitialTempState returns initial object typed TempState.
func InitialTempState() TempState {
	return TempState{}
}

// IsProper judges
func IsProper(static types.StaticState) bool {
	if content, ok := static.ReceivedContent.(talk.TextContent); ok {
		exp := regexp.MustCompile(`(kondate|こんだて|献立|学食|めにゅー|メニュー)`)
		return exp.MatchString(content.Text())
	}
	return false
}

// Talk ...
func Talk(static types.StaticState, temp struct{}, perm types.PermState) (struct{}, types.PermState, bool) {
	fmt.Println("talk")
	talk.SendText(strings.Join(RtCafeInfo(time.Now()), "\n"), static.OpponentID)
	return temp, perm, false
}
