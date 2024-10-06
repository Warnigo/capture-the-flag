package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var currentDir = "/app/c"

func main() {
	fmt.Println("Windows CTF Challenge")
	fmt.Println("Find the flags hidden in the system!")
	fmt.Println("Type 'help' for available commands.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s> ", strings.ReplaceAll(currentDir, "/app/c", "C:"))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch strings.ToLower(args[0]) {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "help":
			fmt.Println("Available commands: cd, dir, ls, type, echo, cls, date, time, ver, exit")
		case "cd":
			if len(args) < 2 {
				fmt.Println(currentDir)
				continue
			}
			changeDirectory(args[1])
		case "dir", "ls":
			listDirectory()
		case "type":
			if len(args) < 2 {
				fmt.Println("Usage: type <filename>")
				continue
			}
			typeFile(args[1])
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "cls":
			fmt.Print("\033[H\033[2J")
		case "date":
			fmt.Println(time.Now().Format("Mon 01/02/2006"))
		case "time":
			fmt.Println(time.Now().Format("15:04:05"))
		case "ver":
			fmt.Println("Microsoft Windows [Version 10.0.19042.928]")
		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}

func changeDirectory(dir string) {
	if dir == ".." {
		if currentDir != "/app/c" {
			currentDir = filepath.Dir(currentDir)
		}
	} else if dir == "\\" || dir == "/" {
		currentDir = "/app/c"
	} else {
		newDir := filepath.Join(currentDir, dir)
		if _, err := os.Stat(newDir); os.IsNotExist(err) {
			fmt.Println("Directory not found")
			return
		}
		currentDir = newDir
	}
}

func listDirectory() {
	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, f := range files {
		if f.Name() == "main" || f.Name() == "main.go" {
			continue
		}
		if f.IsDir() {
			fmt.Printf("<DIR>\t%s\n", f.Name())
		} else {
			fmt.Printf("%d\t%s\n", f.Size(), f.Name())
		}
	}
}

func typeFile(filename string) {
	content, err := ioutil.ReadFile(filepath.Join(currentDir, filename))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content))
}
