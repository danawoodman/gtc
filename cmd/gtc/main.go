package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/danawoodman/gtc/internal"
)

func main() {
	args := []string{"test"}
	args = append(args, os.Args[1:]...)

	cmd := exec.Command("go", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating StderrPipe for Cmd", err)
		os.Exit(1)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting Cmd", err)
		os.Exit(1)
	}

	f := internal.NewFormatter()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(f.Format(line))
	}

	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(f.Format(line))
	}

	if err := cmd.Wait(); err != nil {
		os.Exit(1)
	}
}
