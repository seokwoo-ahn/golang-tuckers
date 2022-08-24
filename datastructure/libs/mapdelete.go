package libs

import "fmt"

func MapDelete() {
	m := make(map[string]string)

	m["cran"] = "berry"
	m["water"] = "melon"
	m["izone"] = "chaewon"

	delete(m, "cran")
	v, ok := m["cran"]

	fmt.Println("v:", v, "ok:", ok)
}
