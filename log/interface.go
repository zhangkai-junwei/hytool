package log

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	Hour = 60 * 60
	Day  = Hour * 24
	Week = 7 * Day
)

type HyLog struct {
	consoleOutFlag bool
	infoFlag       bool
	debugFlag      bool
	errorFlag      bool
	path           string
	fileName       string
	buf            bytes.Buffer
	prefixTime     string
	saveTime       int64
}

func (m *HyLog) Start(path, fileName string, saveTime int64) {
	m.path = path
	m.fileName = fileName
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	log.SetOutput(&m.buf)
	m.saveTime = saveTime
}

func (m *HyLog) EnableConsoleOut(flag bool) {
	m.consoleOutFlag = flag
}

func (m *HyLog) EnableInfo(flag bool) {
	m.infoFlag = flag
}

func (m *HyLog) EnableDebug(flag bool) {
	m.debugFlag = flag
}

func (m *HyLog) EnableError(flag bool) {
	m.errorFlag = flag
}

func (m *HyLog) getPrefixTime() string {
	return time.Now().Format("2006-01-02")
}
func (m *HyLog) writeInFile(bytes []byte) error {
	if m.prefixTime != m.getPrefixTime() {
		m.prefixTime = m.getPrefixTime()
		CheckFileIsCanDelete(m.path, m.saveTime)
	}

	fileName := m.path + m.getPrefixTime() + "_" + m.fileName + ".log"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	defer f.Close()

	if err != nil {
		return err
	}

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
func (m *HyLog) LogDebug(format string, a ...interface{}) {
	if m.debugFlag {
		log.SetPrefix("[debug]")
		log.Printf(format, a...)
		len := m.buf.Len()
		buf := make([]byte, len)
		readLen, err := m.buf.Read(buf)
		if readLen != len {
			fmt.Println("read len is err")
			m.writeInFile([]byte("[debug] read len is err"))
		}
		if err != nil {
			fmt.Println(err.Error())
			m.writeInFile([]byte(err.Error()))
		}
		m.writeInFile(buf)
		if m.consoleOutFlag {
			fmt.Println(string(buf))
		}

	}

}

func (m *HyLog) LogInfo(format string, a ...interface{}) {
	if m.infoFlag {
		log.SetPrefix("[info]")
		log.Printf(format, a...)
		len := m.buf.Len()
		buf := make([]byte, len)
		readLen, err := m.buf.Read(buf)
		if readLen != len {
			fmt.Println("read len is err")
			m.writeInFile([]byte("[debug] read len is err"))
		}
		if err != nil {
			fmt.Println(err.Error())
			m.writeInFile([]byte(err.Error()))
		}
		m.writeInFile(buf)
		if m.consoleOutFlag {
			fmt.Println(string(buf))
		}
	}
}

func (m *HyLog) LogError(format string, a ...interface{}) {
	if m.errorFlag {
		log.SetPrefix("[error]")
		log.Printf(format, a...)
		len := m.buf.Len()
		buf := make([]byte, len)
		readLen, err := m.buf.Read(buf)
		if readLen != len {
			fmt.Println("read len is err")
			m.writeInFile([]byte("[debug] read len is err"))
		}
		if err != nil {
			fmt.Println(err.Error())
			m.writeInFile([]byte(err.Error()))
		}
		m.writeInFile(buf)
		if m.consoleOutFlag {
			fmt.Println(string(buf))
		}
	}
}
