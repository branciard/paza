
package main

import (
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)



//==============================================================================================================================
//	 ASSET Status
//==============================================================================================================================
const   ASSET_STATUS_DESTROYED  		=  0
const   ASSET_STATUS_OWNED  			=  1
const   ASSET_STATUS_AVAILABLE   		=  2
const   ASSET_STATUS_RESERVED  			=  3
const   ASSET_STATUS_OWNERSHIP_TRADING  	=  4


//==============================================================================================================================
//	 asset transactions List
//==============================================================================================================================
const   DESTROY_ASSET_TX		  	=  0
const   CREATE_ASSET_TX 	 		=  1
const   UPDATE_ASSET_TX	 	   		=  2
const   OFFER_ASSET_TX  			=  3
const   CANCEL_OFFER_ASSET_TX  			=  4
const   RESERVE_ASSET_TX  			=  5
const   CANCEL_RESERVE_ASSET_REQUEST_TX  	=  6
const   CANCEL_RESERVE_ASSET_REQUESTED_TX  	=  7
const   REQUEST_FOR_ASSET_OWNERSHIP_TX  	=  8
const   ACCEPT_ASSET_OWNERSHIP_REQUEST_TX  	=  9
const   REFUSE_ASSET_OWNERSHIP_REQUEST_TX  	=  10
const   GIVE_ASSET_OWNERSHIP_TX  		=  11
const   ACCEPT_ASSET_OWNERSHIP_GIVEN_TX  	=  12
const   REFUSE_ASSET_OWNERSHIP_GIVEN_TX  	=  13

// ASSET TABLE
// consts associated with chaincode table
const (
	tableColumn       = "ASSETS_TABLE"
	columnASSET_ID   = "ASSET_ID"
	columnASSET_STATUS   = "ASSET_STATUS"
	columnASSET_OWNER_ECERT   = "ASSET_OWNER_ECERT"
	columnASSET_OWNER_TIMESTAMP= "ASSET_OWNER_TIMESTAMP"
	columnASSET_OWNER_HIST= "ASSET_OWNER_HIST"
	columnASSET_DESC= "ASSET_DESC"
	columnASSET_DESC_TIMESTAMP= "ASSET_DESC_TIMESTAMP"
	columnASSET_DESC_HIST= "ASSET_DESC_HIST"
	columnASSET_TX_TCERT= "ASSET_TX_TCERT"
	columnASSET_TX= "ASSET_TX"
	columnASSET_TX_TIMESTAMP= "SSET_TX_TIMESTAMP"
	columnASSET_TX_HIST = "ASSET_TX_HIST"
)

//AssetHandler provides APIs used to perform operations
type assetHandler struct {
}

// NewAssetHandler create a new reference to AssetHandler
func NewAssetHandler() *assetHandler {
	return &assetHandler{}
}

// createTable initiates a new asset table in the chaincode state
// stub: chaincodestub
func (t *assetHandler) createTable(stub shim.ChaincodeStubInterface) error {

	// Create asset depository table
	return stub.CreateTable(tableColumn, []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: columnASSET_ID, Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: columnASSET_STATUS, Type: shim.ColumnDefinition_UINT32, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_OWNER_ECERT, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_OWNER_TIMESTAMP, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_OWNER_HIST, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_DESC, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_DESC_TIMESTAMP, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_DESC_HIST, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_TX_TCERT, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_TX, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_TX_TIMESTAMP, Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: columnASSET_TX_HIST, Type: shim.ColumnDefinition_STRING, Key: false},

	})

}



// queryTable returns the record row matching a correponding asset ID on the chaincode state table
// stub: chaincodestub
// assetID: assetID
func (t *depositoryHandler) queryTable(stub shim.ChaincodeStubInterface, assetID string) (shim.Row, error) {

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: assetID}}
	columns = append(columns, col1)

	return stub.GetRow(tableColumn, columns)
}