package main

import "fmt"
import "time"

func runner(n int) {
	fmt.Printf("Running for %v\n", n)
	for i := 0; i < n; i++ {
		fmt.Println(n, " -> ", i)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
}

func main() {
	fmt.Println("Hello go!")

	for i := 0; i < 10; i++ {
		go runner(i)
	}

	fmt.Println("BYE go!")

	var input string
	fmt.Scanln(&input)
}
