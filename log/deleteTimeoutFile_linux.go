package log

import (
	"os"
	"runtime"
	"syscall"
	"time"
)

func GetFileCreateTime(path string) int64 {
	osType := runtime.GOOS
	fileInfo, _ := os.Stat(path)
	if osType == "linux" {
		stat_t := fileInfo.Sys().(*syscall.Stat_t)
		tCreate := int64(stat_t.Ctim.Sec)
		return tCreate
	}
	return time.Now().Unix()
}

func CheckFileIsCanDelete(path string, timeout int64) {
	files, _ := filepath.Glob(path + "*.log")
	for _, file := range files {
		now := time.Now().Unix()
		createTime := GetFileCreateTime(file)
		diffTime := now - createTime
		if diffTime > timeout {
			os.Remove(file)
		}
	}
}
