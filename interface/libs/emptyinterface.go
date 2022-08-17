package libs

import (
	"fmt"
	"reflect"
)

func EmptyInterface(input interface{}) {
	fmt.Println(input)
	fmt.Println("input type:", reflect.TypeOf(input))
}
