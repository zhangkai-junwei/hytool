package codec

import (
	"encoding/xml"
	"os"
)

type XmlTool struct {
}

func (x *XmlTool) Encode(name string, data interface{}) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return err
	}
	return nil
}

func (x *XmlTool) Decode(name string, data interface{}) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	err = decoder.Decode(data)

	if err != nil {
		return err
	}
	return nil
}
