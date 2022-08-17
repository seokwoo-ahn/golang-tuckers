package libs

import "fmt"

func VariableArgInterface(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}
