package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	"strconv"
	"time"
	"unicontract/src/api"
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
	uniledgerlog.Debug("parseProtoRequestBody:\n", contract)
	// return err init
	if contentType == "application/x-protobuf" {
		err = proto.Unmarshal(requestBody, contract)
		if err != nil {
			uniledgerlog.Error("contract parseRequestBody unmarshal err ", err)
			err = fmt.Errorf("contract parseRequestBody unmarshal err ")
			status = api.RESPONSE_STATUS_INTERNAL_ERROR
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
	c.Ctx.ResponseWriter.Write(output)
}

func (c *ContractController) responseContract(status int32, msg string, data *protos.Contract) {
	responseData := new(protos.ResponseContract)
	if data == nil {
		responseData.Result = nil
	} else {
		responseData.Result = data
	}
	responseData.Code = status
	responseData.Msg = msg

	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseContract ", err.Error())
	}
	c.Ctx.ResponseWriter.Write([]byte(body))
}

func (c *ContractController) responseContractExecuteLog(status int32, msg string, data *protos.ContractExecuteLog) {
	responseData := new(protos.ResponseContractExecuteLog)
	if data == nil {
		responseData.Result = nil
	} else {
		responseData.Result = data
	}
	responseData.Code = status
	responseData.Msg = msg
	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responseContractExecuteLog ", err.Error())
	}
	c.Ctx.ResponseWriter.Write([]byte(body))
}

/********************* todo temp for pagination start *********************/
func (c *ContractController) responsePaginationContract(status int32, msg string, data *protos.PaginationContract) {
	responseData := new(protos.ResponsePaginationContract)
	if data == nil {
		responseData.Result = nil
	} else {
		responseData.Result = data
	}
	responseData.Code = status
	responseData.Msg = msg
	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responsePaginationContract ", err.Error())
	}
	c.Ctx.ResponseWriter.Write([]byte(body))
}

func (c *ContractController) responsePaginationContractExecuteLog(status int32, msg string, data *protos.PaginationContractExecuteLog) {
	responseData := new(protos.ResponsePaginationContractExecuteLog)
	if data == nil {
		responseData.Result = nil
	} else {
		responseData.Result = data
	}
	responseData.Code = status
	responseData.Msg = msg
	body, err := proto.Marshal(responseData)
	if err != nil {
		uniledgerlog.Error("responsePaginationContractExecuteLog ", err.Error())
	}
	c.Ctx.ResponseWriter.Write([]byte(body))
}

func fromContractModelArrayStrToPaginationContracts(contractModelStr string, page int32, pageSize int32, total int32) (protos.PaginationContract, error) {
	// 1. to contractModel
	var contractModel []model.ContractModel
	err := json.Unmarshal([]byte(contractModelStr), &contractModel)
	// 2. to contract
	var pagination protos.PaginationContract
	var contracts []*protos.Contract
	if err != nil {
		uniledgerlog.Error("error fromContractModelArrayStrToPaginationContracts", err)
		return pagination, err
	}
	lenResult := len(contractModel)
	contracts = make([]*protos.Contract, lenResult)
	for i := 0; i < int(lenResult); i++ {
		contracts[i], err = model.FromContractModelToContractProto(contractModel[i])
	}
	pagination.Data = contracts
	pagination.Page = page
	pagination.PageSize = pageSize
	pagination.Total = total
	uniledgerlog.Debug("query PaginationContract len is %d page=%d , pageSize=%d, total=%d  ", len(contractModel), page, pageSize, total)
	return pagination, nil
}

/********************* todo temp for pagination end *********************/

