package GoroutinePool

import (
	"fmt"
	"runtime"
)

type Pool struct {
	workers       []*worker
	idles         chan *worker
	tasksQueue    chan Gotask
	waitCompleted chan struct{}
}

func init() {
	numCPUs := runtime.NumCPU()
	fmt.Println("numCPUs=", numCPUs)
	runtime.GOMAXPROCS(numCPUs)
}

/*
*启动协程池
 */
func (self *Pool) Start() {
	go func() {
		defer close(self.waitCompleted)

		for task := range self.tasksQueue {
			(<-self.idles).tasks <- task
		}
		for _, worker := range self.workers {
			close(worker.tasks)
			<-worker.waitExit
		}
	}()
}

/*
*关闭线程池，阻塞等待所有的worker协程完成退出后，此函数才会返回
 */
func (self *Pool) Shutdown() {
	close(self.tasksQueue)
	<-self.waitCompleted
}

/*
*添加需要调度的任务。不能再Shutdown之后调用。会引发panic
 */
func (self *Pool) Add(task Gotask) {
	self.tasksQueue <- task
}

/*
*创建一个协程池对象，并指定协程池的大小
 */

func NewPool(num int) *Pool {
	numWorks := max(1, num)
	taskQueueSize := max(1, max(1, num/2))

	goes := &Pool{
		workers:       make([]*worker, numWorks),
		idles:         make(chan *worker, numWorks),
		tasksQueue:    make(chan Gotask, taskQueueSize),
		waitCompleted: make(chan struct{}),
	}

	for i := 0; i < numWorks; i++ {
		worker := newWorker()
		go worker.work(goes.idles)
		goes.workers[i] = worker
		goes.idles <- worker
	}

	return goes
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
