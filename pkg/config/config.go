package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv(filePath string) error {
	// Open the .env file 
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open .env file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line[0] == '#' {
			continue
		}

 		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			err := os.Setenv(parts[0], parts[1])
			if err != nil {
				return fmt.Errorf("could not set environment variable: %v", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .env file: %v", err)
	}

	return nil
}
