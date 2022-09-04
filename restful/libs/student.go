package libs

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var StudentsMap map[int]Student
var LastId int

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}
