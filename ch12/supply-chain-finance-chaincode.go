package main
import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)

type SCFChaincode struct {

}

//========================================================================
// 定义审核信息
// Name: 审核企业名
// Operator: 操作员
// Approve: 通过与否
// Reason: 理由
// Timestamp: 审核时间
//========================================================================
type Audit struct {
	Name    string `json:"name"`
	Operator string `json:"operator"`
	Approve bool   `json:"approve"`
	Reason  string `json:"reason"`
	Timestamp    string `json:"time"`
}

//========================================================================
// 定义融资申请信息
// TxId: 交易ID
// TxData: 交易指纹
// CoreEnterprise: 核心企业
// Logistics: 物流服务商
// Bank: 金融机构
// Timestamp: 申请时间
// Audits: 审核信息
//========================================================================
type SCF struct {
	TxId string `json:"txId"`
	TxData string `json:"txData"`
	CoreEnterprise string `json:"coreEnterprise"`
	Logistics string `json:"logistics"`
	Bank string `json:"bank"`
	Timestamp string `json:"timestamp"`
	Audits []Audit `json:"audits"`
} 

var logger = shim.NewLogger("main")

//========================================================================
// func-map
//========================================================================
var scfFunctions = map[string]func(shim.ChaincodeStubInterface, []string) pb.Response {
	"financingApply"           : financingApply, // 融资申请
	"applyReview"              : applyReview, // 申请审核
}

