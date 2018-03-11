package files

import (
	"os"
	"log"
	"bufio"
	"errors"
)

var fileError = errors.New("File error")

func FileLinesExtractor(configFile string) ([]string, error) {
	data, err := os.Open(configFile)
	if err != nil {
		log.Print(fileError.Error())
		return nil, fileError
	}
	var lines []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}