package input

import (
	"testing"
	"strings"
	"bufio"
	"github.com/stretchr/testify/assert"
)

func TestParseHosts(t *testing.T) {

	inputFixture := bufio.NewReader(strings.NewReader("host1:1111\nhost2:2222\nhost3:3333"))
	hosts, err := ParseHosts(inputFixture)
	assert.Equal(t, nil, err, "Unexpected error while testing parse hosts")
	expectedHosts := []string{"host1:1111", "host2:2222", "host3:3333"}
	assert.Equal(t, hosts, expectedHosts, "Mismatch parsing and comparing hosts")
}

func TestHostsFormatIsChecked(t *testing.T) {

	inputFixture := bufio.NewReader(strings.NewReader("host1.com:1111\nhost2\nhost3.net:3333"))
	hosts, err := ParseHosts(inputFixture)
	assert.Equal(t, []string{}, hosts)
	assert.Equal(t, "Host 'host2' must accomplish format, host:port", err.Error(), "Must check for port.")
}
