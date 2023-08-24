package manager

import (
	"context"
	"fmt"
)

func (m *Manager) stateIdle(ctx context.Context, args ...any) error {
	fmt.Println("manager:", StateIdle)
	return nil
}

func (m *Manager) stateProcessing(ctx context.Context, args ...any) error {
	fmt.Println("manager:", StateProcessing)
	fmt.Println("start processing...")
	m.backup.FireEventStartBackup("message eiei")
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
		fmt.Println("manager: cannot fire EventFinishBackup")
		fmt.Println("manager: current state is", m.state.MustState())
	}
}

func (m *Manager) FireEventStartBackup() {
	err := m.state.Fire(EventStartBackup)
	if err != nil {
		fmt.Println("manager: cannot fire EventStartBackup")
		fmt.Println("manager: current state is", m.state.MustState())
	}
}
