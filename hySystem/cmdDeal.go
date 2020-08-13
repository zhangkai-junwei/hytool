package hySystem

import "net"

func subcriptionCmd(conn net.Conn, bytes []byte) {
	sendBytes := system.EncodeMsg(0x01, bytes[3], nil)
	if msgTypes, ok := system.connSubcriptionMsgMap[conn]; ok {
		msgTypes = append(msgTypes, bytes[3])
		system.connSubcriptionMsgMap[conn] = msgTypes
	} else {
		buf := []byte{bytes[3]}
		system.connSubcriptionMsgMap[conn] = buf
	}
	conn.Write(sendBytes)
}

func publishCmd(conn net.Conn, bytes []byte) {
	for connList, msgTypes := range system.connSubcriptionMsgMap {
		if connList != conn {
			for _, msgType := range msgTypes {
				if msgType == bytes[3] {
					connList.Write(bytes)
				}
			}

		}
	}
}
