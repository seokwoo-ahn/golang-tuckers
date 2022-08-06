package libs

import (
	"fmt"
	"reflect"
	"unsafe"
)

func AddString() {
	var str string = "carn"
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	addr1 := stringHeader.Data

	str += "berry"
	addr2 := stringHeader.Data

	str += "delicious"
	addr3 := stringHeader.Data

	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)
}
