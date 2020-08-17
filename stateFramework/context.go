package stateFramework

import (
	"errors"
	"time"
)

type Context struct {
	stateMap     map[interface{}]AbsState
	currentState string
	interval     int64
}

func (m *Context) RegisterState(stateName interface{}, state AbsState) {
	if m.stateMap == nil {
		m.stateMap = make(map[interface{}]AbsState)
	}
	m.stateMap[stateName] = state
}

func (m *Context) StartStateMachine(startState string) error {
	if state, ok := m.stateMap[startState]; ok {
		m.currentState = startState
		state.OnEntryOnce()
	} else {
		return errors.New("not register this state")
	}
	if m.interval < 10 {
		m.interval = 10
	}
	go m.routine()
	return nil
}

func (m *Context) ChangeState(nextState string) error {
	if state, ok := m.stateMap[nextState]; ok {
		m.currentState = nextState
		state.OnEntryOnce()
	} else {
		return errors.New("not register this state")
	}
	return nil
}

func (m *Context) SetInterval(t int64) error {
	if t < 10 {
		return errors.New("interval time is too short")
	}
	m.interval = t
	return nil
}

func (m *Context) GetInterval() int64 {
	return m.interval
}

func (m *Context) GetCurrentState() string {
	return m.currentState
}

func (m *Context) routine() {
	for {
		if state, ok := m.stateMap[m.currentState]; ok {
			state.OnEntryCircle()
		}
		time.Sleep(time.Millisecond * time.Duration(m.interval))
	}
}
