package sync

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("uppss")

	time.Sleep(5 * time.Second)

}
func DisplayNUmber(num int) {
	fmt.Println(num)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		DisplayNUmber(i)
	}

	time.Sleep(5 * time.Second)

}
