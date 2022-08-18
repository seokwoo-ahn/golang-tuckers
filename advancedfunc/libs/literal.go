package libs

type OpFunc func(a, b int) int

func Literal(op string) OpFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else {
		return nil
	}
}
