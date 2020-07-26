package tcpClient

import (
	"fmt"
	"testing"
	"time"
)

var client *ClientInterface

func clientRec(bytes []byte) {
	fmt.Println("len=", len(bytes))
	fmt.Println(string(bytes))
	client.Send(bytes)
}

func TestClient(t *testing.T) {
	client = &ClientInterface{}
	client.Start("127.0.0.1:8000", clientRec)
	time.Sleep(time.Second)
	client.Stop()
}
