package sync

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for time := range ticker.C {
		fmt.Println(time)
	}
}
