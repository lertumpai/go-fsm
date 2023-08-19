package app

import (
	"context"
	"fmt"
	"github.com/qmuntal/stateless"
)

const (
	eventStartService = "StartService"
	eventStartBackup  = "StartBackup"
	eventFinishBackup = "FinishBackup"
)

const (
	stateStart      = "StateAppStart"
	stateIdle       = "StateAppIdle"
	stateProcessing = "StateProcessing"
)

type Manager struct {
	state *stateless.StateMachine
}

func CreateManager() *Manager {
	return &Manager{
		state: nil,
	}
}

func (m *Manager) Init() {
	state := stateless.NewStateMachine(stateStart)

	// Initialize main state
	state.Configure(stateStart).
		Permit(eventStartService, stateIdle)

	state.Configure(stateIdle).
		OnEntry(m.stateIdle).
		Permit(eventStartBackup, stateProcessing)

	state.Configure(stateProcessing).
		OnEntry(m.stateProcessing).
		Permit(eventFinishBackup, stateIdle)

	m.state = state
	m.FireIdle()
}

func (m *Manager) stateStart(ctx context.Context, args ...any) error {
	fmt.Println("current:", stateStart)
	return nil
}

func (m *Manager) stateIdle(ctx context.Context, args ...any) error {
	fmt.Println("current:", stateIdle)
	return nil
}

func (m *Manager) stateProcessing(ctx context.Context, args ...any) error {
	fmt.Println("current:", stateProcessing)
	fmt.Println("start processing...")
	fmt.Println("finish processing...")
	// TODO: call 'fsm task' and waiting channel
	m.state.Fire(eventFinishBackup)
	return nil
}

func (m *Manager) PrintGraph() {
	graph := m.state.ToGraph()
	fmt.Println("Manager graph")
	fmt.Println(graph)
}

func (m *Manager) FireIdle() {
	err := m.state.Fire(eventStartService)
	if err != nil {
		fmt.Println("cannot fire idle")
		fmt.Println("current state is", m.state.MustState())
	}
}

func (m *Manager) FireProcessing() {
	err := m.state.Fire(eventStartBackup)
	if err != nil {
		fmt.Println("cannot fire processing")
		fmt.Println("current state is", m.state.MustState())
	}
}
