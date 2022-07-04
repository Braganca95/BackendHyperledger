/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	projetoIvContract := new(ProjetoIvContract)
	projetoIvContract.Info.Version = "0.0.1"
	projetoIvContract.Info.Description = "My Smart Contract"
	projetoIvContract.Info.License = new(metadata.LicenseMetadata)
	projetoIvContract.Info.License.Name = "Apache-2.0"
	projetoIvContract.Info.Contact = new(metadata.ContactMetadata)
	projetoIvContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(projetoIvContract)
	chaincode.Info.Title = "braganca95 chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from ProjetoIvContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
