package sync

import (
	"fmt"
	"sync"
	"testing"
)

var counter int = 0

func AddOnce() {
	counter++
}

func TestAddOnce(t *testing.T) {

	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(AddOnce)
			group.Done()
		}()
	}
	fmt.Println(counter)

}
