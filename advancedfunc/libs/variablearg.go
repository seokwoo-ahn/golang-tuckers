package libs

import "fmt"

func VariableArg(nums ...int) {
	count := 0
	sum := 0
	for _, v := range nums {
		sum += v
		count++
	}
	fmt.Println("count:", count)
	fmt.Println("sum:", sum)
}