// special for contractOutputs Array to proto[] only for queryLog
func fromContractOutputsModelArrayStrToPaginationContractsExecuteLog(contractOutputsModelStr string, page int32, pageSize int32, total int32) (protos.PaginationContractExecuteLog, error) {
	// 1. to contractOutputModel
	var contractOutput []model.ContractOutput
	err := json.Unmarshal([]byte(contractOutputsModelStr), &contractOutput)
	// 2. to contract
	var pagination protos.PaginationContractExecuteLog
	var contractExecuteLogs []*protos.ContractExecuteLog
	if err != nil {
		uniledgerlog.Error("error fromContractOutputsModelArrayStrToContractsForLog", err)
		return pagination, err
	}
	lenResult := len(contractOutput)
	contractExecuteLogs = make([]*protos.ContractExecuteLog, lenResult)
	for i := 0; i < lenResult; i++ {
		tempTransaction := contractOutput[i].Transaction
		tempRelation := tempTransaction.Relation
		tempContractBody := tempTransaction.ContractModel.ContractBody
		taskId := tempRelation.TaskId
		if taskId == "" {
			uniledgerlog.Error("taskId is blank!", err)
			return pagination, err
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
	pagination.Data = contractExecuteLogs
	pagination.Page = page
	pagination.PageSize = pageSize
	pagination.Total = total

	uniledgerlog.Debug("query PaginationContractExecuteLog len is ", len(contractExecuteLogs))
	data_test, _ := json.Marshal(pagination)
	uniledgerlog.Warn(string(data_test))
	return pagination, nil
}

// Create [POST]
func (c *ContractController) Create() {
	cost_start := time.Now()
	resultMsg := fmt.Sprintf("%s 查询创建成功!", "API[Create]")
	uniledgerlog.Debug("Create contractModel:\n", cost_start)
	contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseProto(status, err.Error(), "")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, status, err.Error())()
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
		resultMsg = fmt.Sprintf("%s %s ", "API[Create]", "contract 验证不通过, Head or Body is blank!")
		c.responseProto(api.RESPONSE_STATUS_CONTRACT_ERROR_MODEL, resultMsg, "")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONTRACT_ERROR_MODEL, resultMsg)()
		return
	}

	contractValid := contractModel.Validate()
	if !contractValid {
		resultMsg = fmt.Sprintf("%s %s ", "API[Create]", "contract 验证不通过!")
		c.responseProto(api.RESPONSE_STATUS_CONTRACT_ERROR_MODEL, resultMsg, "")
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_CONTRACT_ERROR_MODEL, resultMsg)()
		return
	}
	contract_write_time := monitor.Monitor.NewTiming()
	ok := core.WriteContract(*contractModel)
	if !ok {
		resultMsg = fmt.Sprintf("%s 合约写入失败(WriteContract) ", "API[Create]")
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		uniledgerlog.Error(resultMsg)
		monitor.Monitor.Count("request_fail", 1)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}

	contract_write_time.Send("contract_write")
	c.responseProto(api.RESPONSE_STATUS_OK, resultMsg, contract.Id)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()

}

// QueryContractContent
func (c *ContractController) QueryContractContent() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	contractProductId := c.GetString(api.REQUEST_FIELD_CONTRACT_PRODUCT_ID)
	owner := c.GetString(api.REQUEST_FIELD_CONTRACT_OWNER)
	resultMsg := fmt.Sprintf("%s 查询合约成功!", "API[QueryContractContent]")
	// verify the must length
	if len(contractProductId) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryContractContent]", "contractProductId")
		c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}
	if len(owner) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryContractContent]", "owner")
		c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}

	/*------------------- requestParams end ------------------*/
	contractModelStr, err := rethinkdb.GetContractContentByCondition(contractProductId, owner)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)查询错误! ", "API[QueryContractContent]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg+err.Error())()
		return
	}

	if contractModelStr == "" {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)不存在!", "API[QueryContractContent]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}

	c.responseProto(api.RESPONSE_STATUS_OK, resultMsg, contractModelStr)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

