package libs

import "fmt"

func Map() {
	m := make(map[string]string)
	m["cran"] = "berry"
	m["water"] = "melon"
	m["izone"] = "chaewon"

	fmt.Println(m["izone"])
	fmt.Println(m["cran"])
}
