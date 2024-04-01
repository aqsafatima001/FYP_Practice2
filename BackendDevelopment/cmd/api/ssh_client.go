package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func (app *application) createUserOnCentOS(username, password string) error {
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
	sshClient, err := ssh.Dial("tcp", "192.168.100.133:22", sshConfig)
	if err != nil {
		log.Println("Error connecting to CentOS VM:", err)
		return err
	}
	defer sshClient.Close()

	log.Println("SSH connection established")

	session, err := sshClient.NewSession()
	if err != nil {
		log.Println("Error creating SSH session:", err)
		return err
	}
	defer session.Close()

	log.Println("SSH session created")

	// Replace "new_username" with the desired username and "new_password" with the desired password
	cmd := fmt.Sprintf("sudo useradd -m %s && echo '%s:%s' | sudo chpasswd", username, username, password)
	log.Println("Executing command:", cmd)
	if err := session.Run(cmd); err != nil {
		log.Println("Error executing command:", err)
		return err
	}

	log.Println("User created successfully")

	return nil
}

// func (app *application) createUserInCentOS(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Starting user creation process")

// 	os.Setenv("PATH", "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/bin:/usr/sbin:/bin:/sbin:/c/Windows/System32/OpenSSH")

// 	// Replace these values with your CentOS VM SSH credentials
// 	sshConfig := &ssh.ClientConfig{
// 		User: "root",
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password("aqsafatima"),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}

// 	// Replace "your_centos_vm_ip" with the actual IP address of your CentOS VM
// 	log.Println("Connecting to CentOS VM via SSH")
// 	sshClient, err := ssh.Dial("tcp", "192.168.100.133:22", sshConfig)
// 	if err != nil {
// 		log.Println("Error connecting to CentOS VM:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer sshClient.Close()

// 	log.Println("SSH connection established")

// 	session, err := sshClient.NewSession()
// 	if err != nil {
// 		log.Println("Error creating SSH session:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer session.Close()

// 	log.Println("SSH session created")

// 	// Replace "new_username" with the desired username and "new_password" with the desired password
// 	cmd := "sudo useradd -m fatima && echo 'fatima:aqsafatima' | sudo chpasswd"
// 	log.Println("Executing command:", cmd)
// 	if err := session.Run(cmd); err != nil {
// 		log.Println("Error executing command:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	log.Println("User created successfully")

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("User created successfully"))
// }
