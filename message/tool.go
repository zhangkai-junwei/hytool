package message

import (
	"errors"
)

type MessageTool interface {
	SendMessage(message interface{}) ([]byte, error)
	RecMessage(message interface{}, buf []byte) error
}

func CreateMessageTool(messageType string) (MessageTool, error) {
	switch messageType {
	case "json":
		return &JsonMessageTool{}, nil
	case "gob":
		return &GobMessageTool{}, nil
	default:
		return nil, errors.New("no exist type")
	}
}
