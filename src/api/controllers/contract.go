package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	api "unicontract/src/api"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/engine/execengine/function"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"
)

// Operations about Contract
type ContractController struct {
	beego.Controller
}

func (c *ContractController) parseProtoRequestBody() (contract *protos.Contract, err error, status int32) {
	contentType := c.Ctx.Input.Header("Content-Type")
	requestBody := c.Ctx.Input.RequestBody
	contract = &protos.Contract{}
	// return err init
	if contentType == "application/x-protobuf" {
		err = proto.Unmarshal(requestBody, contract)
		if err != nil {
			uniledgerlog.Error("contract parseRequestBody unmarshal err ", err)
			err = fmt.Errorf("contract parseRequestBody unmarshal err ")
			status = api.RESPONSE_STATUS_BadRequest
			return
		}
		//todo temp
		uniledgerlog.Warn(contract)

		fmt.Sprintf("[api] match |%s [token =%s, Content-Type =%s]",
			contentType)
		uniledgerlog.Info(fmt.Sprintf("[API] match|%-32s \t[token = %s, Content-Type = %s]", c.Ctx.Request.RequestURI,
			c.Ctx.Request.Method, contentType))
	}
	return
}

// responseJsonBody
func (c *ContractController) responseJsonBody(data string, msg string) {
	responseData := new(protos.Response)

	responseData.Code = api.RESPONSE_STATUS_OK
	responseData.Msg = msg
	data = base64.StdEncoding.EncodeToString([]byte(data))
	responseData.Result = data
	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseJsonBodyCode ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write(body)
}

func (c *ContractController) responseJsonBodyCode(status int32, data string, msg string) {
	responseData := new(protos.Response)
	responseData.Code = status
	responseData.Msg = msg
	//todo test
	data = base64.StdEncoding.EncodeToString([]byte(data))
	responseData.Result = data

	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseJsonBodyCode ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(body))
}

func (c *ContractController) responseWithCode(status int32, data string) {
	responseData := new(protos.Response)
	responseData.Code = status
	responseData.Msg = ""
	data = base64.StdEncoding.EncodeToString([]byte(data))
	responseData.Result = data
	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseJsonBodyCode ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(body))
}

// special for contractArray to proto[]
func fromContractModelArrayStrToContracts(contractModelStr string) (protos.ContractList, error) {
	// 1. to contractModel
	var contractModel []model.ContractModel
	err := json.Unmarshal([]byte(contractModelStr), &contractModel)
	// 2. to contract
	var contractList protos.ContractList
	var contracts []*protos.Contract
	if err != nil {
		uniledgerlog.Error("error fromContractModelArrayStrToContracts", err)
		return contractList, err
	}
	contracts = make([]*protos.Contract, len(contractModel))
	for i := 0; i < len(contractModel); i++ {
		//contracts[i] = &contractModel[i].Contract
		contracts[i], err = model.FromContractModelToContractProto(contractModel[i])
	}
	contractList.Contracts = contracts
	uniledgerlog.Info("query contract len is ", len(contractModel))
	return contractList, nil
}

