package manager

import (
	"context"
	"fmt"
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

func (m *Manager) Init() {
	state := stateless.NewStateMachine(StateStart)

	// Initialize main state
	state.Configure(StateStart).
		Permit(EventStartService, StateIdle)

	state.Configure(StateIdle).
		OnEntry(m.StateIdle).
		Permit(EventStartBackup, StateProcessing)

	state.Configure(StateProcessing).
		OnEntry(m.StateProcessing).
		Permit(EventFinishBackup, StateIdle)

	m.state = state
	m.FireIdle()
}

func (m *Manager) StateStart(ctx context.Context, args ...any) error {
	fmt.Println("current:", StateStart)
	return nil
}

func (m *Manager) StateIdle(ctx context.Context, args ...any) error {
	fmt.Println("current:", StateIdle)
	return nil
}

func (m *Manager) StateProcessing(ctx context.Context, args ...any) error {
	fmt.Println("current:", StateProcessing)
	fmt.Println("start processing...")
	fmt.Println("finish processing...")
	// TODO: call 'fsm task' and waiting channel
	m.state.Fire(EventFinishBackup)
	return nil
}

func (m *Manager) PrintGraph() {
	graph := m.state.ToGraph()
	fmt.Println("Manager graph")
	fmt.Println(graph)
}

func (m *Manager) FireIdle() {
	err := m.state.Fire(EventStartService)
	if err != nil {
		fmt.Println("cannot fire idle")
		fmt.Println("current state is", m.state.MustState())
	}
}

func (m *Manager) FireProcessing() {
	err := m.state.Fire(EventStartBackup)
	if err != nil {
		fmt.Println("cannot fire processing")
		fmt.Println("current state is", m.state.MustState())
	}
}
