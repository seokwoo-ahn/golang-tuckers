package libs

import (
	"bufio"
	"fmt"
	"os"
)

func read(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func write(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

func ReadFile(filename string) {
	line, err := read(filename)
	if err != nil {
		err = write(filename, "this is wirte file")
		if err != nil {
			fmt.Println("파일 생성에 실패하였습니다", err)
			return
		}
		line, err = read(filename)
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다", err)
			return
		}
	}
	fmt.Println("파일 내용:", line)
}
