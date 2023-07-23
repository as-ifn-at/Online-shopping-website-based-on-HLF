/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/chaincode"
)

func main() {
	// assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	chaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{}, &chaincode.UserRegistrationContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	// chaincode, err := contractapi.NewChaincode(&CashbackTokenContract{}, &UserRegistrationContract{})

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}