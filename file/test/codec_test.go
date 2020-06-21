package test

import (
	"fmt"
	"hyTool/file/codec"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestGob(t *testing.T) {
	p := Person{
		Name: "zhangkai",
		Age:  10,
	}
	handle, err := codec.CreateCodec("gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	handle.Encode("person.gob", p)
	var q Person
	handle.Decode("person.gob", &q)
	fmt.Println(q)
}

func TestJson(t *testing.T) {
	p := Person{
		Name: "zhangkai",
		Age:  10,
	}
	handle, err := codec.CreateCodec("json")
	if err != nil {
		fmt.Println(err)
		return
	}
	handle.Encode("person.json", p)
	var q Person
	handle.Decode("person.json", &q)
	fmt.Println(q)
}

func TestXml(t *testing.T) {
	p := Person{
		Name: "zhangkai",
		Age:  10,
	}
	handle, err := codec.CreateCodec("xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	handle.Encode("person.xml", p)
	var q Person
	handle.Decode("person.xml", &q)
	fmt.Println(q)
}
