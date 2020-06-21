package codec

import "errors"

type Handle interface {
	Encode(name string, v interface{}) error
	Decode(name string, v interface{}) error
}

func CreateCodec(fileType string) (Handle, error) {
	switch fileType {
	case "gob":
		return &GobTool{}, nil
	case "json":
		return &JsonTool{}, nil
	case "xml":
		return &XmlTool{}, nil
	default:
		return nil, errors.New("no support file type")
	}
}
