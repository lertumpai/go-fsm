package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// MySQL database credentials
	dbHost := "localhost" // Replace with the actual database host
	dbUsername := "root"
	dbPassword := "root"
	dbName := "my_database"

	// Path to mysqldump executable
	mysqldumpPath := "mysqldump" // Use the actual path if it's different

	// Output file for the dump
	outputFile := "backup.sql"

	// Construct the command
	cmd := exec.Command(mysqldumpPath, "--host="+dbHost, "--user="+dbUsername, "--password="+dbPassword, dbName)

	// Create the output file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	// Set the command's stdout to the output file
	cmd.Stdout = file

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running mysqldump:", err)
		return
	}

	fmt.Println("Database dump completed successfully.")
}
