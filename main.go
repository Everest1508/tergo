package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tergo <command> [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  touch <filename>: Create a new file")
		fmt.Println("  ls: List all files and folders in the current directory")
		return
	}

	command := os.Args[1]

	switch command {
	case "touch":
		if len(os.Args) != 3 {
			fmt.Println("Usage: tergo touch <filename>")
			return
		}
		filename := os.Args[2]
		err := createFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("File", filename, "created successfully")
	case "ls":
		err := listFiles()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	default:
		fmt.Println("Invalid command. Use 'tergo <command> [arguments]'")
	}
}

func createFile(filename string) error {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	file, err := os.Create(absPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
)

func listFiles() error {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for i, entry := range dirEntries {
		if i > 0 {
			fmt.Print(" , ")
		}
		if entry.IsDir() {
			fmt.Print(colorGreen)
		} else {
			fmt.Print(colorRed)
		}
		fmt.Print(entry.Name())
		fmt.Print(colorReset)
	}
	fmt.Println()

	return nil
}