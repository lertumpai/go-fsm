package backup

import (
	"context"
	"fmt"
	"go-fsm/config"
	"os"
	"os/exec"
	"time"
)

func (b *Backup) stateIdle(ctx context.Context, args ...any) error {
	fmt.Println("backup:", StateIdle)
	return nil
}

func (b *Backup) stateExtracting(ctx context.Context, args ...any) error {
	fmt.Println("backup:", StateExtracting)
	// Path to mysqldump executable
	mysqldumpPath := "mysqldump" // Use the actual path if it's different

	// Output file for the dump
	outputFile := "./backup-" + time.Now().Format(time.RFC3339) + ".sql"

	// Construct the command
	cmd := exec.Command(mysqldumpPath, "--host="+config.AppConfig.DbHost, "--user="+config.AppConfig.DbUsername, "--password="+config.AppConfig.DbPassword, config.AppConfig.DbName)

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
	}
	defer file.Close()

	// Set the command's stdout to the output file
	cmd.Stdout = file

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running mysqldump:", err)
	}

	fmt.Println("Database dump completed successfully.")
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
