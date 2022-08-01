package libs

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age} // 함수가 끝나면 변수 u는 사라짐
	return &u               // 하지만 탈출 분석으로 u 메모리가 사라지지 않음, dangling error X
}

func EscapeAnalysis() {
	userPointer := NewUser("AAA", 23)

	fmt.Println(userPointer)
}
