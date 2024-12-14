package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Asset struct {
	ID             string `json:"ID"`
	Color          string `json:"Color"`
	Size           int    `json:"Size"`
	Owner          string `json:"Owner"`
	AppraisedValue int    `json:"AppraisedValue"`
	EncryptedColor string `json:"EncryptedColor"`
	EncryptedOwner string `json:"EncryptedOwner"`
}

func main() {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		fmt.Println("An identity for the client user \"appUser\" not exists in the wallet")
		os.Exit(1)
	}

	ccpPath := "./connection.yaml"

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	contract := network.GetContract("asset-transfer-basic")

	err = createAsset(contract, "asset7", "blue", 20, "Tom", 900)
	if err != nil {
		fmt.Printf("Failed to create asset: %v\n", err)
		os.Exit(1)
	}

	err = searchAssetsByColor(contract, "blue")
	if err != nil {
		fmt.Printf("Failed to search assets: %v\n", err)
		os.Exit(1)
	}
}

func createAsset(contract *gateway.Contract, assetID string, color string, size int, owner string, appraisedValue int) error {
	fmt.Printf("Submit Transaction: CreateAsset, %s with color %s, size %d, owner %s, appraisedValue %d\n", assetID, color, size, owner, appraisedValue)

	_, err := contract.SubmitTransaction("CreateAsset", assetID, color, fmt.Sprint(size), owner, fmt.Sprint(appraisedValue))
	if err != nil {
		return fmt.Errorf("failed to submit transaction: %v", err)
	}

	fmt.Println("*** Transaction committed successfully")
	return nil
}

func searchAssetsByColor(contract *gateway.Contract, color string) error {
	fmt.Printf("Evaluate Transaction: SearchAssetsByColor, color %s\n", color)

	result, err := contract.EvaluateTransaction("SearchAssetsByColor", color)
	if err != nil {
		return fmt.Errorf("failed to evaluate transaction: %v", err)
	}

	var assets []*Asset
	err = json.Unmarshal(result, &assets)
	if err != nil {
		return fmt.Errorf("failed to unmarshal assets: %v", err)
	}

	fmt.Println("*** Result:")
	for _, asset := range assets {
		fmt.Printf("Asset %s: color %s, size %d, owner %s, appraisedValue %d\n", asset.ID, asset.Color, asset.Size, asset.Owner, asset.AppraisedValue)
	}

	return nil
}
