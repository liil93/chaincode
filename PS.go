package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type PS struct { // Petsitting chaincode
}

type TradeRec struct { // Trade record (KEY: PSID#CSID#Trade complete time)
	PSID string // Petsitter ID
	CSID string // Consumer ID
	TS   string // Transaction start time
	TE   string // Transaction end time
	TC   string // Transaction complete time
	TA   string // Transaction amount
	TH   string // Transaction history
}

var _CCstr string

type CityCode struct {
	R103 string // Region code 103 (nowongu)
	R104 string // Region code 104 (gangnamgu)
	R105 string // Region code 105 (zongrogu)
}

type UserInfo struct { // User information (KEY: User email)
	PW string // User Password
	PN string // Pet number
	CC string // City code
	AP string // Companion animal
	// AllHomeAsset []HomeAsset
	// AllPetAsset []PetAsset
}
type HomeAsset struct { // Information about home (KEY: User email#home)
	Address  string // Address about home
	HomeType string // House type
	Room     string // Room count
	Area     string // The floor of my asset
	Elevator string // Presence of elevator
	Parking  string // Parking applicability
}
type PetAsset struct { // Information about pet (KEY: User email#pet)
	Name   string // Pet name
	Birth  string // Pet birth
	Gender string // Pet gender
	Kind   string // Pet kind
	Size   string // Pet size (S: ~5kg, M: 5~10, L: 10~)
	NS     string // Neutralization surgery
	Vac    string // Vaccine check
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(PS))
	if err != nil {
		fmt.Printf("Error starting PS chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *PS) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, errors.New("[INIT] Incorrect number of arguments. Expecting 0")
	}
	_CCstr = "_CityCodeStruct"
	cityCode := CityCode{}
	cityCode.R103 = "/"
	cityCode.R104 = "/"
	cityCode.R105 = "/"

	jsonAsBytes, _ := json.Marshal(cityCode)
	stub.PutState(_CCstr, jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *PS) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("[INVOKE] invoke is running: " + function)

	if function == "user_insert" {
		return t.UInsert(stub, args)
	} else if function == "home_insert" {
		return t.HInsert(stub, args)
	} else if function == "pet_insert" {
		return t.PInsert(stub, args)
	} else if function == "home_delete" {
		return t.HDelete(stub, args)
	} else if function == "pet_delete" {
		return t.PDelete(stub, args)
	}
	// ....adding....
	// user_change (PW, PN, AP)           ### 0%
	// home_change (All home info) + CC   ### 0%
	// pet_change (Size, NS, Vac)         ### 0%
	// home_delete + CC, AP               ### 70% (Region change...)
	// pet_delete + PN                    ### 100%

	fmt.Println("[INVOKE] invoke did not find func: " + function)
	return nil, errors.New("[INVOKE] Received unknown function invocation: " + function)
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *PS) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("[QUERY] query is running: " + function)
	if function == "user_read" {
		return t.URead(stub, args)
	} else if function == "home_read" {
		return t.HRead(stub, args)
	} else if function == "pet_read" {
		return t.PRead(stub, args)
	} else if function == "city_search" {
		return t.CSearch(stub, args)
	}
	fmt.Println("[QUERY] query did not find func: " + function) //error
	return nil, errors.New("[QUERY] Received unknown function query: " + function)
}

