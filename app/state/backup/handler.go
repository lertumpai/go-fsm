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

	folderPath := "mysql-backup"

	// Create the folder (forcefully) and any necessary parent folders
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Output file for the dump
	outputFile := fmt.Sprintf("./%s/%s.sql", folderPath, time.Now().Format(time.RFC3339))

	// Path to mysqldump executable
	mysqldumpPath := "mysqldump" // Use the actual path if it's different

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
	b.FireEventFinishExtract(outputFile)
	return nil
}

func (b *Backup) stateUploading(ctx context.Context, args ...any) error {
	fmt.Println("backup:", StateUploading)

	fileName := args[0].(string)
	fmt.Println("Start backup file", fileName)

	b.FireEventFinishUpload(fileName)
	return nil
}

func (b *Backup) stateFinish(ctx context.Context, args ...any) error {
	fmt.Println("backup:", StateFinish)
	fmt.Println("backup: cleanup file")

	fileName := args[0].(string)
	fmt.Println("backup: removing file", fileName)
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Error removing file:", err)
	}

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

func (b *Backup) FireEventFinishExtract(outputFile string) {
	fmt.Println("backup: finish extract")
	err := b.state.Fire(EventFinishExtract, outputFile)
	if err != nil {
		fmt.Println("backup: cannot fire EventFinishExtract")
		fmt.Println("backup: current state is", b.state.MustState())
	}
}

func (b *Backup) FireEventFinishUpload(fileName string) {
	fmt.Println("backup: finish upload")
	err := b.state.Fire(EventFinishUpload, fileName)
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
