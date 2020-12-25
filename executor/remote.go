package executor

import (
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

// The function used to open files when necessary for matching
// Allows the file IO to be overridden during tests
var ReadFile = ioutil.ReadFile

type SSHExecutor struct {
	Address      string
	User         string
	IdentityFile string
	InitOpts     InitOpts
	ClientConfig *ssh.ClientConfig
}

func NewSSHExecutor(address string, user string, identityFile string, opts InitOpts) SSHExecutor {
	e := SSHExecutor{
		Address:      address,
		User:         user,
		IdentityFile: identityFile,
		InitOpts:     opts,
	}

	authMethods := []ssh.AuthMethod{}

	signer := e.createSigner()
	authMethods = append(authMethods, ssh.PublicKeys(signer))

	e.ClientConfig = &ssh.ClientConfig{
		User: e.User,
		Auth: authMethods,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	return e
}

func (e SSHExecutor) createSigner() ssh.Signer {
	buffer, err := ReadFile(e.IdentityFile)
	if err != nil {
		log.Fatal(err)
	}
	signer, err := ssh.ParsePrivateKey(buffer)
	return signer
}
