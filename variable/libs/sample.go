package libs

import "fmt"

func Sample() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)

	var minimumWage int = 10
	var workingHour int = 20

	var income int = minimumWage * workingHour

	fmt.Println(minimumWage, workingHour, income)
}
