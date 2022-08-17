package libs

import (
	"fmt"
	"sort"
)

type Member struct {
	Id   int
	Name string
}

type Members []Member

func (m Members) Len() int           { return len(m) }
func (m Members) Less(i, j int) bool { return m[i].Id < m[j].Id }
func (m Members) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func StructSort() {
	slice := []Member{}
	fmt.Println(slice)

	slice = append(slice, Member{1, "name1"}, Member{2, "name2"}, Member{6, "name6"}, Member{3, "name3"})
	fmt.Println(slice)

	sort.Sort(Members(slice))
	fmt.Println(slice)
}
