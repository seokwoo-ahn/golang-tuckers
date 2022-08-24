package libs

import "fmt"

func MapRange() {
	m := make(map[string]string)

	m["cran"] = "berry"
	m["water"] = "melon"
	m["izone"] = "chaewon"

	for k, v := range m {
		fmt.Println("k:", k, "v:", v)
	}
}
