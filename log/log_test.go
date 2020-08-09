package log

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLog(t *testing.T) {
	log := &HyLog{}
	log.Start("./", "test", Day)
	log.EnableConsoleOut(true)
	log.EnableDebug(true)

	log.LogDebug("1jh2")
	log.LogDebug("3ggggg4")
	log.LogDebug("5ggjjjj6")
}

func TestDir(t *testing.T) {
	readerInfos, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, info := range readerInfos {
		if info.IsDir() {

		}
		fmt.Println(info.Name(), info.ModTime())
	}
}
