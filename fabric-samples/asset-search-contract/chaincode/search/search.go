package search

import (
	"encoding/json"
	"fmt"

	"asset-search-contract/chaincode/encryption"
	"asset-search-contract/chaincode/models"

	"github.com/hyperledger/fabric-chaincode-go/v2/shim"
)

func SearchAssetsByColor(stub shim.ChaincodeStubInterface, color string) ([]*models.Asset, error) {
	colorIndex := encryption.BuildIndex(color)

	query := fmt.Sprintf("{\"selector\":{\"EncryptedColor\":\"%s\"}}", colorIndex)

	resultsIterator, err := stub.GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*models.Asset

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset models.Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}

		assets = append(assets, &asset)
	}

	return assets, nil
}

func SearchAssetsByOwner(stub shim.ChaincodeStubInterface, owner string) ([]*models.Asset, error) {
	ownerIndex := encryption.BuildIndex(owner)

	query := fmt.Sprintf("{\"selector\":{\"EncryptedOwner\":\"%s\"}}", ownerIndex)

	resultsIterator, err := stub.GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*models.Asset

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset models.Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}

		assets = append(assets, &asset)
	}

	return assets, nil
}
