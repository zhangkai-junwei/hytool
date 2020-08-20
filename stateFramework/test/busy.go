package test

import "fmt"

type busyState struct {
}

func (m *busyState) OnEntry() {
	fmt.Println("busy OnEntryOnce")
}

func (m *busyState) OnExit() {
	fmt.Println("busy OnExit")
}
