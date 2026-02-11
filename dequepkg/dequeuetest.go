package main

import (
	"container/list"
	"fmt"
)

func main() {
	DequeTest()
}

func DequeTest() {
	deq := list.New()
	deq.PushBack(10)
	deq.PushBack(20)
	deq.PushBack(30)

	fmt.Println(deq.Len())
	for e := deq.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		fmt.Println(e.Value)
	}

}
