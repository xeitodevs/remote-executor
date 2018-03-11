package transport

import (
	"io/ioutil"
	"log"
	"golang.org/x/crypto/ssh"
	"os/user"
	"bytes"
	"net"
)

func ExecuteCommand(command string, server string, channel chan<- string) {

	conn, err := ssh.Dial("tcp", server, getConfig())
	if err != nil {
		channel <- "Cannot stablish ssh connection (dial phase)." + err.Error()
		return
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		channel <- "Error getting ssh session." + err.Error()
		return
	}
	defer session.Close()

	var output bytes.Buffer
	session.Stdout = &output
	session.Run(command)
	channel <- parseResponse(server, output.String())
}

func parseResponse(server string, message string) string {
	return server + " says :\n" + message
}

func getConfig() *ssh.ClientConfig {
	userInfo := getSystemUserInfo()
	config := &ssh.ClientConfig{
		User: userInfo.Username,
		Auth: []ssh.AuthMethod{
			publicKey(userInfo.HomeDir + "/.ssh/id_rsa"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	return config
}

func getSystemUserInfo() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Cannot gather default system usr.")
	}
	return usr
}

func publicKey(file string) ssh.AuthMethod {

	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Cannot read ssh file.")
	}
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		log.Fatal("Incorrect parse of private key.")
	}
	return ssh.PublicKeys(key)
}
