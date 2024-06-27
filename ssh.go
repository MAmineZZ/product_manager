package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func connectSSH(user, password, host string, port int) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	output, err := session.Output("uname -a")
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
	fmt.Println(string(output))
}
