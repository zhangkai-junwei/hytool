package hySystem

import (
	"fmt"
	"hyTool/tcpServer"
	"io"
	"net"
	"strings"
)

type cmdFunc func(conn net.Conn, bytes []byte)

type HyMsgServer struct {
	Protocol
	server                *tcpServer.ServerInterface
	connSubcriptionMsgMap map[net.Conn][]byte
	cmdFuncMap            map[byte]cmdFunc
}

var system HyMsgServer

func Start() {
	system.connSubcriptionMsgMap = make(map[net.Conn][]byte)
	system.cmdFuncMap = make(map[byte]cmdFunc)
	system.server = &tcpServer.ServerInterface{}
	system.server.Start("8080", recData)
	registerCmdFunc(0x01, subcriptionCmd)
	registerCmdFunc(0x02, publishCmd)
	runAllApp()
}

func registerCmdFunc(cmd byte, cmdCallback cmdFunc) {
	system.cmdFuncMap[cmd] = cmdCallback
}

func recData(conn net.Conn) {
	for {
		bytes := make([]byte, 1024)
		n, err := tcpServer.ReadInTime(conn, bytes, 1000)
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
		recBytes, err := system.DecodeMsg(bytes[:n])
		if err != nil {
			fmt.Println(err)
			continue
		}
		cmdFunc, ok := system.cmdFuncMap[recBytes[2]]
		if ok {
			cmdFunc(conn, recBytes)
		}
	}
}
