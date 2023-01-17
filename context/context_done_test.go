package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	fmt.Println("TOtal go routine", runtime.NumGoroutine())

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(3 * time.Second)
	fmt.Println("TOtal go routine", runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	fmt.Println("TOtal go routine", runtime.NumGoroutine())

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("counter", n)

	}

	time.Sleep(3 * time.Second)
	fmt.Println("TOtal go routine", runtime.NumGoroutine())
}
func TestContextWithDeadline(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()
	fmt.Println("TOtal go routine", runtime.NumGoroutine())

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("counter", n)

	}

	time.Sleep(3 * time.Second)
	fmt.Println("TOtal go routine", runtime.NumGoroutine())
}
