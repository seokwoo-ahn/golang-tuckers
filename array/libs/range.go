package libs

import "fmt"

func Range() {
	var array [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for iter, value := range array {
		fmt.Println(iter, value)
	}
}
