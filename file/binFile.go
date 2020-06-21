package file

import (
	"bytes"
	"encoding/binary"
	"os"
)

func WriteBinFile(name string, data interface{}, appendFlag bool) error {
	var flag int

	if appendFlag {
		flag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	} else {
		flag = os.O_CREATE | os.O_WRONLY
	}

	file, err := os.OpenFile(name, flag, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	var binBuf bytes.Buffer
	binary.Write(&binBuf, binary.LittleEndian, data)
	b := binBuf.Bytes()
	_, err = file.Write(b)

	if err != nil {
		return err
	}
	return nil
}

func ReadBinFile(name string, bytes []byte) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	len, err := file.Read(bytes)

	if err != nil {
		return 0, err
	}
	return len, nil
}
