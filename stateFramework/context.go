package stateFramework

import (
	"errors"
)

type Context struct {
	stateMap     map[interface{}]AbsState
	currentState interface{}
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
		state.OnEntry()
	} else {
		return errors.New("not register this state")
	}
	return nil
}

func (m *Context) ChangeState(nextState interface{}) error {
	if state, ok := m.stateMap[nextState]; ok {
		m.currentState = nextState
		state.OnEntry()
	} else {
		return errors.New("not register this state")
	}
	return nil
}

func (m *Context) GetCurrentState() interface{} {
	return m.currentState
}
