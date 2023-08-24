package backup

import (
	"context"
	"fmt"
)

func (b *Backup) stateIdle(ctx context.Context, args ...any) error {
	fmt.Println("backup:", StateIdle)
	return nil
}

func (b *Backup) stateExtracting(ctx context.Context, args ...any) error {
	fmt.Println("backup:", StateExtracting)
	b.FireEventFinishExtract()
	return nil
}

func (b *Backup) stateUploading(ctx context.Context, args ...any) error {
	fmt.Println("backup:", StateUploading)
	b.FireEventFinishUpload()
	return nil
}

func (b *Backup) stateFinish(ctx context.Context, args ...any) error {
	fmt.Println("backup:", StateFinish)
	b.FireEventFinishBackup()
	return nil
}

func (b *Backup) FireEventStartBackup(msg string) {
	fmt.Println("backup: start backup")
	err := b.state.Fire(EventStartBackup, msg)
	if err != nil {
		fmt.Println("backup: cannot fire EventStartBackup")
		fmt.Println("backup: current state is", b.state.MustState())
	}
}

func (b *Backup) FireEventFinishExtract() {
	fmt.Println("backup: finish extract")
	err := b.state.Fire(EventFinishExtract)
	if err != nil {
		fmt.Println("backup: cannot fire EventFinishExtract")
		fmt.Println("backup: current state is", b.state.MustState())
	}
}

func (b *Backup) FireEventFinishUpload() {
	fmt.Println("backup: finish upload")
	err := b.state.Fire(EventFinishUpload)
	if err != nil {
		fmt.Println("backup: cannot fire EventFinishUpload")
		fmt.Println("backup: current state is", b.state.MustState())
	}
}

func (b *Backup) FireEventFinishBackup() {
	fmt.Println("backup: finish backup")
	err := b.state.Fire(EventFinishBackup)
	if err != nil {
		fmt.Println("backup: cannot fire EventFinishBackup")
		fmt.Println("backup: current state is", b.state.MustState())
	}
}

func (b *Backup) PrintGraph() {
	graph := b.state.ToGraph()
	fmt.Println("Backup graph")
	fmt.Println(graph)
}
