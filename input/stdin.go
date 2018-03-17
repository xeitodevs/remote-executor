package input

import (
	"os"
	"bufio"
)

func GetHosts(input *os.File) []string {
	inputData, err := input.Stat()
	if err != nil {
		panic(err)
	}
	if inputData.Mode()&os.ModeNamedPipe == 0 {
		panic("We need pipe hosts from file for this process.")
	}
	scanner := bufio.NewScanner(input)
	var hosts [] string
	for scanner.Scan() {
		hosts = append(hosts, scanner.Text())
	}
	return hosts
}
