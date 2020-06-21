package message

import (
	"bytes"
	"encoding/json"
)

type JsonMessageTool struct {
}

func (m *JsonMessageTool) SendMessage(message interface{}) ([]byte, error) {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	err := enc.Encode(message)
	return buf.Bytes(), err
}

func (m *JsonMessageTool) RecMessage(message interface{}, buf []byte) error {
	dec := json.NewDecoder(bytes.NewReader(buf))
	err := dec.Decode(message)
	return err
}
