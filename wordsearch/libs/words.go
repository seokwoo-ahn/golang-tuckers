package libs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	LineNo int
	Line   string
}

type FindInfo struct {
	Filename string
	Result   []Result
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}
	files, err := GetFileList(path)
	if err != nil {
		fmt.Println("해당 경로에 입력한 파일이 존재하지 않습니다 err:", err, "path:", path)
		return findInfos
	}
	ch := make(chan FindInfo)
	cnt := len(files)
	recvCnt := 0

	for _, v := range files {
		go FindWordInFile(word, v, ch)
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}
	return findInfos
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []Result{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.Result = append(findInfo.Result, Result{LineNo: lineNo, Line: line})
		}
		lineNo++
	}
	ch <- findInfo
}
