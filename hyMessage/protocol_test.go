package hyMessage

import (
	"testing"
)

type Person struct {
	AbsMessage
	Name string
	Age  int
}

func TestEncode(t *testing.T) {
	s := &Service{}
	p := &Person{
		AbsMessage: AbsMessage{0x11, 0x12, true},
		Name:       "ll",
		Age:        0,
	}
	s.Publish(p)
}
