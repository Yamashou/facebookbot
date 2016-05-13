package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/m2mtu/facebookbot/reply"
	"github.com/m2mtu/facebookbot/endpoint"
	"github.com/m2mtu/facebookbot/state"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")
	for scanner.Scan() {
		content := reply.Get(state.State{ReceivedContents: []interface{}{endpoints.TextContent{Text: scanner.Text()}}})
		if c, ok := content.(endpoint.TextContent); ok {
			fmt.Printf(c.Text)
		}
		fmt.Print(">>> ")
	}
}