// QueryPublishContract GET
func (c *ContractController) QueryPublishContract() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	contractProductId := c.GetString(api.REQUEST_FIELD_CONTRACT_PRODUCT_ID)
	owner := c.GetString(api.REQUEST_FIELD_CONTRACT_OWNER)
	contractState := c.GetString(api.REQUEST_FIELD_CONTRACT_STATE, "Contract_Create")
	resultMsg := fmt.Sprintf("%s 查询合约成功!", "API[QueryPublishContract]")
	/*------------------- requestParams end ------------------*/
	if len(contractProductId) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryPublishContract]", "contractProductId")
		c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}
	if len(owner) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryPublishContract]", "owner")
		c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}
	if len(contractState) != 0 && !api.REQUEST_CONTRACT_STATE_MAP[contractState] {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryPublishContract]", "contractState")
		c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}

	contractModelStr, err := rethinkdb.GetPublishContractByCondition(contractProductId, owner, contractState)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(Id=%s)查询错误! ", "API[QueryPublishContract]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg+err.Error())()
		return
	}

	if contractModelStr == "" {
		resultMsg = fmt.Sprintf("%s(Id=%s)不存在!", "API[QueryPublishContract]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}

	c.responseProto(api.RESPONSE_STATUS_OK, resultMsg, contractModelStr)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

// Query GET
func (c *ContractController) Query() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	contractProductId := c.GetString(api.REQUEST_FIELD_CONTRACT_PRODUCT_ID)
	owner := c.GetString(api.REQUEST_FIELD_CONTRACT_OWNER)
	contractState := c.GetString(api.REQUEST_FIELD_CONTRACT_STATE)
	resultMsg := fmt.Sprintf("%s 查询合约成功!", "API[Query]")

	/*------------------- requestParams end ------------------*/
	if len(contractProductId) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[Query]", "contractProductId")
		c.responseContract(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}
	if len(owner) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[Query]", "owner")
		c.responseContract(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}
	if len(contractState) != 0 && !api.REQUEST_CONTRACT_STATE_MAP[contractState] {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[Query]", "contractState")
		c.responseContract(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}

	contractModelStr, err := rethinkdb.GetOneContractByCondition(contractProductId, owner, contractState)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)查询错误! msg=%s", "API[Query]", contractProductId, err)
		uniledgerlog.Error(resultMsg)
		c.responseContract(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, nil)
		//c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg+err.Error())()
		return
	}

	if contractModelStr == "" {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)不存在!", "API[Query]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responseContract(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}

	contractProto, err := model.FromContractModelStrToContractProto(contractModelStr)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)转换失败(model.FromContractModelStrToContractProto)! ", "API[Query]", contractProductId)
		c.responseContract(api.RESPONSE_STATUS_INTERNAL_ERROR, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_INTERNAL_ERROR, resultMsg+err.Error())()
		return
	}
	c.responseContract(api.RESPONSE_STATUS_OK, resultMsg, contractProto)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

//QueryAll QueryRecords
func (c *ContractController) QueryAll() {
	cost_start := time.Now()
	resultMsg := fmt.Sprintf("%s 查询成功!", "API[QueryAll]")
	/*------------------- requestParams start ------------------*/
	startTime := c.GetString(api.REQUEST_FIELD_CONTRACT_STARTTIME)
	endTime := c.GetString(api.REQUEST_FIELD_CONTRACT_ENDTIME)
	//todo deal
	uniledgerlog.Warn("startTime", startTime)
	uniledgerlog.Warn("endTime", endTime)

	contractProductId := c.GetString(api.REQUEST_FIELD_CONTRACT_PRODUCT_ID)
	owner := c.GetString(api.REQUEST_FIELD_CONTRACT_OWNER)
	contractState := c.GetString(api.REQUEST_FIELD_CONTRACT_STATE)
	page, err := c.GetInt32(api.REQUEST_FIELD_PAGE, 1)
	if err != nil || page <= 0 {
		resultMsg := fmt.Sprintf("%s page(%v) error!", "API[QueryAll]", page)
		c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
	}
	pageSize, err := c.GetInt32(api.REQUEST_FIELD_PAGE_SIZE, 5)
	if err != nil || pageSize <= 0 {
		resultMsg := fmt.Sprintf("%s pageSize(%v) error!", "API[QueryAll]", pageSize)
		c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
	}
	pageNumStart := (page - 1) * pageSize
	pageNumEnd := pageNumStart + pageSize

	//contractName := c.GetString(api.REQUEST_FIELD_CONTRACT_NAME)
	//_=contractName
	/*------------------- requestParams end ------------------*/

	//if len(owner) == 0 {
	//	resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryAll]", "owner")
	//	c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
	//	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
	//	return
	//}
	if len(contractState) != 0 && !api.REQUEST_CONTRACT_STATE_MAP[contractState] {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryAll]", "contractState")
		c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}

	totalRecords, contractModelStr, err := rethinkdb.GetContractsPaginationByCondition(contractProductId, owner, contractState, pageNumStart, pageNumEnd)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)查询错误! ", "API[QueryAll]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg+err.Error())()
		return
	}

	if contractModelStr == "" {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)不存在!", "API[QueryAll]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}
	//contractListProto, err := fromContractModelArrayStrToContracts(contractModelStr)
	paginationContractProto, err := fromContractModelArrayStrToPaginationContracts(contractModelStr, page, pageSize, totalRecords)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)转换失败(fromContractModelArrayStrToContracts)! ", "API[QueryAll]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_INTERNAL_ERROR, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_INTERNAL_ERROR, resultMsg+err.Error())()
		return
	}
	c.responsePaginationContract(api.RESPONSE_STATUS_OK, resultMsg, &paginationContractProto)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()

}

