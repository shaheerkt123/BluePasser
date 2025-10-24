package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shaheerkt123/BluePasser/internal/blueutils"
	"github.com/shaheerkt123/BluePasser/internal/crypto"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run cmd/host/main.go <ssid> <password>")
		return
	}

	ssid := os.Args[1]
	password := os.Args[2]

	encryptedPassword, err := crypto.Encrypt(password)
	if err != nil {
		log.Fatalf("Failed to encrypt password: %v", err)
	}

	log.Println("Starting host...")
	err = blueutils.StartBroadcasting(ssid, encryptedPassword)
	if err != nil {
		log.Fatalf("Failed to start broadcasting: %v", err)
	}
}