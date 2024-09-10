package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	currentDir := filepath.Dir(ex)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s: ", currentDir)
		scanner.Scan()
		inputString := scanner.Text()
		multipleInputStrings := strings.Split(inputString, "|")
		for _, inputString = range multipleInputStrings {
			inputString = strings.TrimSpace(inputString)
			splittedString := strings.Split(inputString, " ")
			switch splittedString[0] {
			case "ps":
				err = showProcesses()
				if err != nil {
					fmt.Println(err)
				}
			case "cd":
				newPath, err := updateCurrentPath(splittedString[1], currentDir)
				if err != nil {
					fmt.Println(err)
					continue
				}
				currentDir = newPath
			case "pwd":
				fmt.Println(currentDir)
			case "kill":
				err = killProcess(splittedString[1])
			case "/exit":
				return
			case "echo":
				for i, v := range splittedString {
					if i != 0 {
						fmt.Printf("%s ", v)
					}
				}
				fmt.Println("")
			default:
				forkExec(splittedString)
			}
		}
	}
}

func showProcesses() error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, v := range processes {
		name, err := v.Name()
		if err != nil {
			continue
		}
		fmt.Println(name)
	}
	return nil
}

func killProcess(name string) error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			continue
		}
		if n == name {
			return p.Kill()
		}
	}
	return fmt.Errorf("process not found")
}

func updateCurrentPath(args, currentPath string) (string, error) {
	if args == ".." {
		return filepath.Dir(currentPath), nil
	}
	s, err := os.Stat(args)
	if err != nil {
		return "", fmt.Errorf("Can't find new path")
	}
	if s.IsDir() {
		return args, err
	}
	return "", fmt.Errorf("Can't find new path")
}

func forkExec(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Ошибка при запуске команды:", err)
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Ошибка при выполнении команды:", err)
	}
}
