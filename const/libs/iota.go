package libs

import "fmt"

func Iota() {
	const (
		Red   int = iota
		Blue  int = iota
		Green int = iota
		Orange
		Purple
	)

	const (
		BitFlag1 uint = 1 << iota
		BitFlag2
		BitFlag3
	)

	fmt.Println("color iota", Red, Blue, Green, Orange, Purple)
	fmt.Println("bit iota", BitFlag1, BitFlag2, BitFlag3)
}
