package libs

import (
	"fmt"
	"sort"
)

func Sort() {
	slice := []int{5, 6, 3, 4, 5}
	sort.Ints(slice)
	fmt.Println(slice)
}
