package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: tergo touch <filename>")
		return
	}

	command := os.Args[1]
	filename := os.Args[2]

	switch command {
	case "touch":
		err := createFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("File", filename, "created successfully")
	default:
		fmt.Println("Invalid command. Use 'tergo touch <filename>'")
	}
}

func createFile(filename string) error {
	// Get the absolute path
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(absPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
