package libs

import "fmt"

func Declaration() {
	str1 := "Hello\t'World'\n"
	str2 := `Hello\t'World'\n`

	fmt.Println("double quotation mark:", str1)
	fmt.Println("back quote:", str2)
}
