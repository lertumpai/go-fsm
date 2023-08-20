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

func (m *Manager) FireEventFinishBackup() {
	err := m.state.Fire(EventFinishBackup)
	if err != nil {
		fmt.Println("cannot fire EventFinishBackup")
		fmt.Println("current state is", m.state.MustState())
	}
}

func (m *Manager) FireEventStartBackup() {
	err := m.state.Fire(EventStartBackup)
	if err != nil {
		fmt.Println("cannot fire EventStartBackup")
		fmt.Println("current state is", m.state.MustState())
	}
}
