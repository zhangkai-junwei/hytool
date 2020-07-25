package tcpServer

import (
	"fmt"
	"net"
)

type Callback func(conn net.Conn, ch chan []byte)

type ServerInterface struct {
	bufSize  int
	ch       chan []byte
	progress Callback
}

func (self *ServerInterface) Start(port string, pCallback Callback) error {
	listen, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		return err
	}
	self.bufSize = 1024
	self.ch = make(chan []byte, self.bufSize)
	self.progress = pCallback
	go self.listenRoutine(listen)
	return nil
}

func (self *ServerInterface) SetBufSize(size int) {
	self.bufSize = size
}

func (self *ServerInterface) listenRoutine(listen net.Listener) {
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go self.recRoutine(conn)
	}
}

func (self *ServerInterface) recRoutine(conn net.Conn) {
	bytes := make([]byte, self.bufSize)
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			continue
		}
		buf := bytes[:n]
		self.ch <- buf
		self.progress(conn, self.ch)
	}
}
