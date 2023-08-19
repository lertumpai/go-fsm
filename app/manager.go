package app

import (
	"context"
	"fmt"
	"github.com/qmuntal/stateless"
)

const (
	EventStartService = "StartService"
	EventStartBackup  = "StartBackup"
	EventFinishBackup = "FinishBackup"
)

const (
	StateStart      = "StateAppStart"
	StateIdle       = "StateAppIdle"
	StateProcessing = "StateProcessing"
)

type Manager struct {
	state *stateless.StateMachine
}

func CreateManager() Manager {
	return Manager{
		state: nil,
	}
}

func (m Manager) Init() {
	state := stateless.NewStateMachine(StateStart)

	// Initialize main state
	state.Configure(StateStart).
		Permit(EventStartService, StateIdle)

	state.Configure(StateIdle).
		OnEntry(m.stateIdle).
		Permit(EventStartBackup, StateProcessing)

	state.Configure(StateProcessing).
		OnEntry(m.stateProcessing).
		Permit(EventFinishBackup, StateIdle)

}

func (m Manager) stateIdle(ctx context.Context, args ...any) error {
	fmt.Println("state:", StateIdle)
	return nil
}

func (m Manager) stateProcessing(ctx context.Context, args ...any) error {
	fmt.Println("state:", StateProcessing)
	fmt.Println("start processing...")
	fmt.Println("finish processing...")
	// TODO: call 'fsm task' and waiting chanel
	m.FireIdle()
	return nil
}

func (m Manager) FireIdle() {
	err := m.state.Fire(StateIdle)
	if err != nil {
		fmt.Println("cannot fire idle")
		fmt.Println("current state is", m.state.MustState())
	}
}

func (m Manager) FireProcessing() {
	err := m.state.Fire(StateProcessing)
	if err != nil {
		fmt.Println("cannot fire processing")
		fmt.Println("current state is", m.state.MustState())
	}
}
