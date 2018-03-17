package input

import (
	"flag"
	"os/user"
	"log"
)

type Arguments struct {
	User           string
	PrivateKeyPath string
	MaxConcurrency int
	Command        string
}

func getSystemUserInfo() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Cannot gather default system usr.")
	}
	return usr
}

func ParseArgs() *Arguments {

	var userName string
	var privateKeyPath string
	var maxConcurrency int

	userInfo := getSystemUserInfo()

	flag.StringVar(&userName, "u", userInfo.Username, "Specify a different user for connections.")
	flag.StringVar(&privateKeyPath, "k", userInfo.HomeDir+"/.ssh/id_rsa", "Specify absolute path of your custom private key.")
	flag.IntVar(&maxConcurrency, "c", 10, "Set the max concurrency for the process.")
	flag.Parse()
	command := flag.Arg(0)

	return &Arguments{
		userName,
		privateKeyPath,
		maxConcurrency,
		command,
	}
}
