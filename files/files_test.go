package files

import (
	"testing"
	"os"
)

func TestHostExtractor(t *testing.T) {
	dir, err := os.Getwd()
	hosts, err := FileLinesExtractor(dir + "/../host.lst")
	if err != nil {
		t.Error(err.Error())
	}
	expectHost := [] string{"example.com:22", "example2.com:2222"}

	for index, host := range hosts {
		if host != expectHost[index] {
			t.Error("File incoherency")
		}
	}
}
