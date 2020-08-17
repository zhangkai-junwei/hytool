package hyMessage

import (
	"errors"
	"fmt"
	"hyTool/message"
	"hyTool/tcpClientEx"
	"reflect"
)

type Callback func(msg interface{})

type Service struct {
	msgFun  map[byte]Callback
	client  *tcpClientEx.ClientInterface
	msgTool message.MessageTool
}

func (m *Service) clientRec(bytes []byte) {
	var msg interface{}
	respBuf, err := decode(bytes)
	if err != nil {
		return
	}

	if callBack, ok := m.msgFun[respBuf[1]]; ok {
		err := m.msgTool.RecMessage(msg, respBuf[4:])
		if err != nil {
			callBack(msg)
		}
	}
}

func (m *Service) Start(addr string) error {
	m.client = &tcpClientEx.ClientInterface{}
	err := m.client.Start(addr, m.clientRec)
	if err != nil {
		return err
	}
	m.msgFun = make(map[byte]Callback, 10)
	m.msgTool, err = message.CreateMessageTool("json")
	if err != nil {
		return err
	}
	return nil
}

func (m *Service) Subscription(msgType byte, pCallback Callback) error {

	bytes := encode(0x01, msgType, nil)
	buf, err := m.client.SendAndWaitResp(bytes)
	if err != nil {
		return err
	}
	respBuf, err := decode(buf)
	if err != nil {
		return err
	}
	if respBuf[0] == 0 {
		m.msgFun[msgType] = pCallback
		return nil
	} else {
		return errors.New("Subscription failed")
	}
}

func (m *Service) Publish(msg interface{}) error {
	dataValue := reflect.ValueOf(msg)

	dataValue = dataValue.Elem()
	msgType := dataValue.FieldByName("MsgType").Uint()
	fmt.Println("msgType=", msgType)

	buf, err := m.msgTool.SendMessage(msg)
	if err != nil {
		return err
	}
	sendBuf := encode(0x02, byte(msgType), buf)
	_, err = m.client.Send(sendBuf)
	return err
}

func (m *Service) ExchangeMsg(reqMsq, respMsg interface{}, timeOutMs int) error {
	m.client.SetInternalMs(timeOutMs)
	dataValue := reflect.ValueOf(reqMsq)

	dataValue = dataValue.Elem()
	msgType := dataValue.FieldByName("MsgType").Uint()
	fmt.Println("msgType=", msgType)

	bytes := encode(0x01, byte(msgType), nil)
	buf, err := m.client.SendAndWaitResp(bytes)
	if err != nil {
		return err
	}
	respBuf, err := decode(buf)
	if err != nil {
		return err
	}
	if respBuf[0] != 0 {
		return errors.New("Subscription failed")
	}

	buf1, err := m.msgTool.SendMessage(reqMsq)
	if err != nil {
		return err
	}
	sendBuf := encode(0x02, byte(msgType), buf1)
	bytes, err = m.client.SendAndWaitResp(sendBuf)
	err = m.msgTool.RecMessage(respMsg, bytes[4:])
	if err != nil {
		return err
	}
	return nil
}
