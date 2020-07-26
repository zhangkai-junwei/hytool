package hyMessage

type AbsMessage struct {
	MsgType    byte
	MsgId      byte
	needReturn bool
}

func (m *AbsMessage) GetMsgType() byte {
	return m.MsgType
}

func (m *AbsMessage) SetMsgType(msgType byte) {
	m.MsgType = msgType
}

func (m *AbsMessage) GetNeedReturn() bool {
	return m.needReturn
}

func (m *AbsMessage) SetNeedReturn(needReturn bool) {
	m.needReturn = needReturn
}

func (m *AbsMessage) GetMsgId() byte {
	return m.MsgId
}

func (m *AbsMessage) SetMsgId(msgId byte) {
	m.MsgId = msgId
}
