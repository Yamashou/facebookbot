package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/m2mtu/facebookbot/reply"
	"github.com/m2mtu/facebookbot/talk"
	"github.com/m2mtu/facebookbot/types"
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
