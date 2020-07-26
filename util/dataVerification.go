package util

func Xor(bytes []byte) byte {
	var xor byte = 0
	for _, b := range bytes {
		xor ^= b
	}
	return xor
}

func GetCrc16(bytes []byte) uint16 {
	var crc16 uint16 = 0

	for _, b := range bytes {
		crc16 ^= uint16(b)
		for i := 0; i < 8; i++ {
			flag := crc16 & 0x0001
			crc16 >>= 1
			if flag == 1 {
				crc16 ^= 0xA001
			}
		}
	}
	return crc16
}
