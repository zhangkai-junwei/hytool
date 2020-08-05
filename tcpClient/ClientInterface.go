package tcpClient

import (
	"net"
	"strings"
	"sync"
	"time"
)

type Callback func(bytes []byte)

type ClientInterface struct {
	host             string
	bufSize          int
	ch               chan []byte
	internalMs       int //连接时读取超时时间
	waitGrountineEnd sync.WaitGroup
	conn             net.Conn
	clientRec        Callback
	connectFlag      bool
	runFlag          bool
	waitRespFlag     bool
}

func (self *ClientInterface) Start(host string, pCallback Callback) (err error) {
	self.conn, err = net.Dial("tcp", host)

	self.host = host
	self.bufSize = 1024
	if err != nil {
		self.connectFlag = false
	} else {
		self.connectFlag = true
	}
	self.clientRec = pCallback
	self.internalMs = 1000
	self.runFlag = true
	self.waitRespFlag = false
	self.waitGrountineEnd.Add(1)
	self.ch = make(chan []byte, 1)
	go self.routine()
	return
}

func (self *ClientInterface) SetBufSize(size int) {
	self.bufSize = size
}

func (self *ClientInterface) GetBufSize() int {
	return self.bufSize
}

func (self *ClientInterface) SetInternalMs(ms int) {
	self.internalMs = ms
}

func (self *ClientInterface) GetInternalMs() int {
	return self.internalMs
}

func (self *ClientInterface) isConnect() bool {
	return self.connectFlag
}

func (self *ClientInterface) Send(bytes []byte) (int, error) {
	n, err := self.conn.Write(bytes)
	return n, err
}

func (self *ClientInterface) SendAndWaitResp(bytes []byte) ([]byte, error) {
	self.waitRespFlag = true
	_, err := self.conn.Write(bytes)
	if err != nil {
		return nil, err
	}
	timeout := time.Millisecond * time.Duration(self.internalMs)
	self.conn.SetReadDeadline(time.Now().Add(timeout))
	recBuf := <-self.ch
	return recBuf, nil
}

func (self *ClientInterface) routine() {
	var err error
	var n int

	timeout := time.Millisecond * time.Duration(self.internalMs)

	if !self.connectFlag {
		for self.runFlag {
			self.conn, err = net.DialTimeout("tcp", self.host, timeout)
			if err == nil {
				self.connectFlag = true
				break
			}
		}
	}

	for self.runFlag {
		buf := make([]byte, self.bufSize)

		self.conn.SetReadDeadline(time.Now().Add(timeout))
		n, err = self.conn.Read(buf)
		if err != nil {
			if strings.Contains(err.Error(), "timeout") {
				if self.waitRespFlag {
					self.waitRespFlag = false
					self.ch <- nil
					continue
				}

				self.clientRec(nil)
				continue
			}
			self.conn.Close()
			self.connectFlag = false

			for self.runFlag {
				self.conn, err = net.DialTimeout("tcp", self.host, timeout)
				if err == nil {
					self.connectFlag = true
					break
				}
			}
			continue
		}
		if self.waitRespFlag {
			self.waitRespFlag = false
			self.ch <- buf[:n]
			continue
		}
		self.clientRec(buf[:n])
	}
	self.waitGrountineEnd.Done()
}

func (self *ClientInterface) Stop() {
	self.runFlag = false
	self.waitGrountineEnd.Wait()
	if self.connectFlag {
		self.conn.Close()
	}
}
