package GoroutinePool

type Gotask struct {
	Param interface{}
	Task  func(interface{})
}

type worker struct {
	tasks    chan Gotask
	waitExit chan struct{}
}

func (self *worker) work(idles chan<- *worker) {
	defer close(self.waitExit)

	finishTask := func() {
		idles <- self
	}

	for taskFunc := range self.tasks {
		func() {
			defer finishTask()
			taskFunc.Task(taskFunc.Param)
		}()
	}
}

func newWorker() *worker {
	return &worker{
		tasks:    make(chan Gotask, 1),
		waitExit: make(chan struct{}),
	}
}
