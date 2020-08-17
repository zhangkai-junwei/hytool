package hyTimer

import "time"

type HyOnceTimer struct {
	timeRunFlag bool
	timer       *time.Timer
	timeout     OnTimeout
}

func (self *HyOnceTimer) CreateTimer(ms int, pCallback OnTimeout) {
	self.timer = time.NewTimer(time.Duration(ms) * time.Millisecond)
	self.timeout = pCallback
	self.timeRunFlag = true
}

func (self *HyOnceTimer) Start() {
	go func() {
		<-self.timer.C
		if self.timeRunFlag {
			self.timeout()
		}
	}()
}

func (self *HyOnceTimer) Reset(ms int) {
	self.timer.Reset(time.Duration(ms) * time.Millisecond)
}

func (self *HyOnceTimer) Stop() {
	self.timer.Stop()
}

func (self *HyOnceTimer) Delete() {
	self.timeRunFlag = false
}
