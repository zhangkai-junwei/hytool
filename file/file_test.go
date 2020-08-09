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
