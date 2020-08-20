package test

import "fmt"

type idleState struct {
}

func (m *idleState) OnEntry() {
	fmt.Println("idle OnEntryOnce")
}

func (m *idleState) OnExit() {
	fmt.Println("idle OnExit")
}
