package test

import "fmt"

type busyState struct {
}

func (m *busyState) OnEntryOnce() {
	fmt.Println("busy OnEntryOnce")
}

func (m *busyState) OnEntryCircle() {
	fmt.Println("busy OnEntryCircle")
}

func (m *busyState) OnExit() {
	fmt.Println("busy OnExit")
}
