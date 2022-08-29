package libs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	result   []Result
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}
	files, err := GetFileList(path)
	if err != nil {
		fmt.Println("해당 경로에 입력한 파일이 존재하지 않습니다 err:", err, "path:", path)
		return findInfos
	}
	for _, v := range files {
		findInfos = append(findInfos, FindWordInFile(word, v))
	}
	return findInfos
}

func FindWordInFile(word, filename string) FindInfo {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
	}
	defer file.Close()

	result := []Result{}
	lineNo := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			result = append(result, Result{lineNo: lineNo, line: line})
		}
		lineNo++
	}
	return FindInfo{filename: filename, result: result}
}
