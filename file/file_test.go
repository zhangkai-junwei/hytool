package file

import (
	"fmt"
	"testing"
)

func TestIni(t *testing.T) {
	val, err := GetValue("./zk.ini", "base", "app1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

func TestIniEx(t *testing.T) {
	iotl := &IniCtl{}
	iotl.LoadIniFile("./zk.ini")
	fmt.Println(iotl.GetValueAsString("base", "app1"))
	iotl.SetValue("base", "app2", "ok")
	fmt.Println(iotl.GetValueAsString("base", "app2"))
	iotl.SetValue("base1", "app3", "1256")
	fmt.Println(iotl.GetValueAsInt("base1", "app3"))
	iotl.SetValue("base1", "app4", "0")
	fmt.Println(iotl.GetValueAsBool("base1", "app4"))
}
