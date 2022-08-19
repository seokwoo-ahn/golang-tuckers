package libs

import (
	"container/list"
	"fmt"
)

func List() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertBefore(4, e4)
	v.InsertAfter(2, e1)

	frontCount := 0
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
		frontCount++
		fmt.Println()
	}
	fmt.Println("Front Count:", frontCount)

	backCount := 0
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
		backCount++
		fmt.Println()
	}
	fmt.Println("Back Count:", backCount)
}
