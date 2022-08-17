package libs

import "fmt"

type Stringer interface {
	Declaration() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) Declaration() string {
	return fmt.Sprintf("안녕! 내 이름은 %s고 %d살이야", s.Name, s.Age)
}
