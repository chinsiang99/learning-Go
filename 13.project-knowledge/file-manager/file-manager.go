package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file", err)
		return nil, errors.New("error opening file")
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	errScanner := scanner.Err()

	if errScanner != nil {
		fmt.Println("Error reading file", err)
		file.Close()
		return nil, errors.New("error reading file")
	}
	file.Close()
	return lines, nil
}
