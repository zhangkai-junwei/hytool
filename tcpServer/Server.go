package tcpServer

import (
	"context"
	"fmt"
	"net"
	"sync"
)

type Callback func(conn net.Conn, ch chan []byte)

type ServerInterface struct {
	bufSize  int
	ch       chan []byte
	progress Callback
	listen   net.Listener
	ctx      context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
}

func (self *ServerInterface) Start(port string, pCallback Callback) (err error) {
	self.listen, err = net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		return err
	}
	self.ctx, self.cancel = context.WithCancel(context.Background())
	self.bufSize = 1024
	self.ch = make(chan []byte, self.bufSize)
	self.progress = pCallback
	self.wg.Add(1)
	go self.listenRoutine(self.ctx)
	return nil
}

func (self *ServerInterface) SetBufSize(size int) {
	self.bufSize = size
}

func (self *ServerInterface) listenRoutine(ctx context.Context) {
	for {
		conn, err := self.listen.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		self.wg.Add(1)
		go self.recRoutine(conn, ctx)
	}
	self.wg.Done()
}

func (self *ServerInterface) recRoutine(conn net.Conn, ctx context.Context) {
	bytes := make([]byte, self.bufSize)
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			continue
		}

		buf := bytes[:n]
		select {
		case self.ch <- buf:
		case <-ctx.Done():
			goto end
		}
		self.progress(conn, self.ch)
	}
end:
	self.wg.Done()
}

func (self *ServerInterface) Stop() {
	self.listen.Close()
	self.cancel()
	self.wg.Wait()
}
