package example

import (
	"fmt"
	"hytool/GoroutinePool"
	"testing"
	"time"
)

func task(v interface{}) {
	fmt.Println("i=:", v)
	time.Sleep(time.Second)
}

func TestQrocde(t *testing.T) {
	goes := GoroutinePool.NewPool(10)
	goes.Start()
	defer goes.Shutdown()

	for i := 0; i < 20; i++ {
		task := GoroutinePool.Gotask{
			i,
			task,
		}
		goes.Add(task)
	}
}
