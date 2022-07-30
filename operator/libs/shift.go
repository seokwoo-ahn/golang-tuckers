package libs

import "fmt"

func Shift() {
	var x int8 = 4
	var y int8 = 64

	fmt.Printf("x:%08b x<<2: %08b x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("y:%08b y<<2: %08b y<<2: %d\n", y, y<<2, y<<2)

	var z int8 = 16
	var w int8 = -128
	var p int8 = -1
	var r uint8 = 128

	fmt.Printf("z:%08b z>>2: %08b z>>2: %d\n", z, z>>2, z>>2)
	fmt.Printf("w:%08b w>>2: %08b w>>2: %d\n", uint8(w), uint8(w>>2), w>>2)
	fmt.Printf("p:%08b p>>2: %08b p>>2: %d\n", uint8(p), uint8(p>>2), p>>2)
	fmt.Printf("r:%08b r>>2: %08b r>>2: %d\n", uint8(r), uint8(r>>2), w>>2)
}
