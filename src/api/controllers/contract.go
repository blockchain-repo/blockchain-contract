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
		//uniledgerlog.Warn(contract)
		uniledgerlog.Debug(fmt.Sprintf("[API] match|%-32s \t[token = %s, Content-Type = %s]", c.Ctx.Request.RequestURI,
			c.Ctx.Request.Method, contentType))
	}
	return
}

// todo un test
func (c *ContractController) responseProto(status int32, msg string, data string) {
	responseData := new(protos.Response)
	responseData.Code = status
	responseData.Msg = msg
	data = base64.StdEncoding.EncodeToString([]byte(data))
	responseData.Result = data
	output, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseProto ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write(output)
}

func (c *ContractController) responseJson(status int32, msg string, data string) {
	responseData := new(protos.Response)
	responseData.Code = status
	responseData.Msg = msg
	responseData.Result = data
	output, err := json.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseJson ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write(output)
}

func (c *ContractController) responseContract(status int32, msg string, data *protos.Contract) {
	responseData := new(protos.ResponseContract)
	responseData.Code = status
	responseData.Msg = msg
	responseData.Result = data
	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseContract ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(body))
}

func (c *ContractController) responseContracts(status int32, msg string, data []*protos.Contract) {
	responseData := new(protos.ResponseContracts)
	responseData.Code = status
	responseData.Msg = msg
	responseData.Result = data
	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseContracts ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(body))
}

/********************* todo temp for pagination start *********************/
func (c *ContractController) responseContractPagination(status int32, msg string, data *protos.ContractPagination) {
	responseData := new(protos.ResponseContractPagination)
	responseData.Code = status
	responseData.Msg = msg
	responseData.Result = data
	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseContractPagination ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(body))
}

func fromContractModelArrayStrToPaginationContracts(contractModelStr string, page int32, pageSize int32, total int32) (protos.ContractPagination, error) {
	// 1. to contractModel
	var contractModel []model.ContractModel
	err := json.Unmarshal([]byte(contractModelStr), &contractModel)
	// 2. to contract
	var pagination protos.ContractPagination
	var contracts []*protos.Contract
	if err != nil {
		uniledgerlog.Error("error fromContractModelArrayStrToPaginationContracts", err)
		return pagination, err
	}
	contracts = make([]*protos.Contract, total)
	for i := 0; i < int(total); i++ {
		contracts[i], err = model.FromContractModelToContractProto(contractModel[i])
	}

	pagination.Data = contracts
	pagination.Page = page
	pagination.PageSize = pageSize
	pagination.Total = total

	uniledgerlog.Debug("query contract len is ", len(contractModel))
	return pagination, nil
}

/********************* todo temp for pagination end *********************/

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
		contracts[i], err = model.FromContractModelToContractProto(contractModel[i])
	}
	contractList.Contracts = contracts
	uniledgerlog.Debug("query contract len is ", len(contractModel))
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
	uniledgerlog.Debug("query contractExecuteLogs len is ", len(contractExecuteLogs))
	return contractExecuteLogList, nil
}

// Create [POST]
func (c *ContractController) Create() {
	cost_start := time.Now()

	contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseProto(status, err.Error(), "")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, status)()
		return
	}
	contractModel, err := model.FromContractProtoToContractModel(*contract)
	uniledgerlog.Debug("contractModel:\n", contractModel)
	contractModel.ContractHead = &model.ContractHead{
		Version: 1,
	}
	//TODO 额外验证 合约基本字段、owners、component为空
	contractHead := contractModel.ContractHead
	contractBody := contractModel.ContractBody
	uniledgerlog.Debug("contractBody:\n", contractBody)
	if contractHead == nil || contractBody == nil {
		c.responseProto(api.RESPONSE_STATUS_BadRequest, "contract 验证不通过, Head or Body is blank!", "")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}

	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseProto(api.RESPONSE_STATUS_BadRequest, "contract 验证不通过!", "")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}
	contract_write_time := monitor.Monitor.NewTiming()
	ok := core.WriteContract(*contractModel)
	if !ok {
		c.responseProto(api.RESPONSE_STATUS_BadRequest, "API[Create] insert contract fail!", "")
		uniledgerlog.Debug(c.Ctx.Request.RequestURI, "API[Create] insert contract fail!")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}
	contract_write_time.Send("contract_write")
	c.responseProto(api.RESPONSE_STATUS_OK, "API[Create] insert contract Id "+contractModel.Id+"]", contract.Id)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()

}

