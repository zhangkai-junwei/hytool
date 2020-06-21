package serial

import (
	serial "hyTool/serial/goserial"
	"io"
)

func OpenPort(name string, baud int) (io.ReadWriteCloser, error) {
	c := &serial.Config{Name: name, Baud: baud, Parity: serial.ParityNone}
	s, err := serial.OpenPort(c)
	return s, err
}
