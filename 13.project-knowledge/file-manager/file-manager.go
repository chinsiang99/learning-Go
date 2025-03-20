package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
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

func (fm *FileManager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("error creating file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(data)
}

func NewFileManager(inputFilePath, outputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}