// special for contractOutputs Array to proto[] only for queryLog
func fromContractOutputsModelArrayStrToContractsForLog(contractOutputsModelStr string) (protos.ContractExecuteLogList, error) {
	// 1. to contractOutputModel
	var contractOutput []model.ContractOutput
	err := json.Unmarshal([]byte(contractOutputsModelStr), &contractOutput)
	// 2. to contract
	var contractExecuteLogList protos.ContractExecuteLogList
	var contractExecuteLogs []*protos.ContractExecuteLog
	if err != nil {
		uniledgerlog.Error("error fromContractOutputsModelArrayStrToContractsForLog", err)
		return contractExecuteLogList, err
	}
	contractExecuteLogs = make([]*protos.ContractExecuteLog, len(contractOutput))
	for i := 0; i < len(contractOutput); i++ {
		tempTransaction := contractOutput[i].Transaction
		tempRelation := tempTransaction.Relation
		tempContractBody := tempTransaction.ContractModel.ContractBody
		taskId := tempRelation.TaskId
		if taskId == "" {
			uniledgerlog.Error("taskId is blank!", err)
			return contractExecuteLogList, err
		}
		tempContractComponents := tempContractBody.ContractComponents
		var tempContractComponent model.ContractComponent
		for j := 0; j < len(tempContractComponents); j++ {
			if tempContractComponents[j].TaskId == taskId {
				tempContractComponent = *tempContractComponents[j]
				break
			}
		}

		uniledgerlog.Error(tempContractComponent)
		contractExecuteLogs[i] = &protos.ContractExecuteLog{
			ContractHashId: tempRelation.ContractHashId,
			TaskId:         taskId,
			Timestamp:      tempTransaction.Timestamp,
			Caption:        tempContractComponent.Caption,
			Cname:          tempContractComponent.Cname,
			CreateTime:     tempContractBody.CreateTime,
			Ctype:          tempContractComponent.Ctype,
			Description:    tempContractComponent.Description,
			State:          tempContractComponent.State,
			MetaAttribute:  tempContractComponent.MetaAttribute,
		}

	}
	contractExecuteLogList.ContractLogs = contractExecuteLogs
	uniledgerlog.Info("query contractExecuteLogs len is ", len(contractExecuteLogs))
	return contractExecuteLogList, nil
}

// Create [POST]
func (c *ContractController) Create() {
	cost_start := time.Now()

	contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(status, "", err.Error())
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, status)
		return
	}
	contractModel, err := model.FromContractProtoToContractModel(*contract)
	uniledgerlog.Warn("contractModel:\n", contractModel)
	contractModel.ContractHead = &model.ContractHead{
		Version: 1,
	}
	//TODO 额外验证 合约基本字段、owners、component为空
	contractHead := contractModel.ContractHead
	contractBody := contractModel.ContractBody
	uniledgerlog.Warn("contractBody:\n", contractBody)
	if contractHead == nil || contractBody == nil {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contract 验证不通过, Head or Body is blank!")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}

	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contract 验证不通过!")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}
	contract_write_time := monitor.Monitor.NewTiming()
	ok := core.WriteContract(*contractModel)
	if !ok {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "API[Create] insert contract fail!")
		uniledgerlog.Debug(c.Ctx.Request.RequestURI, "API[Create] insert contract fail!")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}
	contract_write_time.Send("contract_write")
	c.responseJsonBody(contract.Id, "API[Create] insert contract Id "+contractModel.Id+"]")
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)

}

// QueryContractContent QueryAll
func (c *ContractController) QueryContractContent() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)
	/*------------------- requestParams end ------------------*/

	uniledgerlog.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractId=%s]",
		c.Ctx.Request.RequestURI, token, owner, contractId))
	if token == "" {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_Forbidden, "", "服务器拒绝请求")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_Forbidden)
		return
	}
	if contractId == "" {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contractId is blank!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}
	contractModelStr, err := rethinkdb.GetContractContentByMapCondition(requestParamMap)
	//uniledgerlog.Warn("QueryContractContent:\n", contractModelStr)
	if err != nil {
		uniledgerlog.Error("API[QueryContractContent]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_QUERY_ERROR, "", "API[QueryContractContent]合约查询错误!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)
		return
	}

	if contractModelStr == "" {
		uniledgerlog.Warn("API[QueryContractContent]合约(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", "API[QueryContractContent]合约(Id="+contractId+")不存在: ")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}

	c.responseJsonBody(contractModelStr, "API[QueryContractContent]查询合约成功!")
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
}

// QueryPublishContract GET
func (c *ContractController) QueryPublishContract() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractState := "Contract_Create"
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)
	/*------------------- requestParams end ------------------*/
	//uniledgerlog.Warn("Body: ", c.Ctx.Request.Body)

	uniledgerlog.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractState=%s, contractId=%s]",
		c.Ctx.Request.RequestURI, token, owner, contractState, contractId))
	if token == "" {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_Forbidden, "", "服务器拒绝请求")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_Forbidden)
		return
	}
	if contractId == "" {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contractId is blank!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}
	contractModelStr, err := rethinkdb.GetPublishContractByMapCondition(requestParamMap)
	if err != nil {
		uniledgerlog.Error("API[QueryPublishContract]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_QUERY_ERROR, "", "API[QueryPublishContract]合约查询错误!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)
		return
	}

	if contractModelStr == "" {
		uniledgerlog.Warn("API[QueryPublishContract]合约(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", "API[QueryPublishContract]合约(Id="+contractId+")不存在: ")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}

	c.responseJsonBody(contractModelStr, "API[QueryPublishContract]查询合约成功!")
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
}

