package libs

import "fmt"

type Inter interface {
	Inter() string
}

type Str struct {
	Name string
	Age  int
}

func (s *Str) Inter() string {
	return fmt.Sprintln("이름:", s.Name, "나이:", s.Age)
}

func PrintStr(inter Inter) {
	s := inter.(*Str)
	fmt.Println("이름:", s.Name, "나이:", s.Age)
}
