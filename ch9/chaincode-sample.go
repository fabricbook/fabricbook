package main

//引入必要的包
import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//声明一个结构体
type SimpleChaincode struct {
}

//为结构体添加Init方法
//Init方法用于Chaincode实例化时初始化数据
//注意，Chaincode升级时也会调用该方法，避免因为升级导致数据被重置
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// 通过GetStringArgs获取请求参数
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("参数个数错误")
	}

	key := args[0]
	value := args[1]

	// 通过PutState方法将key-value对存入账本
	err := stub.PutState(key, []byte(value))
	if err != nil {
		return shim.Error("数据初始化失败，Chaincode实例化失败")
	}

	return shim.Success(nil)
}

//为结构体添加Invoke方法
//在该方法中实现链码运行中被调用或查询时的处理逻辑
//在Chaincode上的每笔交易都会请求该方法
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// 从交易请求终获取方法名和参数
	fn, args := stub.GetFunctionAndParameters()

	// 定义处理结果字符串和错误
	var result string
	var err error

	switch fn {
	case "set":
		result, err = set(stub, args)
	case "get":
		result, err = get(stub, args)
	default:
		return shim.Error("智能合约不支持该方法")
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(result))
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	// 参数校验
	if len(args) != 2 {
		return "", fmt.Errorf("参数个数错误")
	}

	var value int
	var key string
	var err error

	// 参数转换 
	key = args[0]
	value, err = strconv.Atoi(args[1])
	if err != nil {
		// 第二个参数无法转化为整数
		return "", fmt.Errorf("参数格式错误")
	}

	// 通过PutState将数据写入到账本中
	err = stub.PutState(key, []byte(strconv.Itoa(value)))
	if err != nil {
		return "", fmt.Errorf("%s数据写入错误", args[0])
	}

	return args[1], nil
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	
	if len(args) != 1 {
		return "", fmt.Errorf("参数个数错误")
	}

	key := args[0]
	
	// 通过key从账本中读取数据
	value, err := stub.GetState(key)
	if err != nil {
			return "", fmt.Errorf("获取: %s 失败:%s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("信息不存在")
	}
	return string(value), nil
}

//主函数，需要调用shim.Start（ ）方法
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
