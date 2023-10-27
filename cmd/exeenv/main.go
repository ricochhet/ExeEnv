package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	File = "exe_env.txt"
)

func main() {
	env, exe, err := readEnvFormat(File)
	if err != nil {
		fmt.Printf("Error reading environment variables from \"%s\": %v\n", env, err)
		return
	}

	for key, value := range env {
		os.Setenv(key, value)
	}

	exe = strings.TrimRight(exe, "\r")
	exe = strings.TrimRight(exe, "\n")
	_, err = os.Stat(exe)
	if err != nil {
		fmt.Printf("The executable file does not exist: %s\n", exe)
		return

	}

	cmd := exec.Command(exe)
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting process \"%s\", %v\n", exe, err)
		return
	}

	cmd.Wait()
}

func readEnvFormat(path string) (map[string]string, string, error) {
	env := make(map[string]string)
	var exeName string

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}

	lines := strings.Split(string(file), "\n")
	if len(lines) == 0 {
		return nil, "", fmt.Errorf("Empty file")
	}

	exeName = lines[0]
	for _, line := range lines[1:] {
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			key := parts[0]
			value := parts[1]
			env[key] = value
		}
	}

	return env, exeName, nil
}
