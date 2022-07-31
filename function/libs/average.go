package libs

import "fmt"

func Average(name string, val1 int, val2 int, val3 int) {
	total := val1 + val2 + val3
	avg := total / 3
	fmt.Println(name, "님 평균은", avg, "입니다.")
}
