package main

import (
	"fmt"
	"gopressor/src/channel-talk/talker"
)

func main() {
	var c = make(chan string)

	for i := 0; i < 1000000; i++ {
		go talker.EchoBack(c)
	}

	c <- "Hello 1"
	c <- "Hello 2"
	c <- "Hello 3"

	var input string
	fmt.Scanln(&input)
}