// QueryRecords QueryLog
func (c *ContractController) QueryLog() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	contractProductId := c.GetString(api.REQUEST_FIELD_CONTRACT_PRODUCT_ID)
	owner := c.GetString(api.REQUEST_FIELD_CONTRACT_OWNER)
	contractState := c.GetString(api.REQUEST_FIELD_CONTRACT_STATE, "Contract_In_Process")
	page, err := c.GetInt32(api.REQUEST_FIELD_PAGE, 1)
	if err != nil || page <= 0 {
		resultMsg := fmt.Sprintf("%s page(%v) error!", "API[QueryLog]", page)
		c.responsePaginationContractExecuteLog(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
	}
	pageSize, err := c.GetInt32(api.REQUEST_FIELD_PAGE_SIZE, 5)
	if err != nil || pageSize <= 0 {
		resultMsg := fmt.Sprintf("%s pageSize(%v) error!", "API[QueryLog]", pageSize)
		c.responsePaginationContractExecuteLog(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
	}
	pageNumStart := (page - 1) * pageSize
	pageNumEnd := pageNumStart + pageSize
	//contractName := c.GetString(api.REQUEST_FIELD_CONTRACT_NAME)
	//_=contractName
	/*------------------- requestParams end ------------------*/
	//uniledgerlog.Debug(fmt.Sprintf("[API] match |%s [owner =%s, contractState=%s, contractId=%s, contractName=%s]",
	//	c.Ctx.Request.RequestURI, owner, contractState, contractId, contractName))
	resultMsg := fmt.Sprintf("%s 查询成功!", "API[QueryLog]")

	if len(contractProductId) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryLog]", "contractProductId")
		c.responsePaginationContractExecuteLog(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}

	totalRecords, contractOutputsModelStr, err := rethinkdb.GetContractsLogPaginationByCondition(contractProductId, owner, contractState, pageNumStart, pageNumEnd)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)查询错误! ", "API[QueryLog]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responsePaginationContractExecuteLog(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg+err.Error())()
		return
	}
	if contractOutputsModelStr == "" {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)不存在!", "API[QueryLog]", contractProductId)
		uniledgerlog.Error(resultMsg)
		c.responsePaginationContractExecuteLog(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}

	//todo 需要过滤字段,只提取需要的字段!
	contractPaginationExecuteLogProto, err := fromContractOutputsModelArrayStrToPaginationContractsExecuteLog(contractOutputsModelStr, page, pageSize, totalRecords)
	//uniledgerlog.Warn(contractPaginationCExecuteLogProto)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(contractProductId=%s)转换失败(fromContractOutputsModelArrayStrToContractsForLog)! ", "API[QueryLog]", contractProductId)
		c.responsePaginationContractExecuteLog(api.RESPONSE_STATUS_INTERNAL_ERROR, resultMsg, nil)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_INTERNAL_ERROR, resultMsg+err.Error())()
		return
	}
	c.responsePaginationContractExecuteLog(api.RESPONSE_STATUS_OK, resultMsg, &contractPaginationExecuteLogProto)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

