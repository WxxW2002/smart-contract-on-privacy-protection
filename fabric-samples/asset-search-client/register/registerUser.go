package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func main() {
	// Configure paths
	walletPath := "./wallet"
	certPath := "/home/wxxw/Desktop/smart-contract-on-privacy-protection/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/cert.pem"
	keyPath := "/home/wxxw/Desktop/smart-contract-on-privacy-protection/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/"

	// Initialize wallet
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	// Check if the user identity already exists
	if wallet.Exists("appUser") {
		fmt.Println("User identity already exists in the wallet")
		return
	}

	// Read user cert and key
	identity := gateway.NewX509Identity("Org1MSP", readFile(certPath), readPrivateKey(keyPath))
	err = wallet.Put("appUser", identity)
	if err != nil {
		log.Fatalf("Failed to put user identity into wallet: %v", err)
	}

	fmt.Println("User identity added to the wallet")
}

func readFile(path string) string {
	// Read content of the certificate file
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return string(contents)
}

func readPrivateKey(dirPath string) string {
	// Find the private key file in the keystore directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Failed to read key directory: %v", dirPath)
	}
	if len(files) == 0 {
		log.Fatalf("No private key file found in directory: %v", dirPath)
	}
	privateKeyPath := fmt.Sprintf("%s/%s", dirPath, files[0].Name())
	return readFile(privateKeyPath)
}
