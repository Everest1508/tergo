package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tergo <command> [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  touch <filename>: Create a new file")
		fmt.Println("  ls: List all files and folders in the current directory")
		fmt.Println("  curl <url>: Perform a GET request to the specified URL")
		fmt.Println("")
		fmt.Println("Author : Ritesh Mahale")
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
	case "neofetch":
		printSystemInfo()

	case "curl":
		if len(os.Args) != 3 {
			fmt.Println("Usage: tergo curl <url>")
			return
		}
		url := os.Args[2]
		err := performCurl(url)
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

func performCurl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}

func printSystemInfo() {
	fmt.Println("System Information:")
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("GOPATH:", os.Getenv("GOPATH"))

	printNetworkInfo()
}
func printNetworkInfo() {
	fmt.Println("\nNetwork Information:")
	fmt.Println("Hostname:", getHostname())
	fmt.Println("IP Addresses:", getIPAddresses())
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "Unknown"
	}
	return hostname
}

func getIPAddresses() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Unknown"
	}

	var ips []string
	for _, addr := range addrs {
		ip := addr.String()
		if strings.Contains(ip, ".") { // Filter IPv4 addresses
			ips = append(ips, ip)
		}
	}

	if len(ips) == 0 {
		return "Unknown"
	}
	return strings.Join(ips, ", ")
}