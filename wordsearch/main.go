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
	findInfos := []libs.FindInfo{}

	fmt.Println("찾으려는 단어:", word)
	for _, path := range files {
		findInfos = append(findInfos, libs.FindWordInAllFiles(word, path)...)
	}
	for _, findInfo := range findInfos {
		fmt.Println("file name:", findInfo.Filename)
		fmt.Println("-------------------------------")
		for _, result := range findInfo.Result {
			fmt.Println("LineNo:", result.LineNo, "\tcontent:", result.Line)
		}
		fmt.Println("-------------------------------")
		fmt.Println()
	}
}