// QueryContractContent QueryAll
func (c *ContractController) QueryContractContent() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)

	//contractId := c.GetString("contractId")
	//owner := c.GetString("owner")
	/*------------------- requestParams end ------------------*/

	uniledgerlog.Debug(fmt.Sprintf("[API] match |%s [owner =%s, contractId=%s]",
		c.Ctx.Request.RequestURI, owner, contractId))

	if contractId == "" {
		c.responseProto(api.RESPONSE_STATUS_BadRequest, "contractId is blank!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}
	contractModelStr, err := rethinkdb.GetContractContentByCondition(contractId, owner)
	//uniledgerlog.Warn("QueryContractContent:\n", contractModelStr)
	if err != nil {
		uniledgerlog.Error("API[QueryContractContent]合约(Id=" + contractId + ")查询错误: ")
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[QueryContractContent]合约查询错误!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	if contractModelStr == "" {
		uniledgerlog.Error("API[QueryContractContent]合约(Id=" + contractId + ")不存在: ")
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[QueryContractContent]合约(Id="+contractId+")不存在: ", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	c.responseProto(api.RESPONSE_STATUS_OK, "API[QueryContractContent]查询合约成功!", contractModelStr)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

// QueryPublishContract GET
func (c *ContractController) QueryPublishContract() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)
	contractState := "Contract_Create"
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)

	//contractId := c.GetString("contractId")
	//owner := c.GetString("owner")
	//// TODO
	//contractState := c.GetString("contractState", "Contract_Create")
	/*------------------- requestParams end ------------------*/

	uniledgerlog.Debug(fmt.Sprintf("[API] match |%s [owner =%s, contractState=%s, contractId=%s]",
		c.Ctx.Request.RequestURI, owner, contractState, contractId))
	if contractId == "" {
		uniledgerlog.Error("API[QueryPublishContract] contractId is blank!")
		c.responseProto(api.RESPONSE_STATUS_BadRequest, "contractId is blank!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}
	contractModelStr, err := rethinkdb.GetPublishContractByCondition(contractId, owner, contractState)
	if err != nil {
		uniledgerlog.Error("API[QueryPublishContract]合约(Id=" + contractId + ")查询错误: ")
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[QueryPublishContract]合约查询错误!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	if contractModelStr == "" {
		uniledgerlog.Error("API[QueryPublishContract]合约(Id=" + contractId + ")不存在: ")
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[QueryPublishContract]合约(Id="+contractId+")不存在: ", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	c.responseProto(api.RESPONSE_STATUS_OK, "API[QueryPublishContract]查询合约成功!", contractModelStr)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

// Query GET
func (c *ContractController) Query() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)
	contractState, _ := requestParamMap["status"].(string)
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)

	//contractId := c.GetString("contractId")
	//owner := c.GetString("owner")
	//// TODO
	//contractState := c.GetString("status", "")
	/*------------------- requestParams end ------------------*/

	uniledgerlog.Debug(fmt.Sprintf("[API] match |%s [owner =%s, contractState=%s, contractId=%s]",
		c.Ctx.Request.RequestURI, owner, contractState, contractId))

	if contractId == "" {
		c.responseJson(api.RESPONSE_STATUS_BadRequest, "contractId is blank!", "")
		//c.responseProto(api.RESPONSE_STATUS_BadRequest, "contractId is blank!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}
	contractModelStr, err := rethinkdb.GetOneContractByCondition(contractId, owner, contractState)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[Query]合约查询错误!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	if contractModelStr == "" {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")不存在: ")
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[Query]合约(Id="+contractId+")不存在: ", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	contractProto, err := model.FromContractModelStrToContractProto(contractModelStr)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + "), 转换失败(fromContractModelStrToContract)")
		c.responseProto(api.RESPONSE_STATUS_CONVERT_ERROR, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)()
		return
	}
	contractProtoBytes, err := proto.Marshal(contractProto)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + "), 转换失败(proto.Marshal) ")
		c.responseProto(api.RESPONSE_STATUS_CONVERT_ERROR, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)()
		return
	}
	contractProtoStr := string(contractProtoBytes)
	c.responseProto(api.RESPONSE_STATUS_OK, "API[Query]查询合约成功!", contractProtoStr)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

// QueryRecords QueryAll
func (c *ContractController) QueryAll() {
	cost_start := time.Now()
	/*------------------- requestParams start ------------------*/
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)
	contractState, _ := requestParamMap["status"].(string)
	owner, _ := requestParamMap["owner"].(string)

	contractId, _ := requestParamMap["contractId"].(string)
	//if !ok {
	//	uniledgerlog.Error("contractId type error")
	//}
	contractName, _ := requestParamMap["contractName"].(string)

	//contractId := c.GetString("contractId")
	//owner := c.GetString("owner")
	//// TODO contractName
	//contractState := c.GetString("status", "")
	//contractName := c.GetString("contractName", "")
	/*------------------- requestParams end ------------------*/

	uniledgerlog.Debug(fmt.Sprintf("[API] match |%s [owner =%s, contractState=%s, contractId=%s, contractName=%s]",
		c.Ctx.Request.RequestURI, owner, contractState, contractId, contractName))

	contractModelStr, err := rethinkdb.GetContractsByCondition(contractId, owner, contractState)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id="+contractId+")查询错误: ", err)
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[Query]合约查询错误!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	if contractModelStr == "" {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")不存在: ")
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[Query]合约(Id="+contractId+")不存在: ", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	contractListProto, err := fromContractModelArrayStrToContracts(contractModelStr)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + "), 转换失败(fromContractModelStrToContract)")
		c.responseProto(api.RESPONSE_STATUS_CONVERT_ERROR, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)()
		return
	}
	contractListProtoBytes, err := proto.Marshal(&contractListProto)
	if err != nil {
		uniledgerlog.Error("API[QueryALl]合约, 转换失败(proto.Marshal) ")
		c.responseProto(api.RESPONSE_STATUS_CONVERT_ERROR, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)()
		return
	}
	contractProtoStr := string(contractListProtoBytes)
	c.responseProto(api.RESPONSE_STATUS_OK, "API[Query]查询合约成功!", contractProtoStr)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

