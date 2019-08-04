package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {

}

// Define the drug structure, with 4 properties.  Structure tags are used by encoding/json library


// 种植基地
type PlantBase struct {
	Name  string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address  string `json:"address"`
	TaxNumber string `json:"taxNumber"`
}

// 药企
type DrugCompany struct {
	Name  string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address  string `json:"address"`
	TaxNumber string `json:"taxNumber"`
}

// 药店
type DrugStore struct {
	Name  string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address  string `json:"address"`
	TaxNumber string `json:"taxNumber"`
}

// 药材采摘信息
type MedicinalMaterialPickInfo struct {
	Name string `json:"variety"`
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
	Weather string `json:"weather"`
	PickTime string `json:"pickTime"`
	PlantBaseID string `json:"plantBaseID"`
	PickMark string `json:"pickMark"`
}

// 药材销售信息
type MedicinalMaterialSalesInfo struct {
	MedicinalMaterialPickID string `json:"medicinalMaterialPickID"`
	Name string `json:"name"`
	SalesTime string `json:"salesTime"`
	PlantBaseID string `json:"plantBaseID"`
	DrugCompanyID string `json:"drugCompanyID"`
}

// 药品生产信息
type DrugProductionInfo struct {
	MedicinalMaterialSalesID string `json:"medicinalMaterialSalesID"`
	Name string `json:"name"`
	DrugProductionTime string `json:"drugProductionTime"`
	DrugProductionAddress string `json:"drugProductionAddress"`
	DrugProductionPersonnal string `json:"drugProductionPersonnal"`
	DrugProductionMark string `json:"drugProductionMark"`
	DrugCompanyID string `json:"drugCompanyID"`
}

// 药品流通信息
type DrugCirculationInfo struct {
	DrugCirculationBeginTime string `json:"drugCirculationBeginTime"`
	DrugCirculationEndTime string `json:"drugCirculationEndTime"`
	DrugProductionInfoID string `json:"drugProductionInfoID"`
	DrugStoreID string `json:"drugStoreID"`
	DrugCirculationPersonnal string `json:"drugCirculationPersonnal"`
	DrugCirculationMark string `json:"drugCirculationMark"`
	DrugCompanyID string `json:"drugCompanyID"`
}

