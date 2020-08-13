package tcpServer

import (
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

func recData(conn net.Conn) {
	for {
		bytes := make([]byte, 10)
		n, err := ReadInTime(conn, bytes, 1000)
		if err != nil {
			if err == io.EOF {
				break
			}
			if strings.Contains(err.Error(), "i/o timeout") {
				continue
			}
			fmt.Println(err)
			break
		}
		conn.Write(bytes[:n])
	}

}

func TestClient(t *testing.T) {
	server := &ServerInterface{}
	server.Start("8080", recData)
	for {
		time.Sleep(time.Second)
	}
}
