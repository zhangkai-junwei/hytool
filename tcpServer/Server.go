package tcpServer

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

type Callback func(conn net.Conn, ch chan []byte)

type ServerInterface struct {
	bufSize    int
	chSize     int
	ch         chan []byte
	internalMs int
	progress   Callback
	listen     net.Listener
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
}

func (self *ServerInterface) Start(port string, pCallback Callback) (err error) {
	self.listen, err = net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		return err
	}
	self.ctx, self.cancel = context.WithCancel(context.Background())
	self.bufSize = 1024
	self.chSize = 10
	self.ch = make(chan []byte, self.chSize)
	self.progress = pCallback
	self.wg.Add(1)
	go self.listenRoutine(self.ctx)
	return nil
}

func (self *ServerInterface) SetBufSize(size int) {
	self.bufSize = size
}

func (self *ServerInterface) GetBufSize() int {
	return self.bufSize
}

func (self *ServerInterface) SetReadTimeout(ms int) {
	self.internalMs = ms
}

func (self *ServerInterface) GetReadTimeout() int {
	return self.internalMs
}

func (self *ServerInterface) SetChSize(chSize int) {
	self.chSize = chSize
}

func (self *ServerInterface) GetChSize() int {
	return self.chSize
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
	timeout := time.Millisecond * time.Duration(self.internalMs)
	for {
		conn.SetReadDeadline(time.Now().Add(timeout))
		n, err := conn.Read(bytes)
		if err != nil {
			select {
			case <-ctx.Done():
				goto end
			default:

			}
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
