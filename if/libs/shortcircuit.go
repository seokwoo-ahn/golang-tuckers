package libs

import "fmt"

func checkshortcircuit() bool {
	// 실행되지 않음
	fmt.Println("short circuit")
	return true
}

func ShortCircuit() {
	if false && checkshortcircuit() {
		fmt.Println("hello world!")
	}
}
