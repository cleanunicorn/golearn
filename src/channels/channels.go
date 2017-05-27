package main

import "fmt"
import "time"

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("ping %d", i)
	}
}

func printer(c <-chan string) {
	for {
		fmt.Println("Waiting for message")
		msg := <-c
		fmt.Println("Received message")
		fmt.Println(msg)
		fmt.Println("---")
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("pong %d", i)
	}
}

func main() {
	var c = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
