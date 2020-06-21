package example

import (
	"fmt"
	"hytool/qrocde"
	"testing"
)

func TestQrocde(t *testing.T) {
	qr := &qrocde.MQr{}
	qr.Gen("123456", "qrcode.png")
	s, e := qr.Parse("qrcode.png")
	fmt.Println(s, e)
}