// Query GET
func (c *ContractController) Query() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	/*------------------- requestParams start ------------------*/
	contractState, _ := requestParamMap["status"].(string)
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)
	/*------------------- requestParams end ------------------*/
	uniledgerlog.Warn("Body: ", c.Ctx.Request.Body)
	//uniledgerlog.Warn("Header: ", c.Ctx.Request.Header)

	uniledgerlog.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractState=%s, contractId=%s]",
		c.Ctx.Request.RequestURI, owner, contractState, contractId))

	if contractId == "" {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", "contractId is blank!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}
	contractModelStr, err := rethinkdb.GetOneContractByMapCondition(requestParamMap)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_QUERY_ERROR, "", "API[Query]合约查询错误!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)
		return
	}

	if contractModelStr == "" {
		uniledgerlog.Warn("API[Query]合约(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", "API[Query]合约(Id="+contractId+")不存在: ")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}

	contractProto, err := model.FromContractModelStrToContractProto(contractModelStr)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + "), 转换失败(fromContractModelStrToContract)")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_CONVERT_ERROR, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)
		return
	}
	contractProtoBytes, err := proto.Marshal(contractProto)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + "), 转换失败(proto.Marshal) ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_CONVERT_ERROR, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)
		return
	}
	contractProtoStr := string(contractProtoBytes)
	c.responseJsonBody(contractProtoStr, "API[Query]查询合约成功!")
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
}

// QueryRecords QueryAll
func (c *ContractController) QueryAll() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractState, _ := requestParamMap["status"].(string)
	owner, _ := requestParamMap["owner"].(string)

	contractId, _ := requestParamMap["contractId"].(string)
	//if !ok {
	//	uniledgerlog.Error("contractId type error")
	//}
	contractName, _ := requestParamMap["contractName"].(string)
	/*------------------- requestParams end ------------------*/
	uniledgerlog.Warn("Body: ", c.Ctx.Request.Body)
	//uniledgerlog.Warn("Header: ", c.Ctx.Request.Header)

	uniledgerlog.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractState=%s, contractId=%s, contractName=%s]",
		c.Ctx.Request.RequestURI, token, owner, contractState, contractId, contractName))
	if token == "" {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_Forbidden, "", "服务器拒绝请求")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_Forbidden)
		return
	}

	contractModelStr, err := rethinkdb.GetContractsByMapCondition(requestParamMap)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id="+contractId+")查询错误: ", err)
		c.responseJsonBodyCode(api.RESPONSE_STATUS_QUERY_ERROR, "", "API[Query]合约查询错误!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)
		return
	}

	if contractModelStr == "" {
		uniledgerlog.Warn("API[Query]合约(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", "API[Query]合约(Id="+contractId+")不存在: ")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}

	contractListProto, err := fromContractModelArrayStrToContracts(contractModelStr)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + "), 转换失败(fromContractModelStrToContract)")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}
	contractListProtoBytes, err := proto.Marshal(&contractListProto)
	if err != nil {
		uniledgerlog.Error("API[QueryALl]合约, 转换失败(proto.Marshal) ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_CONVERT_ERROR, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)
		return
	}
	contractProtoStr := string(contractListProtoBytes)
	c.responseJsonBody(contractProtoStr, "API[Query]查询合约成功!")
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
}

