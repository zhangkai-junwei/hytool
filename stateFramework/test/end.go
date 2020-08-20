package test

import "fmt"

type endState struct {
}

func (m *endState) OnEntry() {
	fmt.Println("end OnEntryOnce")
}

func (m *endState) OnExit() {
	fmt.Println("end OnExit")
}
