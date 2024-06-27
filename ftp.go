package main

import (
	"fmt"
	"log"

	"github.com/jlaffaye/ftp"
)

func connectFTP(host string, port int, user, password string) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := ftp.Dial(addr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	err = conn.Login(user, password)
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}
	defer conn.Logout()

	entries, err := conn.List("/")
	if err != nil {
		log.Fatalf("Failed to list: %v", err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Name)
	}
}
