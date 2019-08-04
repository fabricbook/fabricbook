package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {

}

// Define the electricity structure, with 4 properties.  Structure tags are used by encoding/json library


// 个人光伏发电
type PersonalPhotovoltaicPowerGeneration struct {
	Name  string `json:"name"`
	MoneyAccount  float64 `json:"moneyAccount"`
	ElectricityDegree float64 `json:"electricityDegree"`
	CreateTime string `json:"createTime"`
}

// 发电厂
type PowerPlant struct {
	Name  string `json:"name"`
	MoneyAccount  float64 `json:"moneyAccount"`
	ElectricityDegree float64 `json:"electricityDegree"`
	CreateTime string `json:"createTime"`
}

// 消费者
type Consumer struct {
	Name  string `json:"name"`
	MoneyAccount float64 `json:"moneyAccount"`
	ElectricityDegree float64 `json:"electricityDegree"`
	CreateTime string `json:"createTime"`
}

// 电网公司
type PowerGridCorp struct {
	Name  string `json:"name"`
	MoneyAccount float64 `json:"moneyAccount"`
	ElectricityDegree float64 `json:"electricityDegree"`
	CreateTime string `json:"createTime"`
}

// 交易记录
type PowerTransactionHistory struct {
	Seller string `json:"seller"`
	Buyer string `json:"buyer"`
	PayAmount float64 `json:"payMoneyAmount"`
	ReceiveAmount float64 `json:"receiveMoneyAmount"`
	PowerGridAmount float64 `json:"powerGridAmount"`
	ElectricityDegree float64 `json:"electricityDegree"`
	CreateTime string `json:"createTime"`
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
	case "addPersonalPhotovoltaicPowerGeneration": //新增个人光伏发电
		return s.addPersonalPhotovoltaicPowerGeneration(APIstub, args)
	case "updatePersonalPhotovoltaicPowerGeneration": //更新个人光伏发电
		return s.updatePersonalPhotovoltaicPowerGeneration(APIstub, args)
	case "queryPersonalPhotovoltaicPowerGeneration": //查询个人光伏发电
		return s.queryPersonalPhotovoltaicPowerGeneration(APIstub, args)	
	case "addPowerPlant": //新增发电厂
		return s.addPowerPlant(APIstub, args)
	case "updatePowerPlant": //更新发电厂
		return s.updatePowerPlant(APIstub, args)
	case "queryPowerPlant": //查询发电厂
		return s.queryPowerPlant(APIstub, args)	
	case "addConsumer": //新增消费者
		return s.addConsumer(APIstub, args)
	case "updateConsumer": //更新消费者
		return s.updateConsumer(APIstub, args)
	case "queryConsumer": //查询消费者
		return s.queryConsumer(APIstub, args)
	case "addPowerGridCorp": //新增电网公司
		return s.addPowerGridCorp(APIstub, args)
	case "updatePowerGridCorp": //更新电网公司
		return s.updatePowerGridCorp(APIstub, args)
	case "queryPowerGridCorp": //查询电网公司
		return s.queryPowerGridCorp(APIstub, args)
	case "addPowerTransactionHistory": //新增交易记录
		return s.addPowerTransactionHistory(APIstub, args)
	case "queryPowerTransactionHistory": //查询交易记录
		return s.queryPowerTransactionHistory(APIstub, args)
	default:
		return shim.Error("InvalID Smart Contract function name.")
	}

}


