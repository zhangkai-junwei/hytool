package tcpServer

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func recData(conn net.Conn, ch chan []byte) {
	bytes := <-ch
	fmt.Println("len=", len(bytes))
	conn.Write(bytes)
}

func TestClient(t *testing.T) {
	server := &ServerInterface{}
	server.Start("8080", recData)
	for {
		time.Sleep(time.Second)
	}
}
