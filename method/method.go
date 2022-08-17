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
}
