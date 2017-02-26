package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var _CCstr string

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

type CityCode struct {
	R103 string // Region code 103 (nowongu)
	R104 string // Region code 104 (gangnamgu)
	R105 string // Region code 105 (zongrogu)
} // if string is not good, using map... ex Region[int]string / Region[103] = "key1"...

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
	} else if function == "trade_insert" {
		return t.TInsert(stub, args)
	} else if function == "user_change" {
		return t.UChange(stub, args)
	} else if function == "pet_change" {
		return t.PChange(stub, args)
	}

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
	} else if function == "trade_search" {
		return t.TSearch(stub, args)
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

	userID := args[0]
	cc := args[1]
	address := args[2]
	hometype := args[3]
	room := args[4]
	area := args[5]
	elevator := args[6]
	parking := args[7]

	userInfo.CC = cc
	homeAsset.Address = address
	homeAsset.HomeType = hometype
	homeAsset.Room = room
	homeAsset.Area = area
	homeAsset.Elevator = elevator
	homeAsset.Parking = parking

	jsonAsBytesU, _ := json.Marshal(userInfo)
	jsonAsBytesH, _ := json.Marshal(homeAsset)
	stub.PutState(userID, jsonAsBytesU)
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
	cityCode := CityCode{}
	var cstr string
	var start, end int
	confCC, _ := stub.GetState(_CCstr)
	json.Unmarshal(confCC, &cityCode)
	if userInfo.CC == "R103" {
		cstr = cityCode.R103
	} else if userInfo.CC == "R104" {
		cstr = cityCode.R104
	} else if userInfo.CC == "R105" {
		cstr = cityCode.R105
	}
	for i, v := range cstr {
		if v == 47 {
			end = i
			if cstr[start+1:end+1] == userID+"/" {
				cstr = cstr[:start+1] + cstr[end+1:]
				break
			}
			start = end
		}
	}
	if userInfo.CC == "R103" {
		cityCode.R103 = cstr
	} else if userInfo.CC == "R104" {
		cityCode.R104 = cstr
	} else if userInfo.CC == "R105" {
		cityCode.R105 = cstr
	}
	userInfo.CC = "0"
	userInfo.AP = "0"
	homeAsset := HomeAsset{}
	jsonAsBytesU, _ := json.Marshal(userInfo)
	jsonAsBytesC, _ := json.Marshal(cityCode)
	jsonAsBytesH, _ := json.Marshal(homeAsset)
	stub.PutState(userID, jsonAsBytesU)
	stub.PutState(_CCstr, jsonAsBytesC)
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

// ============================================================================================================================
// TInsert - insert transaction information
// ============================================================================================================================
func (t *PS) TInsert(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 7 {
		return nil, errors.New("[TRADE INSSERT] Incorrect number of arguments. Expecting 7")
	}
	psid := args[0]
	csid := args[1]
	ts := args[2]
	te := args[3]
	tc := args[4]
	ta := args[5]
	th := args[6]

	tradeRec := TradeRec{}
	tradeRec.PSID = psid
	tradeRec.CSID = csid
	tradeRec.TS = ts
	tradeRec.TE = te
	tradeRec.TC = tc
	tradeRec.TA = ta
	tradeRec.TH = th
	jsonAsBytes, _ := json.Marshal(tradeRec)
	stub.PutState(psid+"#"+csid+"#"+tc, jsonAsBytes)

	return nil, nil
}

// ============================================================================================================================
// TSearch - search trade information
// ============================================================================================================================
func (t *PS) TSearch(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, errors.New("[TRADE SEARCH] Incorrect number of arguments. Expecting 3")
	}
	psid := args[0]
	csid := args[1]
	tc := args[2]
	valAsbytes, _ := stub.GetState(psid + "#" + csid + "#" + tc)
	if valAsbytes == nil {
		return []byte("[TRADE SEARCH] Not exist transaction"), errors.New("[TRADE SEARCH] Not exist transaction")
	}
	return valAsbytes, nil
}

// ============================================================================================================================
// UChange - change user information (PW, AP)
// ============================================================================================================================
func (t *PS) UChange(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, errors.New("[USER CHANGE] Incorrect number of arguments. Expecting 3")
	}
	confUser, _ := stub.GetState(args[0])
	if confUser == nil {
		return nil, errors.New("[USER CHANGE] Not exist userID")
	}
	userInfo := UserInfo{}
	json.Unmarshal(confUser, &userInfo)
	if userInfo.CC == "0" && args[2] == "1" {
		return nil, errors.New("[USER CHANGE] Can't change AP to 1...(not exist home)")
	}
	var pw, ap string
	userID := args[0]
	if args[1] != "0" {
		pw = args[1]
		userInfo.PW = pw
	}
	ap = args[2]

	cc := userInfo.CC
	var cstr string
	var start, end int
	cityCode := CityCode{}
	confCC, _ := stub.GetState(_CCstr)
	json.Unmarshal(confCC, &cityCode)
	if ap == "1" && cc != "0" && userInfo.AP == "0" {
		if cc == "R103" {
			cityCode.R103 = cityCode.R103 + userID + "/"
		} else if cc == "R104" {
			cityCode.R104 = cityCode.R104 + userID + "/"
		} else if cc == "R105" {
			cityCode.R105 = cityCode.R105 + userID + "/"
		}
		jsonAsBytesC, _ := json.Marshal(cityCode)
		stub.PutState(_CCstr, jsonAsBytesC)
	} else if ap == "0" && cc != "0" && userInfo.AP == "1" {
		for i, v := range cstr {
			if v == 47 {
				end = i
				if cstr[start+1:end+1] == userID+"/" {
					cstr = cstr[:start+1] + cstr[end+1:]
					break
				}
				start = end
			}
		}
		if userInfo.CC == "R103" {
			cityCode.R103 = cstr
		} else if userInfo.CC == "R104" {
			cityCode.R104 = cstr
		} else if userInfo.CC == "R105" {
			cityCode.R105 = cstr
		}
		jsonAsBytesC, _ := json.Marshal(cityCode)
		stub.PutState(_CCstr, jsonAsBytesC)
	}
	userInfo.AP = ap
	jsonAsBytesU, _ := json.Marshal(userInfo)
	stub.PutState(userID, jsonAsBytesU)
	return nil, nil
}

// ============================================================================================================================
// PChange - change pet information (SIZE, NS, Vac)
// ============================================================================================================================
func (t *PS) PChange(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
		return nil, errors.New("[PET CHANGE] Incorrect number of arguments. Expecting 4")
	}
	confUser, _ := stub.GetState(args[0])
	if confUser == nil {
		return nil, errors.New("[PET CHANGE] Not exist userID")
	}
	userInfo := UserInfo{}
	json.Unmarshal(confUser, &userInfo)
	if userInfo.PN != "1" {
		return nil, errors.New("[PET CHANGE] Not exist pet")
	}
	valAsbytes, _ := stub.GetState(args[0] + "#pet")
	petAsset := PetAsset{}
	json.Unmarshal(valAsbytes, &petAsset)
	var size, ns, vac string
	userID := args[0]
	if args[1] != "0" {
		size = args[1]
		petAsset.Size = size
	}
	if args[2] != "0" {
		ns = args[2]
		petAsset.NS = ns
	}
	if args[3] != "0" {
		vac = args[3]
		petAsset.Vac = vac
	}
	jsonAsBytesP, _ := json.Marshal(petAsset)
	stub.PutState(userID+"#pet", jsonAsBytesP)
	return nil, nil
}
