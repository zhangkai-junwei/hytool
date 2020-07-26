package hyMessage

type absMessage struct {
	Cmd     byte
	MsgType byte
	MsgId   byte
}

func (m *absMessage) GetCmd() byte {
	return m.Cmd
}

func (m *absMessage) SetCmd(cmd byte) {
	m.Cmd = cmd
}

func (m *absMessage) GetMsgType() byte {
	return m.MsgType
}

func (m *absMessage) SetMsgType(msgType byte) {
	m.MsgType = msgType
}

func (m *absMessage) GetMsgId() byte {
	return m.MsgId
}

func (m *absMessage) SetMsgId(msgId byte) {
	m.MsgId = msgId
}
