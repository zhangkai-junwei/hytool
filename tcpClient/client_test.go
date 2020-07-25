package tcpClient

import (
	"fmt"
	"testing"
	"time"
)

var client *ClientInterface

func clientRec(bytes []byte, len int) {
	fmt.Println("len=", len)
	fmt.Println(string(bytes))
	recbyte := bytes[:len]
	client.Send(recbyte)
}

func TestClient(t *testing.T) {
	client = &ClientInterface{}
	client.Start("127.0.0.1:8000", clientRec)
	time.Sleep(time.Second)
	client.Stop()
}
