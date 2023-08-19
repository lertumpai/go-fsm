package manager

import (
	"context"
	"fmt"
)

func (m *Manager) stateIdle(ctx context.Context, args ...any) error {
	fmt.Println("current:", StateIdle)
	return nil
}

func (m *Manager) stateProcessing(ctx context.Context, args ...any) error {
	fmt.Println("current:", StateProcessing)
	fmt.Println("start processing...")
	// TODO: call 'fsm task' and waiting channel
	fmt.Println("finish processing...")
	m.state.Fire(EventFinishBackup)
	return nil
}

func (m *Manager) PrintGraph() {
	graph := m.state.ToGraph()
	fmt.Println("Manager graph")
	fmt.Println(graph)
}

func (m *Manager) FireIdle() {
	err := m.state.Fire(EventStart)
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