// 药品零售信息
type DrugSalesInfo struct {
	DrugCirculationInfoID string `json:"drugCirculationInfoID"`
	DrugSalesTime string `json:"drugSalesTime"`
	DrugSalesPersonnal string `json:"drugSalesPersonnal"`
	DrugStoreID string `json:"drugStoreID"`
	DrugStoreMark string `json:"drugStoreMark"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	switch function {
	case "addPlantBase": //新增种植基地
		return s.addPlantBase(APIstub, args)
	case "updatePlantBase": //更新种植基地
		return s.updatePlantBase(APIstub, args)
	case "queryPlantBase": //查询种植基地
		return s.queryPlantBase(APIstub, args)	
	case "addDrugCompany": //新增药企
		return s.addDrugCompany(APIstub, args)
	case "updateDrugCompany": //更新药企
		return s.updateDrugCompany(APIstub, args)
	case "queryDrugCompany": //查询药企
		return s.queryDrugCompany(APIstub, args)	
	case "addDrugStore": //新增药店
		return s.addDrugStore(APIstub, args)
	case "updateDrugStore": //更新药店
		return s.updateDrugStore(APIstub, args)
	case "queryDrugStore": //查询药店
		return s.queryDrugStore(APIstub, args)	
	case "addMedicinalMaterialPickInfo": //新增药材采摘信息
		return s.addMedicinalMaterialPickInfo(APIstub, args)
	case "updateMedicinalMaterialPickInfo": //更新药材采摘信息
		return s.updateMedicinalMaterialPickInfo(APIstub, args)
	case "queryMedicinalMaterialPickInfo": //查询药材采摘信息
		return s.queryMedicinalMaterialPickInfo(APIstub, args)
	case "addMedicinalMaterialSalesInfo": //新增药材销售信息
		return s.addMedicinalMaterialSalesInfo(APIstub, args)
	case "updateMedicinalMaterialSalesInfo": //更新药材销售信息
		return s.updateMedicinalMaterialSalesInfo(APIstub, args)
	case "queryMedicinalMaterialSalesInfo": //查询药材销售信息
		return s.queryMedicinalMaterialSalesInfo(APIstub, args)
	case "addDrugProductionInfo": //新增药品生产信息
		return s.addDrugProductionInfo(APIstub, args)
	case "updateDrugProductionInfo": //更新药品生产信息
		return s.updateDrugProductionInfo(APIstub, args)
	case "queryDrugProductionInfo": //查询药品生产信息
		return s.queryDrugProductionInfo(APIstub, args)	
	case "addDrugCirculationInfo": //新增药品流通信息
		return s.addDrugCirculationInfo(APIstub, args)
	case "updateDrugCirculationInfo": //更新药品流通信息
		return s.updateDrugCirculationInfo(APIstub, args)
	case "queryDrugCirculationInfo": //查询药品流通信息
		return s.queryDrugCirculationInfo(APIstub, args)
	case "addDrugSalesInfo": //新增药品销售信息
		return s.addDrugSalesInfo(APIstub, args)
	case "updateDrugSalesInfo": //更新药品销售信息
		return s.updateDrugSalesInfo(APIstub, args)
	case "queryDrugSalesInfo": //查询药品销售信息
		return s.queryDrugSalesInfo(APIstub, args)
	case "queryDrugTraceabilityInfo": //查询药品溯源信息
		return s.queryDrugTraceabilityInfo(APIstub, args)
	default:
		return shim.Error("InvalID Smart Contract function name.")
	}

}


func (s *SmartContract) addPlantBase(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	plantBaseAsBytes, _ := APIstub.GetState(args[0])

	if plantBaseAsBytes != nil {
		return shim.Error("plantBase " + args[0] + " already exists")
	}

	plantBase := PlantBase{Name: args[1], PhoneNumber: args[2], Address: args[3], TaxNumber: args[4]}

	plantBaseAsBytes, _ = json.Marshal(plantBase)

	APIstub.PutState(args[0], plantBaseAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updatePlantBase(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	plantBaseAsBytes, _ := APIstub.GetState(args[0])

	if plantBaseAsBytes == nil {
		return shim.Error("plantBase " + args[0] + " does not exist")
	}

	plantBase := PlantBase{}

	json.Unmarshal(plantBaseAsBytes, &plantBase)
	if args[1] != ""{
		plantBase.Name = args[1]
	}
	if args[2] != ""{
		plantBase.PhoneNumber = args[2]
	}
	if args[3] != ""{
		plantBase.Address = args[3]
	}
	if args[4] != ""{
		plantBase.TaxNumber = args[4]
	}

	plantBaseAsBytes, _ = json.Marshal(plantBase)
	APIstub.PutState(args[0], plantBaseAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryPlantBase(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	plantBaseAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(plantBaseAsBytes)
}

func (s *SmartContract) addDrugCompany(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	drugCompanyAsBytes, _ := APIstub.GetState(args[0])

	if drugCompanyAsBytes != nil {
		return shim.Error("drugCompany " + args[0] + " already exists")
	}

	drugCompany := DrugCompany{Name: args[1], PhoneNumber: args[2], Address: args[3], TaxNumber: args[4]}

	drugCompanyAsBytes, _ = json.Marshal(drugCompany)

	APIstub.PutState(args[0], drugCompanyAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateDrugCompany(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	drugCompanyAsBytes, _ := APIstub.GetState(args[0])

	if drugCompanyAsBytes == nil {
		return shim.Error("drugCompany " + args[0] + " does not exist")
	}

	drugCompany := DrugCompany{}

	json.Unmarshal(drugCompanyAsBytes, &drugCompany)
	if args[1] != ""{
		drugCompany.Name = args[1]
	}
	if args[2] != ""{
		drugCompany.PhoneNumber = args[2]
	}
	if args[3] != ""{
		drugCompany.Address = args[3]
	}
	if args[4] != ""{
		drugCompany.TaxNumber = args[4]
	}

	drugCompanyAsBytes, _ = json.Marshal(drugCompany)
	APIstub.PutState(args[0], drugCompanyAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryDrugCompany(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	drugCompanyAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(drugCompanyAsBytes)
}

func (s *SmartContract) addDrugStore(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	drugStoreAsBytes, _ := APIstub.GetState(args[0])

	if drugStoreAsBytes != nil {
		return shim.Error("drugStore " + args[0] + " already exists")
	}

	drugStore := DrugStore{Name: args[1], PhoneNumber: args[2], Address: args[3], TaxNumber: args[4]}

	drugStoreAsBytes, _ = json.Marshal(drugStore)

	APIstub.PutState(args[0], drugStoreAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateDrugStore(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	drugStoreAsBytes, _ := APIstub.GetState(args[0])

	if drugStoreAsBytes == nil {
		return shim.Error("drugStore " + args[0] + " does not exist")
	}

	drugStoreAsBytes, _ = APIstub.GetState(args[0])
	drugStore := DrugStore{}

	json.Unmarshal(drugStoreAsBytes, &drugStore)
	if args[1] != ""{
		drugStore.Name = args[1]
	}
	if args[2] != ""{
		drugStore.PhoneNumber = args[2]
	}
	if args[3] != ""{
		drugStore.Address = args[3]
	}
	if args[4] != ""{
		drugStore.TaxNumber = args[4]
	}

	drugStoreAsBytes, _ = json.Marshal(drugStore)
	APIstub.PutState(args[0], drugStoreAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryDrugStore(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	drugStoreAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(drugStoreAsBytes)
}

func (s *SmartContract) addMedicinalMaterialPickInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	medicinalMaterialPickInfoAsBytes, _ := APIstub.GetState(args[0])

	if medicinalMaterialPickInfoAsBytes != nil {
		return shim.Error("medicinalMaterialPickInfo " + args[0] + " already exists")
	}

	medicinalMaterialPickInfo := MedicinalMaterialPickInfo{Name: args[1], Longitude: args[2], Latitude: args[3], Weather: args[4], PickTime: args[5], PlantBaseID: args[6], PickMark: args[7]}

	medicinalMaterialPickInfoAsBytes, _ = json.Marshal(medicinalMaterialPickInfo)

	APIstub.PutState(args[0], medicinalMaterialPickInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateMedicinalMaterialPickInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	medicinalMaterialPickInfoAsBytes, _ := APIstub.GetState(args[0])

	if medicinalMaterialPickInfoAsBytes == nil {
		return shim.Error("medicinalMaterialPickInfo " + args[0] + " does not exist")
	}

	medicinalMaterialPickInfoAsBytes, _ = APIstub.GetState(args[0])
	medicinalMaterialPickInfo := MedicinalMaterialPickInfo{}

	json.Unmarshal(medicinalMaterialPickInfoAsBytes, &medicinalMaterialPickInfo)
	if args[1] != ""{
		medicinalMaterialPickInfo.Name = args[1]
	}
	if args[2] != ""{
		medicinalMaterialPickInfo.Longitude = args[2]
	}
	if args[3] != ""{
		medicinalMaterialPickInfo.Latitude = args[3]
	}
	if args[4] != ""{
		medicinalMaterialPickInfo.Weather = args[4]
	}
	if args[5] != ""{
		medicinalMaterialPickInfo.PickTime = args[5]
	}
	if args[6] != ""{
		medicinalMaterialPickInfo.PlantBaseID = args[6]
	}
	if args[7] != ""{
		medicinalMaterialPickInfo.PickMark = args[7]
	}

	medicinalMaterialPickInfoAsBytes, _ = json.Marshal(medicinalMaterialPickInfo)
	APIstub.PutState(args[0], medicinalMaterialPickInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryMedicinalMaterialPickInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	medicinalMaterialPickInfoAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(medicinalMaterialPickInfoAsBytes)
}

func (s *SmartContract) addMedicinalMaterialSalesInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	medicinalMaterialSalesInfoAsBytes, _ := APIstub.GetState(args[0])

	if medicinalMaterialSalesInfoAsBytes != nil {
		return shim.Error("medicinalMaterialSaleInfo " + args[0] + " already exists")
	}

	medicinalMaterialSalesInfo := MedicinalMaterialSalesInfo{MedicinalMaterialPickID: args[1], Name: args[2], SalesTime: args[3], PlantBaseID: args[4], DrugCompanyID: args[5]}

	medicinalMaterialSalesInfoAsBytes, _ = json.Marshal(medicinalMaterialSalesInfo)

	APIstub.PutState(args[0], medicinalMaterialSalesInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateMedicinalMaterialSalesInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	medicinalMaterialSalesInfoAsBytes, _ := APIstub.GetState(args[0])

	if medicinalMaterialSalesInfoAsBytes == nil {
		return shim.Error("medicinalMaterialSalesInfo " + args[0] + " does not exist")
	}

	medicinalMaterialSalesInfoAsBytes, _ = APIstub.GetState(args[0])
	medicinalMaterialSalesInfo := MedicinalMaterialSalesInfo{}

	json.Unmarshal(medicinalMaterialSalesInfoAsBytes, &medicinalMaterialSalesInfo)
	if args[1] != ""{
		medicinalMaterialSalesInfo.MedicinalMaterialPickID = args[1]
	}
	if args[2] != ""{
		medicinalMaterialSalesInfo.Name = args[2]
	}
	if args[3] != ""{
		medicinalMaterialSalesInfo.SalesTime = args[3]
	}
	if args[4] != ""{
		medicinalMaterialSalesInfo.PlantBaseID = args[4]
	}
	if args[5] != ""{
		medicinalMaterialSalesInfo.DrugCompanyID = args[5]
	}

	medicinalMaterialSalesInfoAsBytes, _ = json.Marshal(medicinalMaterialSalesInfo)
	APIstub.PutState(args[0], medicinalMaterialSalesInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryMedicinalMaterialSalesInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	medicinalMaterialSalesInfoAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(medicinalMaterialSalesInfoAsBytes)
}

func (s *SmartContract) addDrugProductionInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	drugProductionInfoAsBytes, _ := APIstub.GetState(args[0])

	if drugProductionInfoAsBytes != nil {
		return shim.Error("drugProductionInfo " + args[0] + " already exists")
	}

	drugProductionInfo := DrugProductionInfo{MedicinalMaterialSalesID: args[1], Name: args[2], DrugProductionTime: args[3], DrugProductionAddress: args[4], DrugProductionPersonnal: args[5], DrugProductionMark: args[6], DrugCompanyID: args[7]}

	drugProductionInfoAsBytes, _ = json.Marshal(drugProductionInfo)

	APIstub.PutState(args[0], drugProductionInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateDrugProductionInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	drugProductionInfoAsBytes, _ := APIstub.GetState(args[0])

	if drugProductionInfoAsBytes == nil {
		return shim.Error("drugProductionInfo " + args[0] + " does not exist")
	}

	drugProductionInfoAsBytes, _ = APIstub.GetState(args[0])
	drugProductionInfo := DrugProductionInfo{}

	json.Unmarshal(drugProductionInfoAsBytes, &drugProductionInfo)
	if args[1] != ""{
		drugProductionInfo.MedicinalMaterialSalesID = args[1]
	}
	if args[2] != ""{
		drugProductionInfo.Name = args[2]
	}
	if args[3] != ""{
		drugProductionInfo.DrugProductionTime = args[3]
	}
	if args[4] != ""{
		drugProductionInfo.DrugProductionAddress = args[4]
	}
	if args[5] != ""{
		drugProductionInfo.DrugProductionPersonnal = args[5]
	}
	if args[6] != ""{
		drugProductionInfo.DrugProductionMark = args[6]
	}
	if args[7] != ""{
		drugProductionInfo.DrugCompanyID = args[7]
	}

	drugProductionInfoAsBytes, _ = json.Marshal(drugProductionInfo)
	APIstub.PutState(args[0], drugProductionInfoAsBytes)

	return shim.Success(nil)
}


func (s *SmartContract) queryDrugProductionInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	drugProductionInfoAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(drugProductionInfoAsBytes)
}

func (s *SmartContract) addDrugCirculationInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	drugCirculationInfoAsBytes, _ := APIstub.GetState(args[0])

	if drugCirculationInfoAsBytes != nil {
		return shim.Error("drugCirculationInfo " + args[0] + " already exists")
	}

	drugCirculationInfo := DrugCirculationInfo{DrugCirculationBeginTime: args[1], DrugCirculationEndTime: args[2], DrugProductionInfoID: args[3], DrugStoreID: args[4], DrugCirculationPersonnal: args[5], DrugCirculationMark: args[6], DrugCompanyID: args[7]}

	drugCirculationInfoAsBytes, _ = json.Marshal(drugCirculationInfo)

	APIstub.PutState(args[0], drugCirculationInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateDrugCirculationInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	drugCirculationInfoAsBytes, _ := APIstub.GetState(args[0])

	if drugCirculationInfoAsBytes == nil {
		return shim.Error("drugCirculationInfo " + args[0] + " does not exist")
	}

	drugCirculationInfoAsBytes, _ = APIstub.GetState(args[0])
	drugCirculationInfo := DrugCirculationInfo{}

	json.Unmarshal(drugCirculationInfoAsBytes, &drugCirculationInfo)
	if args[1] != ""{
		drugCirculationInfo.DrugCirculationBeginTime = args[1]
	}
	if args[2] != ""{
		drugCirculationInfo.DrugCirculationEndTime = args[2]
	}
	if args[3] != ""{
		drugCirculationInfo.DrugProductionInfoID = args[3]
	}
	if args[4] != ""{
		drugCirculationInfo.DrugStoreID = args[4]
	}
	if args[5] != ""{
		drugCirculationInfo.DrugCirculationPersonnal = args[5]
	}
	if args[6] != ""{
		drugCirculationInfo.DrugCirculationMark = args[6]
	}
	if args[7] != ""{
		drugCirculationInfo.DrugCompanyID = args[7]
	}

	drugCirculationInfoAsBytes, _ = json.Marshal(drugCirculationInfo)
	APIstub.PutState(args[0], drugCirculationInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryDrugCirculationInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	drugCirculationInfoAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(drugCirculationInfoAsBytes)
}

func (s *SmartContract) addDrugSalesInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	drugSalesInfoAsBytes, _ := APIstub.GetState(args[0])

	if drugSalesInfoAsBytes != nil {
		return shim.Error("drugSalesInfo " + args[0] + " already exists")
	}

	drugSalesInfo := DrugSalesInfo{DrugCirculationInfoID: args[1], DrugSalesTime: args[2], DrugSalesPersonnal: args[3], DrugStoreID: args[4], DrugStoreMark: args[5]}

	drugSalesInfoAsBytes, _ = json.Marshal(drugSalesInfo)

	APIstub.PutState(args[0], drugSalesInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateDrugSalesInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	drugSalesInfoAsBytes, _ := APIstub.GetState(args[0])

	if drugSalesInfoAsBytes == nil {
		return shim.Error("drugSalesInfo " + args[0] + " does not exist")
	}

	drugSalesInfoAsBytes, _ = APIstub.GetState(args[0])
	drugSalesInfo := DrugSalesInfo{}

	json.Unmarshal(drugSalesInfoAsBytes, &drugSalesInfo)
	if args[1] != ""{
		drugSalesInfo.DrugCirculationInfoID = args[1]
	}
	if args[2] != ""{
		drugSalesInfo.DrugSalesTime = args[2]
	}
	if args[3] != ""{
		drugSalesInfo.DrugSalesPersonnal = args[3]
	}
	if args[4] != ""{
		drugSalesInfo.DrugStoreID = args[4]
	}
	if args[5] != ""{
		drugSalesInfo.DrugStoreMark = args[5]
	}

	drugSalesInfoAsBytes, _ = json.Marshal(drugSalesInfo)
	APIstub.PutState(args[0], drugSalesInfoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryDrugSalesInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	drugSalesInfoAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(drugSalesInfoAsBytes)
}

func (s *SmartContract) queryDrugTraceabilityInfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	//查询药品销售信息
	drugSalesInfoAsBytes, _ := APIstub.GetState(args[0])
	if drugSalesInfoAsBytes == nil {
		return shim.Error("drugSalesInfo " + args[0] + " does not exist")
	}

	drugSalesInfo := DrugSalesInfo{}
	json.Unmarshal(drugSalesInfoAsBytes, &drugSalesInfo)

	//查询药店信息
	drugStoreAsBytes,_ :=  APIstub.GetState(drugSalesInfo.DrugStoreID)
	if drugStoreAsBytes == nil {
		return shim.Error("drugStore " + drugSalesInfo.DrugStoreID + " does not exist")
	}
	drugStore := DrugStore{}
	json.Unmarshal(drugStoreAsBytes, &drugStore)

	//查询药品流通信息
	drugCirculationInfoAsBytes,_ := APIstub.GetState(drugSalesInfo.DrugCirculationInfoID)
	if drugCirculationInfoAsBytes == nil {
		return shim.Error("drugCirculationInfo " + drugSalesInfo.DrugCirculationInfoID + " does not exist")
	}
	drugCirculationInfo := DrugCirculationInfo{}
	json.Unmarshal(drugCirculationInfoAsBytes, &drugCirculationInfo)

	//查询药企信息
	drugCompanyAsBytes, _ := APIstub.GetState(drugCirculationInfo.DrugCompanyID)
	if drugCompanyAsBytes == nil {
		return shim.Error("drugCompany " + drugCirculationInfo.DrugCompanyID + " does not exist")
	}
	drugCompany := DrugCompany{}
	json.Unmarshal(drugCompanyAsBytes, &drugCompany)

	//查询药品生产信息
	drugProductionInfoAsBytes,_ :=  APIstub.GetState(drugCirculationInfo.DrugProductionInfoID)
	if drugProductionInfoAsBytes == nil {
		return shim.Error("drugProductionInfo " + drugCirculationInfo.DrugProductionInfoID + " does not exist")
	}
	drugProductionInfo := DrugProductionInfo{}
	json.Unmarshal(drugProductionInfoAsBytes, &drugProductionInfo)
	
	//查询药材销售信息
	medicinalMaterialSalesInfoAsBytes, _ := APIstub.GetState(drugProductionInfo.MedicinalMaterialSalesID)
	if medicinalMaterialSalesInfoAsBytes == nil {
		return shim.Error("medicinalMaterialSalesInfo " + drugProductionInfo.MedicinalMaterialSalesID + " does not exist")
	}
	medicinalMaterialSalesInfo := MedicinalMaterialSalesInfo{}
	json.Unmarshal(medicinalMaterialSalesInfoAsBytes, &medicinalMaterialSalesInfo)

	//查询药材采摘信息
	medicinalMaterialPickInfoAsBytes, _ := APIstub.GetState(medicinalMaterialSalesInfo.MedicinalMaterialPickID)

	if medicinalMaterialPickInfoAsBytes == nil {
		return shim.Error("medicinalMaterialPickInfo " + medicinalMaterialSalesInfo.MedicinalMaterialPickID + " does not exist")
	}

	medicinalMaterialPickInfo := MedicinalMaterialPickInfo{}
	json.Unmarshal(medicinalMaterialPickInfoAsBytes, &medicinalMaterialPickInfo)

	//查询种植基地信息
	plantBaseAsBytes, _ := APIstub.GetState(medicinalMaterialPickInfo.PlantBaseID)

	if plantBaseAsBytes == nil {
		return shim.Error("plantBase " + medicinalMaterialPickInfo.PlantBaseID + " does not exist")
	}

	plantBase := PlantBase{}
	json.Unmarshal(plantBaseAsBytes, &plantBase)

	drugTraceabilityInfo := map[string]interface{}{
		"drugSalesInfo": &drugSalesInfo, 
		"drugStore": &drugStore,
		"drugCirculationInfo": &drugCirculationInfo,
		"drugCompany": &drugCompany,
		"drugProductionInfo": &drugProductionInfo,
		"medicinalMaterialSalesInfo": &medicinalMaterialSalesInfo,
		"medicinalMaterialPickInfo": &medicinalMaterialPickInfo,
		"plantBase": &plantBase,
	}
	drugTraceabilityInfoAsBytes,_ := json.Marshal(drugTraceabilityInfo)
	return shim.Success(drugTraceabilityInfoAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
