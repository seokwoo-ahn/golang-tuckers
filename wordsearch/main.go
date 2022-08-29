package main

import (
	"fmt"
	"golang-tuckers/wordsearch/libs"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 인수를 입력해주세요")
		return
	}
	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("찾으려는 단어:", word)
	libs.PrintAllFiles(files)
	// libs.PrintFile("./text/hamlet.txt")
	for _, path := range files {
		fmt.Println(libs.FindWordInAllFiles(word, path))
	}
	fmt.Println(libs.FindWordInFile(word, "./text/hamlet.txt"))
}