// ============================================================================================================================
// UInsert - insert user information
// ============================================================================================================================
func (t *PS) UInsert(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("[USER INSSERT] Incorrect number of arguments. Expecting 2")
	}
	conf, _ := stub.GetState(args[0])
	if conf != nil {
		return nil, errors.New("[USER INSSERT] Already exist user")
	}
	userID := args[0]
	pw := args[1]

	userInfo := UserInfo{}
	userInfo.PW = pw
	userInfo.PN = "0"
	userInfo.CC = "0"
	userInfo.AP = "0"
	jsonAsBytes, _ := json.Marshal(userInfo)
	stub.PutState(userID, jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// HInsert - insert home information
// ============================================================================================================================
func (t *PS) HInsert(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 8 {
		return nil, errors.New("[HOME INSSERT] Incorrect number of arguments. Expecting 8")
	} // home_insert, 103, ~~~HOME INFO~~~
	confUser, _ := stub.GetState(args[0])
	if confUser == nil {
		return nil, errors.New("[HOME INSSERT] Not exist userID")
	}
	userInfo := UserInfo{}
	json.Unmarshal(confUser, &userInfo)
	if userInfo.CC != "0" {
		return nil, errors.New("[HOME INSSERT] Already exist home")
	}
	homeAsset := HomeAsset{}
	cityCode := CityCode{}
	confCC, _ := stub.GetState(_CCstr)
	json.Unmarshal(confCC, &cityCode)

	userID := args[0]
	cc := args[1]
	address := args[2]
	hometype := args[3]
	room := args[4]
	area := args[5]
	elevator := args[6]
	parking := args[7]

	userInfo.CC = cc
	if cc == "R103" {
		cityCode.R103 = cityCode.R103 + userID + "/"
	} else if cc == "R104" {
		cityCode.R104 = cityCode.R104 + userID + "/"
	} else if cc == "R105" {
		cityCode.R105 = cityCode.R105 + userID + "/"
	}
	homeAsset.Address = address
	homeAsset.HomeType = hometype
	homeAsset.Room = room
	homeAsset.Area = area
	homeAsset.Elevator = elevator
	homeAsset.Parking = parking

	jsonAsBytesU, _ := json.Marshal(userInfo)
	jsonAsBytesC, _ := json.Marshal(cityCode)
	jsonAsBytesH, _ := json.Marshal(homeAsset)
	stub.PutState(userID, jsonAsBytesU)
	stub.PutState(_CCstr, jsonAsBytesC)
	stub.PutState(userID+"#home", jsonAsBytesH)
	return nil, nil
}

// ============================================================================================================================
// PInsert - insert pet information
// ============================================================================================================================
func (t *PS) PInsert(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 8 {
		return nil, errors.New("[PET INSSERT] Incorrect number of arguments. Expecting 8")
	}
	confUser, _ := stub.GetState(args[0])
	if confUser == nil {
		return nil, errors.New("[PET INSSERT] Not exist userID")
	}
	userInfo := UserInfo{}
	json.Unmarshal(confUser, &userInfo)
	if userInfo.PN != "0" {
		return nil, errors.New("[PET INSSERT] Already exist pet")
	}
	petAsset := PetAsset{}

	userID := args[0]
	name := args[1]
	birth := args[2]
	gender := args[3]
	kind := args[4]
	size := args[5]
	ns := args[6]
	vac := args[7]

	userInfo.PN = "1"
	petAsset.Name = name
	petAsset.Birth = birth
	petAsset.Gender = gender
	petAsset.Kind = kind
	petAsset.Size = size
	petAsset.NS = ns
	petAsset.Vac = vac
	jsonAsBytesU, _ := json.Marshal(userInfo)
	jsonAsBytesP, _ := json.Marshal(petAsset)
	stub.PutState(userID, jsonAsBytesU)
	stub.PutState(userID+"#pet", jsonAsBytesP)
	return nil, nil
}

// ============================================================================================================================
// URead - read user information
// ============================================================================================================================
func (t *PS) URead(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("[USER QUERY] Incorrect number of arguments. Expecting 1")
	}
	key := args[0]
	valAsbytes, _ := stub.GetState(key) //get the pet information from chaincode state
	if valAsbytes == nil {
		return []byte("[HOME QUERY] Not exist userID"), errors.New("[HOME QUERY] Not exist userID")
	}
	return valAsbytes, nil
}

