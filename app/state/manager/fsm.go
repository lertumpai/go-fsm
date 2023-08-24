package manager

import (
	"github.com/qmuntal/stateless"
	"go-fsm/app/state/backup"
)

type Manager struct {
	state  *stateless.StateMachine
	backup *backup.Backup
}

type Config struct {
	Backup *backup.Backup
}

func Create(config Config) *Manager {
	return &Manager{
		state:  nil,
		backup: config.Backup,
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

	state.Fire(EventStart)
	m.state = state
	return m
}
