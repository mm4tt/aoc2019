package util

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadString(filepath string) (string, error) {
	inputFile, err := os.Open(os.ExpandEnv(filepath))
	if err != nil {
		return "", err
	}
	buffer, err := ioutil.ReadAll(inputFile)
	return string(buffer), err
}

func ReadLines(relativePath string) (<-chan string, error) {
	inputFile, err := os.Open(os.ExpandEnv(filepath.Join("$GOPATH/src/github.com/mm4tt/aoc2019/", relativePath)))
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

func ReadAllLines(relativePath string) ([]string, error) {
	inputFile, err := os.Open(os.ExpandEnv(filepath.Join("$GOPATH/src/github.com/mm4tt/aoc2019/", relativePath)))
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(inputFile)
	ret := []string{}
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret, nil
}
