package serial

import (
	serial "hyTool/serial/goserial"
	"io"
)

type Callback func(bytes []byte, len int)

type SerialInterface struct {
	bufSize   int
	serialIO  io.ReadWriteCloser
	serialRec Callback
}

func (self *SerialInterface) OpenPort(name string, baud int, pCallback Callback) (err error) {
	c := &serial.Config{Name: name, Baud: baud, Parity: serial.ParityNone}
	self.serialIO, err = serial.OpenPort(c)
	if err != nil {
		return err
	}
	self.bufSize = 1024
	self.serialRec = pCallback
	go self.routine()
	return err
}

func (self *SerialInterface) SetBufSize(size int) {
	self.bufSize = size
}

func (self *SerialInterface) Send(bytes []byte) (int, error) {
	return self.serialIO.Write(bytes)
}

func (self *SerialInterface) routine() {
	bytes := make([]byte, self.bufSize)
	for {
		n, err := self.serialIO.Read(bytes)
		if err != nil {
			continue
		}
		buf := bytes[:n]
		self.serialRec(buf, n)
	}
}
