package input

import (
	"bufio"
	"regexp"
	"errors"
)

func ParseHosts(reader *bufio.Reader) ([]string, error) {

	scanner := bufio.NewScanner(reader)
	var hosts [] string
	for scanner.Scan() {
		host := scanner.Text()
		err := checkHostFormat(host)
		if err != nil {
			return []string{}, err
		}
		hosts = append(hosts, host)
	}
	return hosts, nil
}

func checkHostFormat(host string) error {
	const regex = "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9]):([0-9]){1,5}$"
	result, err := regexp.Match(regex, []byte(host))
	if err != nil {
		return errors.New("Host bad parsing againts regex: " + regex)
	}
	if !result {
		return errors.New("Host '" + host + "' must accomplish format, host:port")
	}
	return nil
}
