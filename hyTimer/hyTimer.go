package hyTimer

import "time"

type OnTimeout func()

type HyTimer interface {
	CreateTimer(ms int, pCallback OnTimeout)
	Start()
	Reset(ms int)
	Stop()
	Delete()
}

type HyOnceTimer struct {
	timeRunFlag bool
	timer       *time.Timer
	timeout     OnTimeout
}

type HyAutoloadTimer struct {
	timeRunFlag bool
	timer       *time.Ticker
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
