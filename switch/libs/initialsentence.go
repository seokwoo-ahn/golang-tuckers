package libs

import "fmt"

func getMyAge(age int) int {
	return age
}

func InitialSentence(age int) {
	switch age := getMyAge(age); true {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
}
