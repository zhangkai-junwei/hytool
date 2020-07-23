package tcpClient

import (
	"fmt"
	"net"
)

type Callback func(bytes []byte, len int)

type ClientInterface struct {
	bufSize   uint16
	conn      net.Conn
	ClientRec Callback
}

func (self *ClientInterface) Start(host string, bufSize uint16, pCallback Callback) (err error) {
	self.conn, err = net.Dial("tcp", host)
	go self.routine()
	return
}

func (self *ClientInterface) Send(bytes []byte) (int, error) {
	n, err := self.conn.Write(bytes)
	return n, err
}

func (self *ClientInterface) routine() {
	for {
		buf := make([]byte, self.bufSize)
		n, err := self.conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		self.ClientRec(buf, n)
	}
}

func (self *ClientInterface) Stop() {
	self.conn.Close()
}
