package log

import (
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)

func GetFileCreateTime(path string) int64 {
	osType := runtime.GOOS
	fileInfo, _ := os.Stat(path)
	if osType == "windows" {
		wFileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
		tNanSeconds := wFileSys.CreationTime.Nanoseconds() /// 返回的是纳秒
		tSec := tNanSeconds / 1e9                          ///秒
		return tSec
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