// PressTest POST
func (c *ContractController) PressTest() {
	cost_start := time.Now()

	/*------------------- requestParams start ------------------*/
	startTime := c.GetString(api.REQUEST_FIELD_CONTRACT_STARTTIME)
	endTime := c.GetString(api.REQUEST_FIELD_CONTRACT_ENDTIME)
	resultMsg := fmt.Sprintf("%s 操作成功!", "API[PressTest]")
	if len(startTime) != 0 {
		_, err := strconv.ParseInt(startTime, 10, 64)
		if err != nil {
			resultMsg = fmt.Sprintf("%s %s 格式错误!", "API[PressTest]", "startTime")
			c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_TYPE, resultMsg, "")
			defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_TYPE, resultMsg)()
			return
		}
	}
	if len(endTime) != 0 {
		_, err := strconv.ParseInt(endTime, 10, 64)
		if err != nil {
			resultMsg = fmt.Sprintf("%s %s 格式错误!", "API[PressTest]", "endTime")
			c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_TYPE, resultMsg, "")
			defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_TYPE, resultMsg)()
			return
		}
	}

	if len(startTime) != 0 && len(endTime) != 0 {
		startTimeInt64, _ := strconv.ParseInt(startTime, 10, 64)
		endTimeInt64, _ := strconv.ParseInt(endTime, 10, 64)
		timeStart := time.Unix(startTimeInt64/1000, 0)
		timeEnd := time.Unix(endTimeInt64/1000, 0)
		if timeEnd.Before(timeStart) {
			resultMsg = fmt.Sprintf("%s %s endTime  需要大于 startTime !", "API[PressTest]", "endTime")
			c.responseProto(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
			defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
			return
		}
	}

	contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		resultMsg = fmt.Sprintf("%s 解析失败(parseProtoRequestBody) ", "API[PressTest]")
		c.responseJson(status, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, status, resultMsg+err.Error())()
		return
	}
	if contract == nil {
		uniledgerlog.Warn("23423")
	}

	randomStr := strconv.FormatInt(time.Now().UnixNano(), 10) + "_" + common.GenerateUUID()
	contractCaptionTemp := contract.ContractBody.Caption + "_" + randomStr
	contractIdTemp := contract.ContractBody.ContractId + "_" + randomStr
	contractProductIdTemp := "CP0001-" + time.Now().Format("20060102150405") + "-" + randomStr
	contract.ContractBody.ContractId = contractIdTemp
	contract.ContractBody.ContractProductId = contractProductIdTemp
	contract.ContractBody.Caption = contractCaptionTemp
	contract.ContractBody.ContractState = "Contract_Signature"

	//todo 1. replace createTime, Signatures, owner, start and end time!

	//uniledgerlog.Warn("Input contractDeserialize:\n", common.StructSerialize(contract))
	//contractModel := fromContractToContractModel(contract)
	contractModel, err := model.FromContractProtoToContractModel(*contract)
	uniledgerlog.Warn(contractModel)
	/*-------------------------- this for press test generate Id start---------------------*/
	// add random string
	randomString := common.GenerateUUID() + "_node" + c.Ctx.Request.RequestURI + "_token_"
	contractModel.ContractBody.Caption = randomString
	contractModel.ContractBody.Description = randomString
	contractModel.ContractBody.CreateTime = common.GenTimestamp()
	/************************** startTime endTime deal start **********************/
	if len(startTime) != 13 {
		_startTime := contractModel.ContractBody.StartTime
		if len(_startTime) == 0 {
			startTime = common.GenTimestamp()
		} else if len(_startTime) != 13 {
			startTime, _ = common.GenSpecialTimestamp(_startTime)
		} else {
			startTime = _startTime
		}
	}
	contractModel.ContractBody.StartTime = startTime

	if len(endTime) != 13 {
		_endTime := contractModel.ContractBody.EndTime
		if len(_endTime) == 0 {
			endTime, _ = common.GenSpecialTimestampAfterSeconds(startTime, 3600*24*360)
		} else if len(_endTime) != 13 {
			endTime, _ = common.GenSpecialTimestamp(_endTime)
		} else {
			endTime = _endTime
		}
	}
	contractModel.ContractBody.EndTime = endTime
	uniledgerlog.Warn("%+v ", contractModel.ContractBody)
	/************************** startTime endTime deal end **********************/

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
	//for i := 0; i < contractOwnersLen; i++ {
	//	publicKeyBase58, privateKeyBase58 := common.GenerateKeyPair()
	//	owners[publicKeyBase58] = privateKeyBase58
	//	ownersPubkeys[i] = publicKeyBase58
	//}
	publicKeyBase58 := "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	privateKeyBase58 := "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	owners[publicKeyBase58] = privateKeyBase58
	ownersPubkeys[0] = publicKeyBase58

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
		resultMsg = fmt.Sprintf("%s 签名验证失败(IsSignatureValid) ", "API[PressTest]")
		c.responseJson(api.RESPONSE_STATUS_INTERNAL_ERROR, resultMsg, "")
		uniledgerlog.Error(resultMsg)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_INTERNAL_ERROR, resultMsg)()
		return
	}
	ok := core.WriteContract(*contractModel)
	if !ok {
		resultMsg = fmt.Sprintf("%s 合约写入失败(WriteContract) ", "API[PressTest]")
		c.responseJson(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		uniledgerlog.Error(resultMsg)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}
	uniledgerlog.Debug(resultMsg)
	c.responseJson(api.RESPONSE_STATUS_OK, resultMsg, contract.Id)
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()

}

//demo使用---------------------------------------------------------------------------------------------------------------
// QueryOutput GET
func (c *ContractController) QueryOutput() {
	cost_start := time.Now()

	contractId := c.GetString(api.REQUEST_FIELD_CONTRACT_ID)
	resultMsg := fmt.Sprintf("%s 查询成功!", "API[QueryOutput]")

	if len(contractId) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryOutput]", "contractId")
		c.responseJson(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}

	output, err := rethinkdb.QueryOutput(contractId)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(Id=%s)查询错误! ", "API[QueryOutput]", contractId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg+err.Error())()
		return
	}
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(output))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()

}

