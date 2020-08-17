package test

import "fmt"

type idleState struct {
}

func (m *idleState) OnEntryOnce() {
	fmt.Println("idle OnEntryOnce")
}

func (m *idleState) OnEntryCircle() {
	fmt.Println("idle OnEntryCircle")
}

func (m *idleState) OnExit() {
	fmt.Println("idle OnExit")
}
