package codec

import (
	"encoding/json"
	"os"
)

type JsonTool struct {
}

func (j *JsonTool) Encode(name string, v interface{}) error {

	filePtr, err := os.Create(name)

	if err != nil {
		return err
	}

	defer filePtr.Close()

	encoder := json.NewEncoder(filePtr)

	err = encoder.Encode(v)
	if err != nil {
		return err
	}
	return nil
}

func (j *JsonTool) Decode(name string, v interface{}) error {

	filePtr, err := os.Open(name)

	if err != nil {
		return err
	}

	defer filePtr.Close()

	decoder := json.NewDecoder(filePtr)

	err = decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}
