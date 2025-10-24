package main

import (
	"fmt"
	"log"

	"github.com/shaheerkt123/BluePasser/internal/blueutils"
	"github.com/shaheerkt123/BluePasser/internal/crypto"
)

func main() {
	log.Println("Starting client...")
	err := blueutils.ScanForCredentials(func(creds blueutils.WifiCredentials) {
		log.Println("Received credentials")
		decryptedPassword, err := crypto.Decrypt(creds.Password)
		if err != nil {
			log.Printf("Failed to decrypt password: %v", err)
			return
		}
		fmt.Printf("SSID: %s\nPassword: %s\n", creds.SSID, decryptedPassword)
	})
	if err != nil {
		log.Fatalf("Failed to scan for credentials: %v", err)
	}
}