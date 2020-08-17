package hyTimer

type OnTimeout func()

type HyTimer interface {
	CreateTimer(ms int, pCallback OnTimeout)
	Start()
	Reset(ms int)
	Stop()
	Delete()
}
