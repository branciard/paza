


package main


import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type AssetChaincode struct {
}
var assetHandler = NewAssetHandler()

func (t *AssetChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Printf("Asset Init called")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	return nil, assetHandler.createTable(stub)

}

// Transaction makes payment of X units from A to B
func (t *AssetChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {



	return nil, nil
}

// Deletes an entity from state
func (t *AssetChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *AssetChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {


	return Avalbytes, nil
}
func main() {
	err := shim.Start(new(AssetChaincode))
	if err != nil {
		fmt.Printf("Error starting Asset chaincode: %s", err)
	}
}