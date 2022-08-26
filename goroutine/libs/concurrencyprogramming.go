package libs

import (
	"fmt"
	"sync"
	"time"
)

type job interface {
	do()
}

type squareJob struct {
	index int
}

func (j *squareJob) do() {
	fmt.Printf("%d 작업시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func ConcurrencyProgramming() {
	var jobList [10]job

	for i := 0; i < 10; i++ {
		jobList[i] = &squareJob{i}
	}

	var wgCP sync.WaitGroup
	wgCP.Add(10)

	for i := 0; i < 10; i++ {
		test := jobList[i]
		go func() {
			test.do()
			wgCP.Done()
		}()
	}
	wgCP.Wait()
}
