package internal

import (
	"bufio"
	"fmt"
	"os/exec"
)

func NewCmd(args []string) error {
	args = append([]string{"test"}, args...)
	cmd := exec.Command("go", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating StdoutPipe for gtc: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error creating StderrPipe for gtc: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting gtc: %w", err)
	}

	f := NewFormatter()

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
		return err
	}

	return nil
}
