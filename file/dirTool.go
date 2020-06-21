package file

import (
	"io/ioutil"
	"os"
	"strings"
)

//获取指定目录下的所有文件和目录

func GetFilesAndDirs(dirPath string, suffix string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, nil, err
	}

	pathSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)

	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, dirPath+pathSep+fi.Name())
			GetFilesAndDirs(dirPath+pathSep+fi.Name(), suffix)
		} else {
			if suffix == "*" {
				files = append(files, dirPath+pathSep+fi.Name())
			} else {
				ok := strings.HasSuffix(fi.Name(), suffix)
				if ok {
					files = append(files, dirPath+pathSep+fi.Name())
				}
			}
		}
	}

	return files, dirs, nil
}

//获取指定目录下的所有文件包含子目录下的文件
func GetAllFiles(dirPath string, suffix string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	pathSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, dirPath+pathSep+fi.Name())
		} else {
			if suffix == "*" {
				files = append(files, dirPath+pathSep+fi.Name())
			} else {
				ok := strings.HasSuffix(fi.Name(), suffix)
				if ok {
					files = append(files, dirPath+pathSep+fi.Name())
				}
			}
		}
	}

	for _, table := range dirs {
		temp, _ := GetAllFiles(table, suffix)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}
