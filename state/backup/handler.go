package backup

import (
	"context"
	"fmt"
)

func (b *Backup) stateIdle(ctx context.Context, args ...any) error {
	fmt.Println("current:", StateIdle)
	return nil
}

func (b *Backup) stateExtracting(ctx context.Context, args ...any) error {
	fmt.Println("current:", StateExtracting)
	return nil
}

func (b *Backup) stateUploading(ctx context.Context, args ...any) error {
	fmt.Println("current:", StateUploading)
	return nil
}

func (b *Backup) stateFinish(ctx context.Context, args ...any) error {
	fmt.Println("current:", StateFinish)
	return nil
}

func (b *Backup) FireEventStartBackup() {
	err := b.state.Fire(EventStartBackup)
	if err != nil {
		fmt.Println("cannot fire EventStartBackup")
		fmt.Println("current state is", b.state.MustState())
	}
}

func (b *Backup) FireEventFinishExtract() {
	err := b.state.Fire(EventFinishExtract)
	if err != nil {
		fmt.Println("cannot fire EventFinishExtract")
		fmt.Println("current state is", b.state.MustState())
	}
}

func (b *Backup) FireEventFinishUpload() {
	err := b.state.Fire(EventFinishUpload)
	if err != nil {
		fmt.Println("cannot fire EventFinishUpload")
		fmt.Println("current state is", b.state.MustState())
	}
}

func (b *Backup) FireEventFinishBackup() {
	err := b.state.Fire(EventFinishBackup)
	if err != nil {
		fmt.Println("cannot fire EventFinishBackup")
		fmt.Println("current state is", b.state.MustState())
	}
}

func (b *Backup) PrintGraph() {
	graph := b.state.ToGraph()
	fmt.Println("Backup graph")
	fmt.Println(graph)
}
