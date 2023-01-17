package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{New:func() interface{} {return "-"}}

	pool.Put("Farhan")
	pool.Put("Rafid")
	pool.Put("Syauqi")

	for i := 0; i <10;i++{
		go func(){
			data := pool.Get()
			fmt.Println(data)
			
			
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Selesai")

}