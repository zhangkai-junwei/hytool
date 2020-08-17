package test

import "fmt"

type endState struct {
}

func (m *endState) OnEntryOnce() {
	fmt.Println("end OnEntryOnce")
}

func (m *endState) OnEntryCircle() {
	fmt.Println("end OnEntryCircle")
}

func (m *endState) OnExit() {
	fmt.Println("end OnExit")
}
