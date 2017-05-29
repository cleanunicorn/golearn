package main

import "fmt"
import "container/list"

// import "crypto/"

type ceva struct {
	gigel                    int
	costel                   int
	aldskjflakjfdalkjdflajdf int
}

func main() {
	var x list.List

	x.PushFront(1)
	x.PushFront(2)
	x.PushFront(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
