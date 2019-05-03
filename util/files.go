package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func hasFileFormat(filepath string, format string) bool {
	subStrings := strings.Split(filepath, ".")
	ending := subStrings[len(subStrings)-1]

	if ending != format {
		return false
	}

	return true
}

func readFile(filepath string) ([]string, error) {

	var output []string

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return output, nil
}

func WriteFile(filepath string, input []string) error {

	file, err := os.Create(filepath)
	if err != nil {
		file.Close()
		return err
	}

	for _, login := range input {
		_, err = fmt.Fprintln(file, login)
		if err != nil {
			return err
		}
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func ReadInputFiles(filepathes []string) ([]string, error) {

	var output []string

	for _, filepath := range filepathes {

		if !hasFileFormat(filepath, "txt") {
			return nil, errors.New("wrong file format")
		}

		content, err := readFile(filepath)
		if err != nil {
			return nil, err
		}

		for _, line := range content {
			if line != "" && !Contains(output, line) {
				output = append(output, line)
			}
		}
	}

	return output, nil
}
