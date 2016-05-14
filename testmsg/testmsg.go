package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/m2mtu/facebookbot/endpoint"
	"github.com/m2mtu/facebookbot/reply"
	"github.com/m2mtu/facebookbot/types"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")
	for scanner.Scan() {
		reply.Talk(endpoint.Event{
			SenderID:    types.UserID("test sender"),
			RecepientID: types.UserID("test recepient"),
			Content: endpoint.TextContent{
				Text: scanner.Text(),
			},
		})
		fmt.Print(">>> ")
	}
}
