package hyMessage

/*
*数据存储为小端模式
*请求协议帧格式 55 AA cmd len(2bytes) data crc16(从命令字开始计算)
*应答协议帧格式 55 AA status len(2bytes) data crc16
* cmd-0x01:订阅一条消息
* 返回：status：0x01-失败， 0x00-成功
* cmd-0x02:发布一条消息
* 无相应
 */

import (
	"errors"
	"hyTool/util"
)

func encode(cmd byte, bytes []byte) []byte {
	var buf []byte

	buf = append(buf, 0x55)
	buf = append(buf, 0xAA)
	buf = append(buf, cmd)
	len := len(bytes)
	buf = append(buf, byte(len&0xFF))
	buf = append(buf, byte((len>>8)&0xFF))
	buf = append(buf, bytes...)
	crc16 := util.GetCrc16(buf[2:])
	buf = append(buf, byte(crc16&0xFF))
	buf = append(buf, byte((crc16>>8)&0xFF))
	return buf
}

func decode(bytes []byte) ([]byte, error) {

	if bytes[0] != 0x55 || bytes[1] != 0xAA {
		return nil, errors.New("data frame err")
	}
	if bytes[2] != 0 {
		return nil, errors.New("return data err")
	}
	length := uint16(bytes[3] | (bytes[4] << 8))

	recLen := len(bytes)
	if uint16(recLen) < (length + 7) {
		return nil, errors.New("len is short")
	}
	crc16 := util.GetCrc16(bytes[2 : length+7])
	if byte(crc16&0xFF) != bytes[length+5] || byte((crc16>>8)&0xFF) != bytes[length+6] {
		return nil, errors.New("crc is error")
	}
	return bytes[5 : length+5], nil
}
