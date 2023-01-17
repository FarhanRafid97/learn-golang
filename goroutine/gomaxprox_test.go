package sync

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)


func TestGoMaxProx(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)
	group := sync.WaitGroup{}

	totalThreat := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThreat)
	x := 0
	for i:=0;i<100;i++{
		go func ()  {
			group.Add(1)
			for j:=0;j<100;j++{
				x = x + 1
			}
			group.Done()
		}()
	}

	totalGoRouting := runtime.NumGoroutine()
	fmt.Println("total goroutine",totalGoRouting)
	group.Wait()
	fmt.Println(x)
}