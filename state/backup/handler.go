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
