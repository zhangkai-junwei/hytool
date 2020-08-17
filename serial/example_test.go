package serial

import (
	"fmt"
	"testing"
	"time"
)

func recData(bytes []byte, len int) {
	fmt.Println("len=", len)
	fmt.Println(bytes)
}

func TestQueue(t *testing.T) {
	serial := SerialInterface{}
	serial.OpenPort("Com1", 115200, recData)
	for {
		time.Sleep(time.Second)
	}
}
