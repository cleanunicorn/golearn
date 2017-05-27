package main

import "fmt"
import "time"

func main() {
	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		i := 0
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("[1]", msg1, i)
				i = 0
			case msg2 := <-c2:
				fmt.Println("[2]", msg2, i)
				i = 0
			case <-time.After(time.Millisecond * 100):
				fmt.Printf("Timeout occured %d\n", i)
				i++
			}
		}
	}()

	// https://www.golang-book.com/books/intro/11

	var input string
	fmt.Scanln(&input)

}
