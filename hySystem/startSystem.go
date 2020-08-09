package hySystem

import (
	"fmt"
	"hyTool/file"
	"os/exec"
	"time"
)

func runAllApp() {
	i := 1
	for {
		key := fmt.Sprintf("app%d", i)
		i++
		val, err := file.GetValue("./appConfig.ini", "base", key)
		if err != nil {
			break
		}
		fmt.Println(val)
		go runApp(val)
	}

}

func runApp(path string) {
	for {
		cmd := exec.Command(path)
		cmd.Run()
		fmt.Println("./" + path)
		time.Sleep(5 * time.Second)
	}
}
