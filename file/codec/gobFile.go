package codec

import (
	"encoding/gob"
	"os"
)

type GobTool struct {
}

func (g *GobTool) Encode(name string, v interface{}) error {

	File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	defer File.Close()

	enc := gob.NewEncoder(File)

	if err := enc.Encode(v); err != nil {
		return err
	}
	return nil
}

func (g *GobTool) Decode(name string, v interface{}) error {
	File, err := os.Open(name)
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(File)
	err = dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
}
