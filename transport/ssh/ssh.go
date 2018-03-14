package ssh

import (
	"io/ioutil"
	"log"
	"golang.org/x/crypto/ssh"
	"os/user"
	"bytes"
	"net"
	"time"
	"errors"
)

type SSHAdapter interface {
	Connect() error
	Close()
	Run(command string) string
	config() *sshConfiguration
}

type sshAdapter struct {
	Configuration *sshConfiguration
	Client        *SSHClient
	Session       *session
	server        string
}

type sshConfiguration struct {
	*ssh.ClientConfig
}

func (adapter *sshAdapter) config() *sshConfiguration {
	userInfo := adapter.getSystemUserInfo()
	hostKeyCallback := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	auth := adapter.publicKeyAuthBuilder(userInfo.HomeDir + "/.ssh/golang_rsa")
	authMethods := []ssh.AuthMethod{auth}
	config := &ssh.ClientConfig{
		User:            userInfo.Username,
		Auth:            authMethods,
		HostKeyCallback: hostKeyCallback,
		Timeout:         time.Second * 1,
	}
	adapter.Configuration = &sshConfiguration{config}
	return adapter.Configuration
}

func (adapter *sshAdapter) getSystemUserInfo() *user.User {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Cannot gather default system usr.")
	}
	return usr
}

func (adapter *sshAdapter) publicKeyAuthBuilder(file string) ssh.AuthMethod {

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

func (adapter *sshAdapter) Connect() error {
	config := adapter.config()
	conn, err := ssh.Dial("tcp", adapter.server, config.ClientConfig)
	if err != nil {
		return errors.New("Cannot stablish ssh client (dial phase)." + err.Error())
	}
	adapter.Client = &SSHClient{conn}
	session, err := adapter.createSession()
	adapter.Session = session
	if err != nil {
		return err
	}
	return nil
}

func (adapter *sshAdapter) Close() {
	adapter.Client.Close()
	adapter.Session.Close()
}

func (adapter *sshAdapter) Run(command string) string {
	var output bytes.Buffer
	adapter.Session.SetStdout(&output)
	adapter.Session.Run(command)
	return adapter.server + " says :\n" + output.String()

}

type SSHClient struct {
	*ssh.Client
}

func (adapter *sshAdapter) createSession() (*session, error) {
	newSession, err := adapter.Client.NewSession()
	if err != nil {
		return nil, errors.New("Error getting ssh session." + err.Error())
	}
	return &session{newSession}, nil
}

type Session interface {
	Close()
	SetStdout(output *bytes.Buffer)
	Run(command string) string
}

type session struct {
	*ssh.Session
}

func (s *session) Close() {
	s.Session.Close()
}

func (s *session) SetStdout(output *bytes.Buffer) {
	s.Session.Stdout = output
}

func (s *session) Run(command string) {
	s.Session.Run(command)
}

func New(server string) SSHAdapter {
	return &sshAdapter{server: server}
}
