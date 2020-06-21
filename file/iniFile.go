package file

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetValue(fileName, section, expectKey string) (string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var sectionName string

	for {
		lineStr, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}

		if lineStr[0] == ';' {
			continue
		}

		if lineStr[0] == '[' && lineStr[len(lineStr)-1] == ']' {
			//获取到当前的段名
			sectionName = lineStr[1 : len(lineStr)-1]
		} else if sectionName == section {
			//分割键值对
			pair := strings.Split(lineStr, "=")
			if len(pair) == 2 {
				//去掉键的空白字符
				key := strings.TrimSpace(pair[0])
				if key == expectKey {
					//返回去掉空白字符的值
					return strings.TrimSpace(pair[1]), nil
				}
			}
		}
	}
	return "", errors.New("not fild resource")
}
