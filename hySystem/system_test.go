package hySystem

import (
	"fmt"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	runAllApp()
	for {
		time.Sleep(time.Second)
	}
}

func TestSystem(t *testing.T) {
	Start()
	for {
		time.Sleep(time.Second)
	}
}

func TestRunApp(t *testing.T) {
	runApp("timer.exe")
	for {
		time.Sleep(time.Second)
	}
}

func TestSubscription(t *testing.T) {
	mapBuf := make(map[int][]byte)
	buf := []byte{6, 2, 3}
	mapBuf[1] = buf

	for k, mapL := range mapBuf {
		fmt.Println(k)
		for _, b := range mapL {
			fmt.Println(b)
		}
	}
}
