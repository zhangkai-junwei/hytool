package hyMessage

import (
	"errors"
	"hyTool/tcpClient"
)

type Callback func(bytes []byte)

type service struct {
	msgFun      map[byte]Callback
	client      *tcpClient.ClientInterface
	bytes       []byte
	needRecFlag bool
	recFlag     bool
	timeoutFlag bool
}

func (m *service) clientRec(bytes []byte) {
	if m.needRecFlag {
		m.needRecFlag = false
		if bytes != nil {
			m.recFlag = true
			m.bytes = bytes
		} else {
			m.timeoutFlag = true
		}
	}
}

func (m *service) Start(addr string) error {
	m.client = &tcpClient.ClientInterface{}
	err := m.client.Start(addr, m.clientRec)
	if err != nil {
		return err
	}
	m.needRecFlag = false
	m.msgFun = make(map[byte]Callback, 10)
	return nil
}

func (m *service) Subscription(msgType byte, pCallback Callback) error {

	bytes := encode(0x01, []byte{msgType})
	_, err := m.client.Send(bytes)
	if err != nil {
		return err
	}
	m.timeoutFlag = false
	m.recFlag = false
	m.needRecFlag = true
	for {
		if m.timeoutFlag {
			return errors.New("timeout")
		}
		if m.recFlag {

		}
	}
	return nil
}

func (m *service) Publish(msg interface{}) error {
	return nil
}
