package libs

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func Separate() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("add-on complete!")

	wg.Add(3)
	go makeBody(tireCh)
	go installTire(tireCh, paintCh)
	go painCar(paintCh)

	wg.Wait()
	fmt.Println("not enough minerals!")
}

func makeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Tank"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func installTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "rail tire"
		paintCh <- car
	}
	close(paintCh)
	wg.Done()
}

func painCar(paintCh chan *Car) {
	index := 0
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "khaki"
		index++
		duration := time.Now().Sub(startTime)
		fmt.Printf("number %d tank, %.2f, ready to roll out!\n", index, duration.Seconds())
	}
	wg.Done()
}
