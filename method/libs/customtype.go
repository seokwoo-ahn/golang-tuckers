package libs

type CustomType int

func (c CustomType) Add(b int) int {
	return int(c) + b
}
