package libs

import (
	"math/rand"
	"time"
)

func NumGen() int {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100)
	return n
}
