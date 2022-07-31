package libs

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func Declaration() {
	var house House
	house.Address = "광야"
	house.Size = 500
	house.Price = 16443
	house.Type = "메타버스"

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기 %d평\n", house.Size)
	fmt.Printf("가격: %.2f억원\n", house.Price)
	fmt.Println("타입:", house.Type)
}