// QueryRecords QueryLog
func (c *ContractController) QueryLog() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)
	contractState, _ := requestParamMap["status"].(string)
	owner, _ := requestParamMap["owner"].(string)
	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		uniledgerlog.Error("contractId type error")
	}
	contractName, _ := requestParamMap["contractName"].(string)

	//contractId := c.GetString("contractId")
	//owner := c.GetString("owner")
	//// TODO
	//contractState := c.GetString("status", "")
	//contractName := c.GetString("contractName", "")
	/*------------------- requestParams end ------------------*/

	uniledgerlog.Debug(fmt.Sprintf("[API] match |%s [owner =%s, contractState=%s, contractId=%s, contractName=%s]",
		c.Ctx.Request.RequestURI, owner, contractState, contractId, contractName))

	contractOutputsModelStr, err := rethinkdb.GetContractsLogByCondition(contractId, owner, contractState)

	if err != nil {
		uniledgerlog.Error("API[QueryLog]合约log(Id="+contractId+")查询错误: ", err)
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[QueryLog]查询错误!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	if contractOutputsModelStr == "" {
		uniledgerlog.Error("API[QueryLog]合约log(Id=" + contractId + ")不存在: ")
		c.responseProto(api.RESPONSE_STATUS_QUERY_ERROR, "API[QueryLog](Id="+contractId+")不存在: ", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}
	//todo 需要过滤字段,只提取需要的字段!
	contractExecuteLogListProto, err := fromContractOutputsModelArrayStrToContractsForLog(contractOutputsModelStr)
	uniledgerlog.Warn(contractExecuteLogListProto)
	if err != nil {
		uniledgerlog.Error("API[QueryLog]合约(Id=" + contractId + "), 转换失败(fromContractOutputsModelArrayStrToContractsForLog)")
		c.responseProto(api.RESPONSE_STATUS_CONVERT_ERROR, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)()
		return
	}
	contractExecuteLogListProtoBytes, err := proto.Marshal(&contractExecuteLogListProto)
	if err != nil {
		uniledgerlog.Error("API[QueryLog]合约, 转换失败(proto.Marshal) ")
		c.responseProto(api.RESPONSE_STATUS_CONVERT_ERROR, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONVERT_ERROR)()
		return
	}
	contractExecuteLogListProtoStr := string(contractExecuteLogListProtoBytes)
	c.responseProto(api.RESPONSE_STATUS_OK, "API[QueryLog]查询成功!", contractExecuteLogListProtoStr)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

// PressTest POST
func (c *ContractController) PressTest() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	startTime, _ := requestParamMap["start"].(string)
	endTime, _ := requestParamMap["end"].(string)

	//token := c.GetString("token")
	//startTime := c.GetString("start")
	//endTime := c.GetString("end")

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
		c.responseJson(status, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, status)()
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
		c.responseJson(api.RESPONSE_STATUS_BadRequest, "contract 非法", "")
		uniledgerlog.Error("API[PressTest] token is", token)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}
	ok := core.WriteContract(*contractModel)
	if !ok {
		c.responseJson(api.RESPONSE_STATUS_BadRequest, "API[PressTest] insert contract fail!", "")
		uniledgerlog.Error(c.Ctx.Request.RequestURI, "API[PressTest] insert contract fail!")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}
	uniledgerlog.Debug("API[PressTest] InsertContract success!")
	c.responseJson(api.RESPONSE_STATUS_OK, "API[PressTest] insert contract Id "+contractModel.Id+"]", contract.Id)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()

}

