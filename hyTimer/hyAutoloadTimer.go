package hyTimer

import "time"

type HyAutoloadTimer struct {
	timeRunFlag bool
	timer       *time.Ticker
	timeout     OnTimeout
}

func (self *HyAutoloadTimer) CreateTimer(ms int, pCallback OnTimeout) {
	self.timer = time.NewTicker(time.Duration(ms) * time.Millisecond)
}

func (self *HyAutoloadTimer) Start() {
	go func() {
		for self.timeRunFlag {
			<-self.timer.C
			if self.timeRunFlag {
				self.timeout()
			}
		}
	}()
}

func (self *HyAutoloadTimer) Reset(ms int) {
	self.timer = time.NewTicker(time.Duration(ms) * time.Millisecond)
}

func (self *HyAutoloadTimer) Stop() {
	self.timer.Stop()
}

func (self *HyAutoloadTimer) Delete() {
	self.timeRunFlag = false
}
