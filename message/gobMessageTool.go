package message

import (
	"bytes"
	"encoding/gob"
)

type GobMessageTool struct {
}

func (m *GobMessageTool) SendMessage(message interface{}) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(message)
	return buf.Bytes(), err
}

func (m *GobMessageTool) RecMessage(message interface{}, buf []byte) error {
	dec := gob.NewDecoder(bytes.NewReader(buf))
	err := dec.Decode(message)
	return err
}
