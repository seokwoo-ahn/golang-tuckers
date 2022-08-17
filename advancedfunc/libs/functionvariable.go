package libs

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func FunctionVariable(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}