// ============================================================================================================================
// HRead - read home information
// ============================================================================================================================
func (t *PS) HRead(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("[HOME QUERY] Incorrect number of arguments. Expecting 1")
	}
	key := args[0]
	conf, _ := stub.GetState(key)
	if conf == nil {
		return []byte("[HOME QUERY] Not exist userID"), errors.New("[HOME QUERY] Not exist userID")
	}
	userInfo := UserInfo{}
	json.Unmarshal(conf, &userInfo)
	if userInfo.CC == "0" {
		return []byte("[HOME QUERY] Not exist home information"), errors.New("[HOME QUERY] Not exist home information")
	}
	valAsbytes, _ := stub.GetState(key + "#home")
	return valAsbytes, nil
}

// ============================================================================================================================
// PRead - read pet information
// ============================================================================================================================
func (t *PS) PRead(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("[PET QUERY] Incorrect number of arguments. Expecting 1")
	}
	key := args[0]
	conf, _ := stub.GetState(key)
	if conf == nil {
		return []byte("[PET QUERY] Not exist userID"), errors.New("[PET QUERY] Not exist userID")
	}
	userInfo := UserInfo{}
	json.Unmarshal(conf, &userInfo)
	if userInfo.PN == "0" {
		return []byte("[PET QUERY] Not exist pet information"), errors.New("[PET QUERY] Not exist pet information")
	}
	valAsbytes, _ := stub.GetState(key + "#pet")
	return valAsbytes, nil
}

// ============================================================================================================================
// CSearch - search city information on citycode
// ============================================================================================================================
func (t *PS) CSearch(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("[CITY SEARCH] Incorrect number of arguments. Expecting 1")
	}
	key := args[0]
	cityCode := CityCode{}
	valAsbytes, _ := stub.GetState(_CCstr)
	json.Unmarshal(valAsbytes, &cityCode)
	if key == "R103" {
		return []byte(cityCode.R103), nil
	} else if key == "R104" {
		return []byte(cityCode.R104), nil
	} else if key == "R105" {
		return []byte(cityCode.R105), nil
	}
	return []byte("[CITY SEARCH] Wrong city ccode"), nil
}

// ============================================================================================================================
// HDelete - delete home information
// ============================================================================================================================
func (t *PS) HDelete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("[HOME DELETE] Incorrect number of arguments. Expecting 1")
	}
	userID := args[0]
	conf, _ := stub.GetState(userID)
	if conf == nil {
		return nil, errors.New("[HOME DELETE] Not exist userID")
	}
	userInfo := UserInfo{}
	json.Unmarshal(conf, &userInfo)
	if userInfo.CC == "0" {
		return nil, errors.New("[HOME DELETE] Not exist home information")
	}
	userInfo.CC = "0"
	userInfo.AP = "0"
	homeAsset := HomeAsset{}
	jsonAsBytesU, _ := json.Marshal(userInfo)
	jsonAsBytesH, _ := json.Marshal(homeAsset)
	stub.PutState(userID, jsonAsBytesU)
	stub.PutState(userID+"#home", jsonAsBytesH)
	return nil, nil
}

// ============================================================================================================================
// PDelete - delete pet information
// ============================================================================================================================
func (t *PS) PDelete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("[PET DELETE] Incorrect number of arguments. Expecting 1")
	}
	userID := args[0]
	conf, _ := stub.GetState(userID)
	if conf == nil {
		return nil, errors.New("[PET DELETE] Not exist userID")
	}
	userInfo := UserInfo{}
	json.Unmarshal(conf, &userInfo)
	if userInfo.PN == "0" {
		return nil, errors.New("[PET DELETE] Not exist pet information")
	}
	userInfo.PN = "0"
	petAsset := PetAsset{}
	jsonAsBytesU, _ := json.Marshal(userInfo)
	jsonAsBytesP, _ := json.Marshal(petAsset)
	stub.PutState(userID, jsonAsBytesU)
	stub.PutState(userID+"#pet", jsonAsBytesP)
	return nil, nil
}