func (s *SmartContract) addPersonalPhotovoltaicPowerGeneration(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	personalPhotovoltaicPowerGenerationAsBytes, _ := APIstub.GetState(args[0])

	if personalPhotovoltaicPowerGenerationAsBytes != nil {
		return shim.Error("personalPhotovoltaicPowerGeneration " + args[0] + " already exists")
	}

	moneyAccount, moneyAccountErr := strconv.ParseFloat(args[2], 64)
	if(moneyAccountErr != nil){
		return shim.Error("moneyAccount is not correct")
	}

	electricityDegree, electricityDegreeErr := strconv.ParseFloat(args[3], 64)
	if(electricityDegreeErr != nil){
		return shim.Error("electricityDegree is not correct")
	}

	personalPhotovoltaicPowerGeneration := PersonalPhotovoltaicPowerGeneration{Name: args[1], MoneyAccount: moneyAccount, ElectricityDegree: electricityDegree, CreateTime: args[4]}

	personalPhotovoltaicPowerGenerationAsBytes, _ = json.Marshal(personalPhotovoltaicPowerGeneration)

	APIstub.PutState(args[0], personalPhotovoltaicPowerGenerationAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updatePersonalPhotovoltaicPowerGeneration(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	personalPhotovoltaicPowerGenerationAsBytes, _ := APIstub.GetState(args[0])

	if personalPhotovoltaicPowerGenerationAsBytes == nil {
		return shim.Error("personalPhotovoltaicPowerGeneration " + args[0] + " does not exist")
	}

	personalPhotovoltaicPowerGeneration := PersonalPhotovoltaicPowerGeneration{}

	json.Unmarshal(personalPhotovoltaicPowerGenerationAsBytes, &personalPhotovoltaicPowerGeneration)
	if args[1] != ""{
		personalPhotovoltaicPowerGeneration.Name = args[1]
	}

	moneyAccount, moneyAccountErr := strconv.ParseFloat(args[2], 64)
	if(moneyAccountErr != nil){
		return shim.Error("moneyAccount is not correct")
	}
	if args[2] != ""{
		personalPhotovoltaicPowerGeneration.MoneyAccount = moneyAccount
	}

	electricityDegree, electricityDegreeErr := strconv.ParseFloat(args[3], 64)
	if(electricityDegreeErr != nil){
		return shim.Error("electricityDegree is not correct")
	}
	if args[3] != ""{
		personalPhotovoltaicPowerGeneration.ElectricityDegree = electricityDegree
	}
	if args[4] != ""{
		personalPhotovoltaicPowerGeneration.CreateTime = args[4]
	}

	personalPhotovoltaicPowerGenerationAsBytes, _ = json.Marshal(personalPhotovoltaicPowerGeneration)
	APIstub.PutState(args[0], personalPhotovoltaicPowerGenerationAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryPersonalPhotovoltaicPowerGeneration(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	personalPhotovoltaicPowerGenerationAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(personalPhotovoltaicPowerGenerationAsBytes)
}

func (s *SmartContract) addPowerPlant(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	powerPlantAsBytes, _ := APIstub.GetState(args[0])

	if powerPlantAsBytes != nil {
		return shim.Error("powerPlant " + args[0] + " already exists")
	}

	moneyAccount, moneyAccountErr := strconv.ParseFloat(args[2], 64)
	if(moneyAccountErr != nil){
		return shim.Error("moneyAccount is not correct")
	}

	electricityDegree, electricityDegreeErr := strconv.ParseFloat(args[3], 64)
	if(electricityDegreeErr != nil){
		return shim.Error("electricityDegree is not correct")
	}

	powerPlant := PowerPlant{Name: args[1], MoneyAccount: moneyAccount, ElectricityDegree: electricityDegree, CreateTime: args[4]}

	powerPlantAsBytes, _ = json.Marshal(powerPlant)

	APIstub.PutState(args[0], powerPlantAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updatePowerPlant(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	powerPlantAsBytes, _ := APIstub.GetState(args[0])

	if powerPlantAsBytes == nil {
		return shim.Error("powerPlant " + args[0] + " does not exist")
	}

	powerPlant := PowerPlant{}

	json.Unmarshal(powerPlantAsBytes, &powerPlant)
	if args[1] != ""{
		powerPlant.Name = args[1]
	}

	moneyAccount, moneyAccountErr := strconv.ParseFloat(args[2], 64)
	if(moneyAccountErr != nil){
		return shim.Error("moneyAccount is not correct")
	}
	if args[2] != ""{
		powerPlant.MoneyAccount = moneyAccount
	}

	electricityDegree, electricityDegreeErr := strconv.ParseFloat(args[3], 64)
	if(electricityDegreeErr != nil){
		return shim.Error("electricityDegree is not correct")
	}
	if args[3] != ""{
		powerPlant.ElectricityDegree = electricityDegree
	}

	if args[4] != ""{
		powerPlant.CreateTime = args[4]
	}

	powerPlantAsBytes, _ = json.Marshal(powerPlant)
	APIstub.PutState(args[0], powerPlantAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryPowerPlant(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	powerPlantAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(powerPlantAsBytes)
}

func (s *SmartContract) addConsumer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	consumerAsBytes, _ := APIstub.GetState(args[0])

	if consumerAsBytes != nil {
		return shim.Error("consumer " + args[0] + " already exists")
	}

	moneyAccount, moneyAccountErr := strconv.ParseFloat(args[2], 64)
	if(moneyAccountErr != nil){
		return shim.Error("moneyAccount is not correct")
	}
	
	electricityDegree, electricityDegreeErr := strconv.ParseFloat(args[3], 64)
	if(electricityDegreeErr != nil){
		return shim.Error("electricityDegree is not correct")
	}

	consumer := Consumer{Name: args[1], MoneyAccount: moneyAccount, ElectricityDegree: electricityDegree, CreateTime: args[4]}

	consumerAsBytes, _ = json.Marshal(consumer)

	APIstub.PutState(args[0], consumerAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateConsumer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	consumerAsBytes, _ := APIstub.GetState(args[0])

	if consumerAsBytes == nil {
		return shim.Error("consumer " + args[0] + " does not exist")
	}

	consumerAsBytes, _ = APIstub.GetState(args[0])
	consumer := Consumer{}

	json.Unmarshal(consumerAsBytes, &consumer)
	if args[1] != ""{
		consumer.Name = args[1]
	}

	moneyAccount, moneyAccountErr := strconv.ParseFloat(args[2], 64)
	if(moneyAccountErr != nil){
		return shim.Error("moneyAccount is not correct")
	}
	if args[2] != ""{
		consumer.MoneyAccount = moneyAccount
	}

	electricityDegree, electricityDegreeErr := strconv.ParseFloat(args[3], 64)
	if(electricityDegreeErr != nil){
		return shim.Error("electricityDegree is not correct")
	}
	if args[3] != ""{
		consumer.ElectricityDegree = electricityDegree
	}

	if args[4] != ""{
		consumer.CreateTime = args[4]
	}

	consumerAsBytes, _ = json.Marshal(consumer)
	APIstub.PutState(args[0], consumerAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryConsumer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	consumerAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(consumerAsBytes)
}

func (s *SmartContract) addPowerGridCorp(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	powerGridCorpAsBytes, _ := APIstub.GetState(args[0])

	if powerGridCorpAsBytes != nil {
		return shim.Error("powerGridCorp " + args[0] + " already exists")
	}

	moneyAccount, moneyAccountErr := strconv.ParseFloat(args[2], 64)
	if(moneyAccountErr != nil){
		return shim.Error("moneyAccount is not correct")
	}
	
	electricityDegree, electricityDegreeErr := strconv.ParseFloat(args[3], 64)
	if(electricityDegreeErr != nil){
		return shim.Error("electricityDegree is not correct")
	}

	powerGridCorp := PowerGridCorp{Name: args[1], MoneyAccount: moneyAccount, ElectricityDegree: electricityDegree, CreateTime: args[4]}
	powerGridCorpAsBytes, _ = json.Marshal(powerGridCorp)

	APIstub.PutState(args[0], powerGridCorpAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updatePowerGridCorp(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	powerGridCorpAsBytes, _ := APIstub.GetState(args[0])

	if powerGridCorpAsBytes == nil {
		return shim.Error("powerGridCorp " + args[0] + " does not exist")
	}

	powerGridCorpAsBytes, _ = APIstub.GetState(args[0])
	powerGridCorp := PowerGridCorp{}

	json.Unmarshal(powerGridCorpAsBytes, &powerGridCorp)

	if args[1] != ""{
		powerGridCorp.Name = args[1]
	}

	moneyAccount, moneyAccountErr := strconv.ParseFloat(args[2], 64)
	if(moneyAccountErr != nil){
		return shim.Error("moneyAccount is not correct")
	}
	if args[2] != ""{
		powerGridCorp.MoneyAccount = moneyAccount
	}

	electricityDegree, electricityDegreeErr := strconv.ParseFloat(args[3], 64)
	if(electricityDegreeErr != nil){
		return shim.Error("electricityDegree is not correct")
	}
	if args[3] != ""{
		powerGridCorp.ElectricityDegree = electricityDegree
	}

	if args[4] != ""{
		powerGridCorp.CreateTime = args[4]
	}

	powerGridCorpAsBytes, _ = json.Marshal(powerGridCorp)
	APIstub.PutState(args[0], powerGridCorpAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryPowerGridCorp(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	powerGridCorpAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(powerGridCorpAsBytes)
}


func (s *SmartContract) addPowerTransactionHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	powerTransactionHistoryAsBytes, _ := APIstub.GetState(args[0])

	if powerTransactionHistoryAsBytes != nil {
		return shim.Error("powerTransactionHistory " + args[0] + " already exists")
	}

	payAmount, payAmountErr := strconv.ParseFloat(args[2], 64)
	if(payAmountErr != nil){
		return shim.Error("payAmount is not correct")
	}

	receiveAmount, receiveAmountErr := strconv.ParseFloat(args[3], 64)
	if(receiveAmountErr != nil){
		return shim.Error("receiveAmount is not correct")
	}

	powerGridAmount, powerGridAmountErr := strconv.ParseFloat(args[4], 64)
	if(powerGridAmountErr != nil){
		return shim.Error("powerGridAmount is not correct")
	}

	electricityDegree, electricityDegreeErr := strconv.ParseFloat(args[5], 64)
	if(electricityDegreeErr != nil){
		return shim.Error("electricityDegree is not correct")
	}

	powerTransactionHistory := PowerTransactionHistory{Seller: args[1], Buyer: args[2], PayAmount: payAmount, ReceiveAmount: receiveAmount, PowerGridAmount: powerGridAmount, ElectricityDegree: electricityDegree, CreateTime: args[6]}

	powerTransactionHistoryAsBytes, _ = json.Marshal(powerTransactionHistory)

	APIstub.PutState(args[0], powerTransactionHistoryAsBytes)

	return shim.Success(nil)
}


func (s *SmartContract) queryPowerTransactionHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	powerTransactionHistoryBytes, _ := APIstub.GetState(args[0])
	return shim.Success(powerTransactionHistoryBytes)
}


// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
