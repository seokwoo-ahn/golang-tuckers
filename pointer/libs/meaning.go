package libs

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	fmt.Println("change data")
	arg.value = 999
	arg.data[100] = 999
}

func ChangePointer(arg *Data) {
	fmt.Println("change pointer")
	arg.value = 999
	arg.data[100] = 999
}

func Meaning() {
	var data Data

	ChangeData(data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	ChangePointer(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
