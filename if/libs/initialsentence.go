package libs

import "fmt"

func getMyAge(age int) (int, bool) {
	fmt.Println("my age is", age)
	return age, true
}

func InitialSentence() {
	if age, ok := getMyAge(31); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	// fmt.Println("Your age is", age) age 변수는 소멸 됨
}
