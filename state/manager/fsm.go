package manager

import (
	"github.com/qmuntal/stateless"
)

type Manager struct {
	state *stateless.StateMachine
}

func CreateManager() *Manager {
	return &Manager{
		state: nil,
	}
}

func (m *Manager) Init() *Manager {
	state := stateless.NewStateMachine(StateStart)

	// Initialize main state
	state.Configure(StateStart).
		Permit(EventStart, StateIdle)

	state.Configure(StateIdle).
		OnEntry(m.stateIdle).
		Permit(EventStartBackup, StateProcessing)

	state.Configure(StateProcessing).
		OnEntry(m.stateProcessing).
		Permit(EventFinishBackup, StateIdle)

	m.state = state
	m.FireIdle()

	return m
}
