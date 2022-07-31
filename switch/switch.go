package main

import (
	"fmt"
	"golang-tuckers/switch/libs"
)

func main() {
	// initial sentence
	fmt.Println("initial sentence for switch")
	libs.InitialSentence(31)
	// switch iota
	fmt.Println("switch iota")
	libs.FavoriteColor(0)
}
