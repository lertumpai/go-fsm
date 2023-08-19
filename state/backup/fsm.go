package backup

import (
	"github.com/qmuntal/stateless"
)

type Backup struct {
	state *stateless.StateMachine
}

func CreateBackup() *Backup {
	return &Backup{
		state: nil,
	}
}

func (b *Backup) Init() *Backup {
	state := stateless.NewStateMachine(StateIdle)

	// Initialize main state
	state.Configure(StateIdle).
		OnEntry(b.stateIdle).
		Permit(EventStartBackup, StateExtracting)

	state.Configure(StateExtracting).
		OnEntry(b.stateExtracting).
		Permit(EventFinishExtract, StateUploading)

	state.Configure(StateUploading).
		OnEntry(b.stateUploading).
		Permit(EventFinishUpload, StateFinish)

	state.Configure(StateFinish).
		OnEntry(b.stateFinish).
		Permit(EventFinishBackup, StateIdle)

	b.state = state

	return b
}
