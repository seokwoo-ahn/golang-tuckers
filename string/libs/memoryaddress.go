package libs

import (
	"fmt"
	"reflect"
	"unsafe"
)

func MemoryAddress() {
	var str string = "IZONE"
	var slice []byte = []byte(str)

	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceheader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringheader.Data)
	fmt.Printf("slice:\t%x\n", sliceheader.Data)
}
