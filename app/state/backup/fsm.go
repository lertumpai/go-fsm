package backup

import (
	"github.com/qmuntal/stateless"
)

type Backup struct {
	state    *stateless.StateMachine
	response chan string
}

func Create() *Backup {
	return &Backup{
		state:    nil,
		response: nil,
	}
}

func (b *Backup) Init() *Backup {
	state := stateless.NewStateMachine(StateIdle)

	// Initialize main state
	state.Configure(StateIdle).
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
