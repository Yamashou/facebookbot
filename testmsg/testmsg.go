package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/m2mtu/facebookbot/reply"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")
	for scanner.Scan() {
		fmt.Println(reply.Get(scanner.Text()))
		fmt.Print(">>> ")
	}
}