//========================================================================
// chaincode 初始化
//========================================================================
func(t *SCFChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response{
	logger.Debug("chaincode初始化成功")
	return shim.Success([]byte("chaincode初始化成功!!!!!!!!"))
}

//========================================================================
// chaincode Invoke
//========================================================================
func(t *SCFChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	bcFunc := scfFunctions[function]
	if bcFunc == nil {
		logger.Error("非法调用函数")
		return shim.Error("无效的调用方法")
	}
	return bcFunc(stub, args)
}


/**
 * 发起融资申请
 * args[0]: 交易id
 * args[1]: 交易指纹
 * args[2]: 核心企业（营业全称）
 * args[3]: 物流企业（营业全称）
 * args[4]: 金融机构（营业全称）
 */
func financingApply(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// --- 参数校验 ---
	if len(args) != 5  {
		logger.Error("参数个数错误")
		return shim.Error("参数个数错误")
	}

	txId := args[0]
	if 0 == len(txId)  {
		logger.Error("交易id不能为空")
		return shim.Error("交易id不能为空")
	}
	txData := args[1]
	if 0 == len(txData)   {
		logger.Error("交易原始数据指纹不能为空")
		return shim.Error("交易原始数据指纹不能为空")
	}
	coreEnterprise := args[2]
	if  0 == len(coreEnterprise) {
		logger.Error("核心企业名不能为空")
		return shim.Error("核心企业名不能为空")
	}
	logistics := args[3]
	if 0 == len(logistics) {
		logger.Error("物流服务商名称不能为空")
		return shim.Error("物流服务商名称不能为空")
	}
	bank := args[4]
	if 0 == len(bank) {
		logger.Error("金融机构名称不能为空")
		return shim.Error("金融机构名称不能为空")
	}

	tm := time.Now()
	dateStr := tm.Format("2006010203:04:05PM")

	// 构建融资申请结构
	apply := SCF {
		TxId: txId,
		TxData: txData,
		CoreEnterprise: coreEnterprise,
		Logistics: logistics,
		Bank: bank,
		Timestamp: dateStr,
	}
	applyBytes, err := json.Marshal(apply)
	if err != nil {
		logger.Error("融资申请序列化失败：", err)
		return shim.Error("融资申请序列化失败")
	}

	err = stub.PutState(txId, applyBytes)
	if err != nil {
		logger.Error("融资申请记录添加失败：", err)
		return shim.Error("融资申请记录添加失败")
	}
	logger.Debug(args, "添加融资申请成功")
	return shim.Success([]byte(txId))
}

/**
 * 融资申请审核
 * args[0]：txId
 * args[1]: 1 通过 -1 不通过
 * args[2]: 审核理由
 * args[3]: 操作员
 */
func applyReview(stub shim.ChaincodeStubInterface, args []string) pb.Response  {
	//---------- 参数个数校验 ---------------
	if 3 != len(args)  {
		logger.Error("参数个数有误")
		return shim.Error("参数个数有误")
	}

	//---------- 参数校验 ---------------
	txId := args[0]
	if 0 == len(txId) {
		logger.Error("交易id不能为空")
		return shim.Error("交易id不能为空")
	}

	approve := args[1]
	if "1" != approve && "-1" != approve {
		logger.Error("审核参数有误")
		return shim.Error("审核参数只能为1或者-1")
	}
	approved, _ := strconv.ParseBool(approve)

	reason := args[2]

	operator := args[3]
	if 0 == len(operator) {
		logger.Error("操作员不能为空")
		return shim.Error("操作员不能为空")
	}


	//---------- 获取请求发送者公钥，并从中获取到相关信息 ---------------
	var createrBytes, err = stub.GetCreator()
	if err != nil{
		logger.Error("获取请求发送者信息失败")
		return shim.Error("获取请求发送者信息失败")
	}

	certStart := bytes.IndexAny(createrBytes, "-----")
	if certStart == -1 {
		logger.Error("获取请求发送者信息失败")
		return shim.Error("获取请求发送者信息失败")
	}
	certText := createrBytes[certStart:]
	bl, _ := pem.Decode(certText)
	if bl == nil {
		logger.Error("获取请求发送者信息失败")
		return shim.Error("获取请求发送者信息失败")
	}
	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		logger.Error("获取请求发送者信息失败")
		return shim.Error("获取请求发送者信息失败")
	}
	uname := cert.Subject.CommonName

	tm := time.Now()
	dateStr := tm.Format("2006010203:04:05PM")

	audit := Audit{
		Name:    uname,
		Approve: approved,
		Reason: reason,
		Timestamp: dateStr,
	}

	//--------- 查询该交易id是否申请个融资 ----------
	scfBytesEX, err := stub.GetState(txId)
	if scfBytesEX == nil || err != nil {
		logger.Error("该融资申请不存在")
		return shim.Error("该融资申请不存在")
	}
	var scfEx SCF
 	err = json.Unmarshal(scfBytesEX, &scfEx)
	if err != nil {
		logger.Error("该融资申请解析失败")
		return shim.Error("该融资申请解析失败")
	}

 	audits := scfEx.Audits

 	// 如果是核心企业签名
	if uname == scfEx.CoreEnterprise {
		// 确保核心企业确认之前无人确认过
		if 0 != len(audits) {
			logger.Error("该笔融资申请已经审核过，请勿重复提交审核")
			return shim.Error("该笔融资申请已经审核过，请勿重复提交审核")
		}
		scfEx.Audits =  append(scfEx.Audits, audit)
	} else if uname == scfEx.Logistics {
		// 确保核心企业已经确认过
		if 1 != len(audits) {
			logger.Error("该笔融资申请已经审核过，请勿重复提交审核")
			return shim.Error("该笔融资申请已经审核过，请勿重复提交审核")
		}
		audit_core_enterprise := audits[0]
		if !audit_core_enterprise.Approve {
			logger.Error("核心企业已经拒绝该融资申请")
			return shim.Error("核心企业已经拒绝该融资申请")
		}
		scfEx.Audits =  append(scfEx.Audits, audit)
	} else if uname == scfEx.Bank {
		// 确保核心企业和物流服务商都确认完成
		if 2 != len(audits) {
			logger.Error("该笔融资申请已经审核过，请勿重复提交审核")
			return shim.Error("该笔融资申请已经审核过，请勿重复提交审核")
		}

		audit_core_enterprise := audits[0]
		if !audit_core_enterprise.Approve {
			logger.Error("核心企业已经拒绝该融资申请")
			return shim.Error("核心企业已经拒绝该融资申请")
		}

		audit_logistics := audits[1]
		if !audit_logistics.Approve {
			logger.Error("物流服务商已经拒绝该融资申请")
			return shim.Error("物流服务商已经拒绝该融资申请")
		}
		scfEx.Audits =  append(scfEx.Audits, audit)
	} else {
		logger.Error("无权限操作该申请")
		return shim.Error("无权限操作该申请")
	}

	applyBytes, err := json.Marshal(scfEx)
	if err != nil {
		logger.Error("融资申请序列化失败：", err)
		return shim.Error("融资申请序列化失败")
	}

	err = stub.PutState(txId, applyBytes)
	if err != nil {
		logger.Error("融资申请审核信息添加失败：", err)
		return shim.Error("融资申请审核信息添加失败")
	}
	logger.Debug(args, "融资申请审核信息添加成功")
	return shim.Success([]byte(txId))
}
