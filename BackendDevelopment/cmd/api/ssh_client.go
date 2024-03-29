package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/ssh"
)

func (app *application) createUserInCentOS(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting user creation process")

	os.Setenv("PATH", "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/bin:/usr/sbin:/bin:/sbin:/c/Windows/System32/OpenSSH")

	// Replace these values with your CentOS VM SSH credentials
	sshConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("aqsafatima"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Replace "your_centos_vm_ip" with the actual IP address of your CentOS VM
	log.Println("Connecting to CentOS VM via SSH")
	sshClient, err := ssh.Dial("tcp", "192.168.206.41:22", sshConfig)
	if err != nil {
		log.Println("Error connecting to CentOS VM:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer sshClient.Close()

	log.Println("SSH connection established")

	session, err := sshClient.NewSession()
	if err != nil {
		log.Println("Error creating SSH session:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()

	log.Println("SSH session created")

	// Replace "new_username" with the desired username and "new_password" with the desired password
	cmd := "sudo useradd -m fatima && echo 'fatima:aqsafatima' | sudo chpasswd"
	log.Println("Executing command:", cmd)
	if err := session.Run(cmd); err != nil {
		log.Println("Error executing command:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("User created successfully")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created successfully"))
}

// // ssh_client.go

// package main

// import (
// 	"io/ioutil"

// 	"golang.org/x/crypto/ssh"
// )

// // SSHClient represents an SSH client
// type SSHClient struct {
// 	Config *ssh.ClientConfig
// 	Host   string
// 	Port   string
// }

// // NewSSHClient creates a new SSH client
// func NewSSHClient(user, privateKeyPath, host, port string) (*SSHClient, error) {
// 	privateKey, err := ioutil.ReadFile(privateKeyPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	signer, err := ssh.ParsePrivateKey(privateKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	clientConfig := &ssh.ClientConfig{
// 		User: user,
// 		Auth: []ssh.AuthMethod{
// 			ssh.PublicKeys(signer),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}

// 	return &SSHClient{
// 		Config: clientConfig,
// 		Host:   host,
// 		Port:   port,
// 	}, nil
// }

// // ExecuteCommand executes a command on the remote server
// func (c *SSHClient) ExecuteCommand(cmd string) ([]byte, error) {
// 	conn, err := ssh.Dial("tcp", c.Host+":"+c.Port, c.Config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()

// 	session, err := conn.NewSession()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer session.Close()

// 	output, err := session.CombinedOutput(cmd)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return output, nil
// }