//demo使用---------------------------------------------------------------------------------------------------------------
// QueryOutput GET
func (c *ContractController) QueryOutput() {
	cost_start := time.Now()
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		c.responseJson(api.RESPONSE_STATUS_BadRequest, "", "contractId type is error!")
		return
	}
	//contractId := c.GetString("contractId")
	if len(contractId) == 0 {
		c.responseJson(api.RESPONSE_STATUS_BadRequest, "contractId is blank!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}

	output, err := rethinkdb.QueryOutput(contractId)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJson(api.RESPONSE_STATUS_QUERY_ERROR, "API[Query]合约查询错误!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(output))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()

}

// QueryOutputNum GET
func (c *ContractController) QueryOutputNum() {
	cost_start := time.Now()

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	//contractId := c.GetString("contractId")
	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		c.responseJson(api.RESPONSE_STATUS_BadRequest, "", "contractId type is error!")
		return
	}
	if len(contractId) == 0 {
		c.responseJson(api.RESPONSE_STATUS_BadRequest, "contractId is blank!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}

	count, err := rethinkdb.QueryOutputNum(contractId)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJson(api.RESPONSE_STATUS_QUERY_ERROR, "API[QueryOutputNum]合约查询错误!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"count":%d}`, count)))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

// QueryOutputDuration GET
func (c *ContractController) QueryOutputDuration() {
	cost_start := time.Now()
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)
	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		c.responseJson(api.RESPONSE_STATUS_BadRequest, "", "contractId type is error!")
		return
	}

	//contractId := c.GetString("contractId")

	if len(contractId) == 0 {
		c.responseJson(api.RESPONSE_STATUS_BadRequest, "contractId is blank!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}

	startTime, err := rethinkdb.QueryContractStartTime(contractId)
	if err != nil {
		uniledgerlog.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJson(api.RESPONSE_STATUS_QUERY_ERROR, "API[QueryOutputDuration]合约查询错误!", "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	nowTime := common.GenTimestamp()

	start, err := strconv.Atoi(startTime)
	now, err := strconv.Atoi(nowTime)

	hours := ((now - start) / 1000) / 3600

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"duration":%d}`, hours)))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

// QueryAccountBalance GET
func (c *ContractController) QueryAccountBalance() {
	cost_start := time.Now()

	result, err := function.FuncQueryAccountBalance()
	if err != nil {
		c.responseJson(api.RESPONSE_STATUS_OK, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
		return
	}

	data, _ := result.GetData().(string)

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(data))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

// QueryAmmeterBalance GET
func (c *ContractController) QueryAmmeterBalance() {
	cost_start := time.Now()

	result, err := function.FuncQueryAmmeterBalance()
	if err != nil {
		c.responseJson(api.RESPONSE_STATUS_OK, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
		return
	}

	data, _ := result.GetData().(string)

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(data))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

// QueryRecords GET
func (c *ContractController) QueryRecords() {
	cost_start := time.Now()

	str, err := rethinkdb.GetTransactionRecords()
	if err != nil {
		c.responseJson(api.RESPONSE_STATUS_QUERY_ERROR, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_QUERY_ERROR)()
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(api.HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(str))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK)()
}

//demo使用---------------------------------------------------------------------------------------------------------------
