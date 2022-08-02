package libs

import "fmt"

func StringCount(input string) {
	runes := []rune(input)

	fmt.Println("number of char =", len(runes))
	fmt.Println("length of string =", len(input))
}
