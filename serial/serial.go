package serial

import (
	serial "github.com/serial"
	"io"
	"time"
)

type Callback func(bytes []byte, len int)

type SerialInterface struct {
	bytes     []byte
	serialIO  io.ReadWriteCloser
	serialRec Callback
}

func (self *SerialInterface) OpenPortNoBlock(name string, baud int, timeout time.Duration, pCallback Callback) (err error) {
	c := &serial.Config{Name: name, Baud: baud, ReadTimeout: timeout * time.Millisecond}
	self.serialIO, err = serial.OpenPort(c)
	if err != nil {
		return err
	}
	self.serialRec = pCallback
	self.bytes = make([]byte, 1024)

	go self.routine()
	return err
}

func (self *SerialInterface) OpenPort(name string, baud int, pCallback Callback) (err error) {
	c := &serial.Config{Name: name, Baud: baud}
	self.serialIO, err = serial.OpenPort(c)
	if err != nil {
		return err
	}
	self.serialRec = pCallback
	self.bytes = make([]byte, 1024)

	go self.routine()
	return err
}

func (self *SerialInterface) Send(bytes []byte) (int, error) {
	return self.serialIO.Write(bytes)
}

func (self *SerialInterface) routine() {

	for {
		n, err := self.serialIO.Read(self.bytes)
		if err != nil {
			continue
		}
		buf := self.bytes[:n]
		self.serialRec(buf, n)
	}
}
