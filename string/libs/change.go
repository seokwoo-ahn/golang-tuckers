package libs

import "fmt"

func Change() {
	var str string = "IZONE"
	var slice []byte = []byte(str)

	slice[2] = 'o'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}
