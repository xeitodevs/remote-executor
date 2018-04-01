package input

import (
	"os"
)

func StdinChecker(input *os.File) {
	inputData, err := input.Stat()
	if err != nil {
		panic(err)
	}
	if inputData.Mode()&os.ModeNamedPipe == 0 {
		panic("We need pipe hosts from file for this process.")
	}

}