// QueryOutputNum GET
func (c *ContractController) QueryOutputNum() {
	cost_start := time.Now()
	contractId := c.GetString(api.REQUEST_FIELD_CONTRACT_ID)
	resultMsg := fmt.Sprintf("%s 查询成功!", "API[QueryOutputNum]")

	if len(contractId) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryOutputNum]", "contractId")
		c.responseJson(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}

	count, err := rethinkdb.QueryOutputNum(contractId)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(Id=%s)查询错误! ", "API[QueryOutputNum]", contractId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg+err.Error())()
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"count":%d}`, count)))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

// QueryOutputDuration GET
func (c *ContractController) QueryOutputDuration() {
	cost_start := time.Now()
	contractId := c.GetString(api.REQUEST_FIELD_CONTRACT_ID)
	resultMsg := fmt.Sprintf("%s 查询成功!", "API[QueryOutputDuration]")

	if len(contractId) == 0 {
		resultMsg = fmt.Sprintf("%s %s 值错误!", "API[QueryOutputDuration]", "contractId")
		c.responseJson(api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_PARAMETER_ERROR_VALUE, resultMsg)()
		return
	}

	startTime, err := rethinkdb.QueryContractStartTime(contractId)
	if err != nil {
		resultMsg = fmt.Sprintf("%s(Id=%s)查询错误! ", "API[QueryOutputDuration]", contractId)
		uniledgerlog.Error(resultMsg)
		c.responseProto(api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg, "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg+err.Error())()
		return
	}

	nowTime := common.GenTimestamp()

	start, err := strconv.Atoi(startTime)
	now, err := strconv.Atoi(nowTime)

	hours := ((now - start) / 1000) / 3600

	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"duration":%d}`, hours)))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

// QueryAccountBalance GET
func (c *ContractController) QueryAccountBalance() {
	cost_start := time.Now()
	resultMsg := fmt.Sprintf("%s 查询成功!", "API[QueryAccountBalance]")

	result, err := function.FuncQueryAccountBalance()
	if err != nil {
		resultMsg = fmt.Sprintf("%s查询错误! ", "API[QueryOutputDuration]")
		c.responseJson(api.RESPONSE_STATUS_DB_ERROR_OP, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}

	data, _ := result.GetData().(string)

	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(data))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

// QueryAmmeterBalance GET
func (c *ContractController) QueryAmmeterBalance() {
	cost_start := time.Now()
	resultMsg := fmt.Sprintf("%s 查询成功!", "API[QueryAmmeterBalance]")

	result, err := function.FuncQueryAmmeterBalance()
	if err != nil {
		resultMsg = fmt.Sprintf("%s查询错误! ", "API[QueryOutputDuration] ")
		c.responseJson(api.RESPONSE_STATUS_DB_ERROR_OP, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}

	data, _ := result.GetData().(string)

	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(data))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

// QueryRecords GET
func (c *ContractController) QueryRecords() {
	cost_start := time.Now()
	resultMsg := fmt.Sprintf("%s 查询成功!", "API[QueryRecords]")

	str, err := rethinkdb.GetTransactionRecords()
	if err != nil {
		resultMsg = fmt.Sprintf("%s查询错误! ", "API[QueryRecords] ")
		c.responseJson(api.RESPONSE_STATUS_DB_ERROR_OP, err.Error(), "")
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg+err.Error())()
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(str))))
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

func (c *ContractController) Terminate() {
	cost_start := time.Now()
	resultMsg := fmt.Sprintf("%s 操作成功!", "API[Terminate]")
	contractProductId := c.GetString(api.REQUEST_FIELD_CONTRACT_PRODUCT_ID)
	contractId := c.GetString(api.REQUEST_FIELD_CONTRACT_ID)
	//todo
	ok := true
	if !ok {
		resultMsg = fmt.Sprintf("%s操作失败[%s]! ", "API[Terminate] ", contractId)
		c.responseJson(api.RESPONSE_STATUS_ERROR, "", resultMsg)
		defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_DB_ERROR_OP, resultMsg)()
		return
	}
	uniledgerlog.Warn("API[Terminate]contractProductId: " + contractProductId + "contractId: " + contractId)
	c.responseJson(api.RESPONSE_STATUS_OK, resultMsg, "terminate success")
	defer api.TimeCost(cost_start, c.Ctx, api.RESPONSE_STATUS_OK, resultMsg)()
}

//demo使用---------------------------------------------------------------------------------------------------------------
