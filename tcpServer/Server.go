package tcpServer

import (
	"fmt"
	"net"
	"time"
)

type Callback func(conn net.Conn)

type ServerInterface struct {
	progress Callback
	listen   net.Listener
}

func (self *ServerInterface) Start(port string, pCallback Callback) (err error) {
	self.listen, err = net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		return err
	}
	self.progress = pCallback
	go self.listenRoutine()
	return nil
}

func (self *ServerInterface) listenRoutine() {
	for {
		conn, err := self.listen.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		go self.progress(conn)
	}
}

func ReadInTime(conn net.Conn, b []byte, ms int64) (n int, err error) {
	timeout := time.Millisecond * time.Duration(ms)
	conn.SetReadDeadline(time.Now().Add(timeout))
	n, err = conn.Read(b)
	return
}
