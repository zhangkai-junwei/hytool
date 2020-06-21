package test

import (
	"hyTool/serial"
	"testing"
)

func TestQueue(t *testing.T) {
	s, _ := serial.OpenPort("COM1", 115200)
	s.Write([]byte("asda"))
	buf := make([]byte, 128)
	s.Read(buf)
}
