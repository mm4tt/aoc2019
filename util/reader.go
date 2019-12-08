package util

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ReadString(filepath string) (string, error) {
	inputFile, err := os.Open(os.ExpandEnv(filepath))
	if err != nil {
		return "", err
	}
	buffer, err := ioutil.ReadAll(inputFile)
	return string(buffer), err
}

func ReadLines(filepath string) (<-chan string, error) {
	inputFile, err := os.Open(os.ExpandEnv(filepath))
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(inputFile)
	output := make(chan string)
	go func() {
		for scanner.Scan() {
			output <- scanner.Text()
		}
		close(output)
	}()
	return output, nil
}