// QueryRecords QueryLog
func (c *ContractController) QueryLog() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractState, _ := requestParamMap["status"].(string)
	owner, _ := requestParamMap["owner"].(string)

	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		uniledgerlog.Error("contractId type error")
	}
	contractName, _ := requestParamMap["contractName"].(string)
	/*------------------- requestParams end ------------------*/
	uniledgerlog.Warn("Body: ", c.Ctx.Request.Body)
	//uniledgerlog.Warn("Header: ", c.Ctx.Request.Header)

	uniledgerlog.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractState=%s, contractId=%s, contractName=%s]",
		c.Ctx.Request.RequestURI, token, owner, contractState, contractId, contractName))
	if token == "" {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_Forbidden, "", "服务器拒绝请求")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_Forbidden)
		return
	}

	contractOutputsModelStr, err := rethinkdb.GetContractsLogByMapCondition(requestParamMap)

	if err != nil {
		uniledgerlog.Error("API[QueryLog]合约log(Id="+contractId+")查询错误: ", err)
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", "API[QueryLog]查询错误!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}

	if contractOutputsModelStr == "" {
		uniledgerlog.Warn("API[QueryLog]合约log(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", "API[QueryLog](Id="+contractId+")不存在: ")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}
	//todo 需要过滤字段,只提取需要的字段!
	contractExecuteLogListProto, err := fromContractOutputsModelArrayStrToContractsForLog(contractOutputsModelStr)
	uniledgerlog.Warn(contractExecuteLogListProto)
	if err != nil {
		uniledgerlog.Error("API[QueryLog]合约(Id=" + contractId + "), 转换失败(fromContractOutputsModelArrayStrToContractsForLog)")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_CONVERT_ERROR, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)
		return
	}
	contractExecuteLogListProtoBytes, err := proto.Marshal(&contractExecuteLogListProto)
	if err != nil {
		uniledgerlog.Error("API[QueryLog]合约, 转换失败(proto.Marshal) ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_CONVERT_ERROR, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)
		return
	}
	contractExecuteLogListProtoStr := string(contractExecuteLogListProtoBytes)
	c.responseJsonBody(contractExecuteLogListProtoStr, "API[QueryLog]查询成功!")
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
	//c.responseJsonBody(contractProtoStr, true, "API[Query] success!")
}

// PressTest POST
func (c *ContractController) PressTest() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	startTime, _ := requestParamMap["start"].(string)
	endTime, _ := requestParamMap["end"].(string)
	contract, err, status := c.parseProtoRequestBody()
	if startTime == "" {
		startTime = common.GenTimestamp()
	}
	if endTime == "" {
		endTime, _ = common.GenSpecialTimestampAfterSeconds(startTime, 300)
	}
	contractCaptionTemp := contract.ContractBody.Caption + "_" + time.Nanosecond.String()
	contractIdTemp := contract.ContractBody.ContractId + "_" + time.Nanosecond.String()
	contract.ContractBody.ContractId = contractIdTemp
	contract.ContractBody.Caption = contractCaptionTemp
	contract.ContractBody.ContractState = "Contract_Signature"

	if err != nil {
		c.responseJsonBodyCode(status, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, status)
		return
	}

	//todo 1. replace createTime, Signatures, owner, start and end time!

	//uniledgerlog.Warn("Input contractDeserialize:\n", common.StructSerialize(contract))
	//contractModel := fromContractToContractModel(contract)
	contractModel, err := model.FromContractProtoToContractModel(*contract)
	/*-------------------------- this for press test generate Id start---------------------*/
	// add random string
	randomString := common.GenerateUUID() + "_node" + c.Ctx.Request.RequestURI + "_token_" + token
	contractModel.ContractBody.Caption = randomString
	contractModel.ContractBody.Description = randomString

	contractModel.ContractBody.CreateTime = common.GenTimestamp()
	contractModel.ContractBody.StartTime = startTime
	contractModel.ContractBody.EndTime = endTime

	// lost head lead to nil pointer
	if contractModel.ContractHead == nil {
		contractModel.ContractHead = &model.ContractHead{
			Version: 1,
		}
	}
	contractOwnersLen := 1
	// 生成的合约签名人个数
	contractSignaturesLen := contractOwnersLen

	//generate contractOwnersLen keypair
	owners := make(map[string]string)
	ownersPubkeys := make([]string, contractOwnersLen)
	for i := 0; i < contractOwnersLen; i++ {
		publicKeyBase58, privateKeyBase58 := common.GenerateKeyPair()
		owners[publicKeyBase58] = privateKeyBase58
		ownersPubkeys[i] = publicKeyBase58
	}

	// random contractOwners 随机生成的合约拥有者数组
	contractOwners := ownersPubkeys
	contractModel.ContractBody.ContractOwners = contractOwners
	// 生成 签名
	contractSignatures := make([]*model.ContractSignature, contractSignaturesLen)
	for i := 0; i < contractSignaturesLen; i++ {
		ownerPubkey := ownersPubkeys[i]
		privateKey := owners[ownerPubkey]
		contractSignatures[i] = &model.ContractSignature{
			OwnerPubkey:   ownerPubkey,
			Signature:     contractModel.Sign(privateKey),
			SignTimestamp: common.GenTimestamp(),
		}
	}

	contractModel.ContractBody.ContractSignatures = contractSignatures
	contractModel.Id = contractModel.GenerateId()

	/*-------------------------- this for press test end----------------------*/

	// no verify id again!
	contractValid := contractModel.IsSignatureValid()
	if !contractValid {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contract 非法")
		uniledgerlog.Debug("API[PressTest] token is", token)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}
	ok := core.WriteContract(*contractModel)
	if !ok {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "API[PressTest] insert contract fail!")
		uniledgerlog.Debug(c.Ctx.Request.RequestURI, "API[PressTest] insert contract fail!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}
	uniledgerlog.Warn("API[PressTest] InsertContract success!")
	c.responseJsonBody(contract.Id, "API[PressTest] insert contract Id "+contractModel.Id+"]")
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)

}

