package test

import (
	"fmt"
	"hyTool/stateFramework"
	"testing"
	"time"
)

func TestState(t *testing.T) {
	idle := &idleState{}
	busy := &busyState{}
	end := &endState{}
	s := &stateFramework.Context{}

	s.RegisterState("idle", idle)
	s.RegisterState("busy", busy)
	s.RegisterState("end", end)

	s.StartStateMachine("idle")

	for {
		time.Sleep(time.Second)
		fmt.Println("change busy")
		s.ChangeState("busy")
		time.Sleep(time.Second)
		fmt.Println("change end")
		s.ChangeState("end")
		time.Sleep(time.Second)
		fmt.Println("change idle")
		s.ChangeState("idle")
	}
}

func TestMap(t *testing.T) {
	s := make(map[interface{}]interface{})

	s["s"] = 5
	s[3] = 6
	s[6] = "dfas"
	for v, k := range s {
		fmt.Println(v, k)
	}
}
