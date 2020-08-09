package hySystem

import (
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	runAllApp()
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
