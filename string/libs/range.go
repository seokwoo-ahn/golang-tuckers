package libs

import "fmt"

func Range(input string) {
	for _, v := range input {
		fmt.Printf("타입:%T 값:%d 문자:%c\n", v, v, v)
	}
}
