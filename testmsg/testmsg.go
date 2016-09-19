package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Yamashou/facebookbot/reply"
	"github.com/Yamashou/facebookbot/talk"
	"github.com/Yamashou/facebookbot/types"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")
	for scanner.Scan() {
		content := talk.TextContent{}
		content.SetText(scanner.Text())
		reply.Talk(talk.Event{
			SenderID:    types.UserID("test sender"),
			RecepientID: types.UserID("test recepient"),
			Content: content,
		})
		fmt.Print(">>> ")
	}
}
