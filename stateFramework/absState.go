package stateFramework

type AbsState interface {
	OnEntry()
	OnExit()
}
