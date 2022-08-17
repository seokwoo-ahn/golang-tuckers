package main

import (
	"fmt"
	"golang-tuckers/method/libs"
)

func main() {
	// method delcaration
	fmt.Println("method declaration")
	a := &libs.Account{Balance: 100}
	libs.WithdrawFunc(a, 30)
	a.WithdrawMethod(30)
	// custom type
	fmt.Println("custom type")
	b := libs.CustomType(30)
	fmt.Println(b.Add(50))
}
