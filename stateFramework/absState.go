package stateFramework

type AbsState interface {
	OnEntryOnce()
	OnEntryCircle()
	OnExit()
}
