package main

import (
	"fmt"
	"golang-tuckers/datastructure/libs"
)

func main() {
	// list
	fmt.Println("list")
	libs.List()
	// queue
	fmt.Println("queue")
	queue := libs.NewQueue()
	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	fmt.Println(queue.Pop())
	// stack
	fmt.Println("stack")
	stack := libs.NewStack()
	for i := 1; i < 5; i++ {
		stack.Push(i)
	}
	fmt.Println(stack.Pop())
}
