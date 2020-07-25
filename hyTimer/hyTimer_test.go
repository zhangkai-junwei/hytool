package hyTimer

import (
	"fmt"
	"testing"
	"time"
)

func event1() {
	fmt.Println("====", time.Now())
}

func event2() {
	fmt.Println("******", time.Now())
}
func TestTimer(t *testing.T) {
	timer := &hyTimer{}
	timer.CreateTimerOnce(1500, event1)
	timer.CreateTimerAutoLoad(1800, event2)
	fmt.Println("start", time.Now())
	for {
		time.Sleep(time.Second)
	}
}
