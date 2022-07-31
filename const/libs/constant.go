package libs

func Constant() {
	const C int = 10
	// C = 20 cannot assign to constant
	// fmt.Println(&C) cannot take the address to constant
}
