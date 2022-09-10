package todostruct

type Todo struct {
	ID        int    `json:"ID,omitempty"`
	Name      string `json:"Name"`
	Completed bool   `json:"Completed,omitempty"`
}

type Todos []Todo

type Success struct {
	Success bool `json:"Success"`
}

var TodoMap map[int]Todo
var LastId int = 0

func (t Todos) Len() int {
	return len(t)
}

func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Todos) Less(i, j int) bool {
	return t[i].ID > t[j].ID
}