//demo使用---------------------------------------------------------------------------------------------------------------
// QueryOutput GET
func (c *ContractController) QueryOutput() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	if len(token) == 0 {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_Forbidden, "", "服务器拒绝请求")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_Forbidden)
		return
	}

	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contractId type is error!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}
	if len(contractId) == 0 {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contractId is blank!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}

	output, err := rethinkdb.QueryOutput(contractId)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_QUERY_ERROR, "", "API[Query]合约查询错误!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(output))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)

}

// QueryOutputNum GET
func (c *ContractController) QueryOutputNum() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	if len(token) == 0 {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_Forbidden, "", "服务器拒绝请求")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_Forbidden)
		return
	}

	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contractId type is error!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}
	if len(contractId) == 0 {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contractId is blank!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}

	count, err := rethinkdb.QueryOutputNum(contractId)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_QUERY_ERROR, "", "API[QueryOutputNum]合约查询错误!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"count":%d}`, count)))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
}

// QueryOutputDuration GET
func (c *ContractController) QueryOutputDuration() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	if len(token) == 0 {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_Forbidden, "", "服务器拒绝请求")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_Forbidden)
		return
	}

	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contractId type is error!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}
	if len(contractId) == 0 {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_BadRequest, "", "contractId is blank!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)
		return
	}

	startTime, err := rethinkdb.QueryContractStartTime(contractId)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(api.RESPONSE_STATUS_QUERY_ERROR, "", "API[QueryOutputDuration]合约查询错误!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)
		return
	}

	nowTime := common.GenTimestamp()

	start, err := strconv.Atoi(startTime)
	now, err := strconv.Atoi(nowTime)

	hours := ((now - start) / 1000) / 3600

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"duration":%d}`, hours)))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
}

// QueryAccountBalance GET
func (c *ContractController) QueryAccountBalance() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	result, err := function.FuncQueryAccountBalance()
	if err != nil {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}

	data, _ := result.GetData().(string)

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(data))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
}

// QueryAmmeterBalance GET
func (c *ContractController) QueryAmmeterBalance() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	result, err := function.FuncQueryAmmeterBalance()
	if err != nil {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_OK, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
		return
	}

	data, _ := result.GetData().(string)

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(data))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
}

// QueryRecords GET
func (c *ContractController) QueryRecords() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	if len(token) == 0 {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_Forbidden, "", "服务器拒绝请求")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_Forbidden)
		return
	}

	str, err := rethinkdb.GetTransactionRecords()
	if err != nil {
		c.responseJsonBodyCode(api.RESPONSE_STATUS_QUERY_ERROR, "", err.Error())
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(str))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)
}

//demo使用---------------------------------------------------------------------------------------------------------------
