package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/golang/protobuf/proto"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
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

const (
	HTTP_STATUS_CODE_OK             = 200 //200 - 客户端请求已成功
	HTTP_STATUS_CODE_BadRequest     = 400 //400 - 请求出现语法错误
	HTTP_STATUS_CODE_Unauthorized   = 401 //401 - 访问被拒绝
	HTTP_STATUS_CODE_Forbidden      = 403 //403 - 禁止访问 资源不可用
	HTTP_STATUS_CODE_NotFound       = 404 //404 - 无法找到指定位置的资源
	HTTP_STATUS_CODE_NotAcceptable  = 406 //406 - 指定的资源已经找到，但它的MIME类型和客户在Accpet头中所指定的不兼容
	HTTP_STATUS_CODE_RequestTimeout = 408 //408 - 在服务器许可的等待时间内，客户一直没有发出任何请求。客户可以在以后重复同一请求。
)

func (c *ContractController) parseProtoRequestBody() (token string, contract *protos.Contract, err error, status int) {
	contentType := c.Ctx.Input.Header("Content-Type")
	token = c.Ctx.Input.Header("token")
	if token == "" {
		err = fmt.Errorf("token is blank!")
		status = HTTP_STATUS_CODE_BadRequest
		return
	}

	requestBody := c.Ctx.Input.RequestBody
	contract = &protos.Contract{}
	// return err init
	if contentType == "application/x-protobuf" {
		err = proto.Unmarshal(requestBody, contract)
		if err != nil {
			logs.Error("contract parseRequestBody unmarshal err ", err)
			err = fmt.Errorf("contract parseRequestBody unmarshal err ")
			status = HTTP_STATUS_CODE_BadRequest
			return
		}
		//todo temp
		logs.Warn(contract)
		//if contract == nil || contract.Id == "" {
		//	err = fmt.Errorf("contract nil or contract.Id is blank!")
		//	status = HTTP_STATUS_CODE_BadRequest
		//	return
		//}
		fmt.Sprintf("[API] match |%s [token =%s, Content-Type =%s]", token, c.Ctx.Request.RequestURI,
			contentType)
		logs.Info(fmt.Sprintf("[API] match|%-32s \t[token = %s, Content-Type = %s]", c.Ctx.Request.RequestURI,
			c.Ctx.Request.Method, contentType))
	}
	return
}

//response the json body
//serializeStrData
func (c *ContractController) responseJsonBody(data string, ok bool, msg string) {
	responseData := new(protos.ResponseData)
	responseData.Ok = ok
	responseData.Msg = msg
	data = base64.StdEncoding.EncodeToString([]byte(data))
	responseData.Data = data
	//body, _ := json.Marshal(responseData)
	body, err := proto.Marshal(responseData)
	if err != nil {
		logs.Error("responseJsonBodyCode ", err.Error())
	}
	//c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/x-protobuf")
	c.Ctx.ResponseWriter.WriteHeader(HTTP_STATUS_CODE_OK)
	//c.Ctx.ResponseWriter.Write([]byte(body))
	c.Ctx.ResponseWriter.Write(body)
	//c.ServeJSON()
}

func (c *ContractController) responseJsonBodyCode(status int, data string, ok bool, msg string) {
	responseData := new(protos.ResponseData)
	responseData.Ok = ok
	responseData.Msg = msg
	//todo test
	data = base64.StdEncoding.EncodeToString([]byte(data))
	responseData.Data = data

	body, err := proto.Marshal(responseData)
	if err != nil {
		logs.Error("responseJsonBodyCode ", err.Error())
	}
	//c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/x-protobuf")
	c.Ctx.ResponseWriter.WriteHeader(status)
	c.Ctx.ResponseWriter.Write([]byte(body))
}

func (c *ContractController) responseWithCode(status int, data string) {
	responseData := new(protos.ResponseData)
	responseData.Ok = true
	responseData.Msg = ""
	//todo test
	data = base64.StdEncoding.EncodeToString([]byte(data))
	responseData.Data = data
	body, err := proto.Marshal(responseData)
	if err != nil {
		logs.Error("responseJsonBodyCode ", err.Error())
	}
	// last panic user string
	//c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/x-protobuf")
	c.Ctx.ResponseWriter.WriteHeader(status)
	c.Ctx.ResponseWriter.Write([]byte(body))
}

// API receive and transfer it to contractModel
func fromContractToContractModel(contract *protos.Contract) model.ContractModel {
	var contractModel model.ContractModel
	contractModel.Contract = *contract
	return contractModel
}

// go rethink get contractModel string and transfer it to contract
func fromContractModelStrToContract(contractModelStr string) (protos.Contract, error) {
	// 1. to contractModel
	var contractModel model.ContractModel
	err := json.Unmarshal([]byte(contractModelStr), &contractModel)
	// 2. to contract
	contract := contractModel.Contract
	if err != nil {
		logs.Error("error fromContractModelStrToContract", err)
		return contract, err
	}

	return contract, nil
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
		logs.Error("error fromContractModelArrayStrToContracts", err)
		return contractList, err
	}
	contracts = make([]*protos.Contract, len(contractModel))
	for i := 0; i < len(contractModel); i++ {
		contracts[i] = &contractModel[i].Contract
	}
	contractList.Contracts = contracts
	logs.Info("query contract len is ", len(contractModel))
	return contractList, nil
}

// special for contractArray to proto[] only for queryLog
func fromContractModelArrayStrToContractsForLog(contractModelStr string) (protos.ContractList, error) {
	// 1. to contractModel
	var contractModel []model.ContractModel
	err := json.Unmarshal([]byte(contractModelStr), &contractModel)
	// 2. to contract
	var contractList protos.ContractList
	var contracts []*protos.Contract
	if err != nil {
		logs.Error("error fromContractModelArrayStrToContracts", err)
		return contractList, err
	}
	contracts = make([]*protos.Contract, len(contractModel))
	for i := 0; i < len(contractModel); i++ {
		tempContract := &contractModel[i].Contract
		contracts[i] = &protos.Contract{
			Id: tempContract.Id,
			ContractHead: &protos.ContractHead{
				OperateTime: tempContract.ContractHead.OperateTime,
			},
			ContractBody: &protos.ContractBody{
				Caption:            tempContract.ContractBody.Caption,
				Cname:              tempContract.ContractBody.Cname,
				ContractId:         tempContract.ContractBody.ContractId,
				ContractOwners:     tempContract.ContractBody.ContractOwners,
				ContractSignatures: tempContract.ContractBody.ContractSignatures,
				ContractState:      tempContract.ContractBody.ContractState,
				Description:        tempContract.ContractBody.Description,
				StartTime:          tempContract.ContractBody.StartTime,
				EndTime:            tempContract.ContractBody.EndTime,
				Creator:            tempContract.ContractBody.Creator,
			},
		}
	}
	contractList.Contracts = contracts
	logs.Info("query contract len is ", len(contractModel))
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
		logs.Error("error fromContractOutputsModelArrayStrToContractsForLog", err)
		return contractExecuteLogList, err
	}
	contractExecuteLogs = make([]*protos.ContractExecuteLog, len(contractOutput))
	for i := 0; i < len(contractOutput); i++ {
		tempTransaction := contractOutput[i].Transaction
		tempRelation := tempTransaction.Relation
		tempContractBody := tempTransaction.ContractModel.ContractBody
		taskId := tempRelation.TaskId
		if taskId == "" {
			logs.Error("taskId is blank!", err)
			return contractExecuteLogList, err
		}
		tempContractComponents := tempContractBody.ContractComponents
		var tempContractComponent protos.ContractComponent
		for j := 0; j < len(tempContractComponents); j++ {
			if tempContractComponents[j].TaskId == taskId {
				tempContractComponent = *tempContractComponents[j]
				break
			}
		}

		logs.Error(tempContractComponent)
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
	logs.Info("query contractExecuteLogs len is ", len(contractExecuteLogs))
	return contractExecuteLogList, nil
}

// @Title AuthSignature
// @Description AuthSignature for contract
// @Param	body		body 	models.Contract	true		"body for contract content"
// @Success 200 {} models.Contract.Head.Id
// @Failure 403 body is empty
// @router /authSignature [post]
func (c *ContractController) AuthSignature() {
	_, contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(status, "", false, err.Error())
		return
	}

	contractModel := fromContractToContractModel(contract)
	signatureValid := contractModel.IsSignatureValid()
	if !signatureValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "合约签名验证失败")
		return
	}
	c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "验证签名 success")
}

// @Title CreateContract
// @Description create contract
// @Param	body		body 	models.Contract	true		"body for contract content"
// @Success 200 {int} models.Contract.Head.Id
// @Failure 403 body is empty
// @router /create [post]
func (c *ContractController) Create() {
	token, contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(status, "", false, err.Error())
		monitor.Monitor.Count("request_fail", 1)
		return
	}

	contractModel := fromContractToContractModel(contract)
	logs.Warn("contractModel:\n", contractModel)
	contractModel.ContractHead = &protos.ContractHead{
		Version: 1,
	}
	//TODO 额外验证 合约基本字段、owners、component为空
	//contractHead := contractModel.ContractHead
	//contractBody := contractModel.ContractBody
	//components := contractModel.ContractBody.ContractComponents
	//if contractHead == nil || contractBody == nil || contractBody.ContractOwners == nil || components == nil {
	//	c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract 验证不通过!")
	//	logs.Debug("API[Create] token is", token)
	//	monitor.Monitor.Count("request_fail", 1)
	//	return
	//}

	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract 验证不通过!")
		logs.Debug("API[Create] token is", token)
		monitor.Monitor.Count("request_fail", 1)
		return
	}
	contract_write_time := monitor.Monitor.NewTiming()
	ok := core.WriteContract(contractModel)
	if !ok {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "API[Create] insert contract fail!")
		logs.Debug(c.Ctx.Request.RequestURI, "API[Create] insert contract fail!")
		monitor.Monitor.Count("request_fail", 1)
		return
	}
	contract_write_time.Send("contract_write")
	c.responseJsonBody(contract.Id, true, "API[Create] insert contract Id "+contractModel.Id+"]")

}

// @Title Signature
// @Description signature the contract
// @Param	body		body 	interface{}	true		"body for contract id"
// @Success 200 {object} models.Contract
// @Failure 403 body is empty
// @router /signature [post]
func (c *ContractController) Signature() {
	token, contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(status, "", false, err.Error())
		return
	}

	contractModel := fromContractToContractModel(contract)
	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract 非法")
		logs.Debug("API[Signature] token is", token)
		return
	}
	ok := core.WriteContract(contractModel)
	if !ok {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "API[Signature] insert contract fail!")
		logs.Debug(c.Ctx.Request.RequestURI, "API[Signature] insert contract fail!")
		return
	}
	c.responseJsonBody(contract.Id, true, "API[Signature] insert contract Id "+contractModel.Id+"]")
}

// @Title Terminate
// @Description terminate the contract
// @Param	body		body 	interface{}	true		"body for contract id"
// @Success 200 {string} terminate success!
// @Failure 403 body is empty
// @router /terminate [post]
func (c *ContractController) Terminate() {
	_, contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(status, "", false, err.Error())
		return
	}

	if contract.Id == "" {
		logs.Debug("API[Terminate]合约(Id=" + contract.Id + ")不存在: ")
		c.responseJsonBody("", false, "合约终止失败!")
		return
	}

	logs.Warn(c.Ctx.Request.RequestURI, "API[Signature]缺少终止合约方法!", "合约Id:", contract.Id)
	c.responseJsonBody(contract.Id, false, "API[Terminate]合约终止失败!")
	//c.responseJsonBody(contract.Id, true, "合约终止成功!")
}

// query the contract content with the contractId and oners
func (c *ContractController) QueryContractContent() {

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)
	/*------------------- requestParams end ------------------*/

	logs.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractId=%s]",
		c.Ctx.Request.RequestURI, token, owner, contractId))
	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}
	if contractId == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId is blank!")
		return
	}
	contractModelStr, err := rethinkdb.GetContractContentByMapCondition(requestParamMap)
	//logs.Warn("QueryContractContent:\n", contractModelStr)
	if err != nil {
		logs.Error("API[QueryContractContent]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[QueryContractContent]合约查询错误!")
		return
	}

	if contractModelStr == "" {
		logs.Warn("API[QueryContractContent]合约(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[QueryContractContent]合约(Id="+contractId+")不存在: ")
		return
	}

	c.responseJsonBody(contractModelStr, true, "API[QueryContractContent]查询合约成功!")
}

func (c *ContractController) QueryPublishContract() {

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractState := "Contract_Create"
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)
	/*------------------- requestParams end ------------------*/
	//logs.Warn("Body: ", c.Ctx.Request.Body)

	logs.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractState=%s, contractId=%s]",
		c.Ctx.Request.RequestURI, token, owner, contractState, contractId))
	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}
	if contractId == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId is blank!")
		return
	}
	contractModelStr, err := rethinkdb.GetPublishContractByMapCondition(requestParamMap)
	if err != nil {
		logs.Error("API[QueryPublishContract]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[QueryPublishContract]合约查询错误!")
		return
	}

	if contractModelStr == "" {
		logs.Warn("API[QueryPublishContract]合约(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[QueryPublishContract]合约(Id="+contractId+")不存在: ")
		return
	}

	c.responseJsonBody(contractModelStr, true, "API[QueryPublishContract]查询合约成功!")
}

// @Title Query
// @Description get contract by cid
// @Param	body		body 	interface{}	true			"The key for contract"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /query [post]
func (c *ContractController) Query() {

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractState, _ := requestParamMap["status"].(string)
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)
	/*------------------- requestParams end ------------------*/
	logs.Warn("Body: ", c.Ctx.Request.Body)
	//logs.Warn("Header: ", c.Ctx.Request.Header)

	logs.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractState=%s, contractId=%s]",
		c.Ctx.Request.RequestURI, token, owner, contractState, contractId))
	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}
	if contractId == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId is blank!")
		return
	}
	contractModelStr, err := rethinkdb.GetOneContractByMapCondition(requestParamMap)
	if err != nil {
		logs.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[Query]合约查询错误!")
		return
	}

	if contractModelStr == "" {
		logs.Warn("API[Query]合约(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[Query]合约(Id="+contractId+")不存在: ")
		return
	}

	contractProto, err := fromContractModelStrToContract(contractModelStr)
	if err != nil {
		logs.Error("API[Query]合约(Id=" + contractId + "), 转换失败(fromContractModelStrToContract)")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	contractProtoBytes, err := proto.Marshal(&contractProto)
	if err != nil {
		logs.Error("API[Query]合约(Id=" + contractId + "), 转换失败(proto.Marshal) ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	contractProtoStr := string(contractProtoBytes)
	c.responseJsonBody(contractProtoStr, true, "API[Query]查询合约成功!")
	//c.responseJsonBody(contractProtoStr, true, "API[Query] success!")
}

// @Title Query
// @Description get contract by cid
// @Param	body		body 	interface{}	true			"The key for contract"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /queryList [post]
func (c *ContractController) QueryAll() {
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractState, _ := requestParamMap["status"].(string)
	owner, _ := requestParamMap["owner"].(string)

	contractId, _ := requestParamMap["contractId"].(string)
	//if !ok {
	//	logs.Error("contractId type error")
	//}
	contractName, _ := requestParamMap["contractName"].(string)
	/*------------------- requestParams end ------------------*/
	logs.Warn("Body: ", c.Ctx.Request.Body)
	//logs.Warn("Header: ", c.Ctx.Request.Header)

	logs.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractState=%s, contractId=%s, contractName=%s]",
		c.Ctx.Request.RequestURI, token, owner, contractState, contractId, contractName))
	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	contractModelStr, err := rethinkdb.GetContractsByMapCondition(requestParamMap)
	if err != nil {
		logs.Error("API[Query]合约(Id="+contractId+")查询错误: ", err)
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[Query]合约查询错误!")
		return
	}

	if contractModelStr == "" {
		logs.Warn("API[Query]合约(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[Query]合约(Id="+contractId+")不存在: ")
		return
	}

	contractListProto, err := fromContractModelArrayStrToContracts(contractModelStr)
	if err != nil {
		logs.Error("API[Query]合约(Id=" + contractId + "), 转换失败(fromContractModelStrToContract)")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	contractListProtoBytes, err := proto.Marshal(&contractListProto)
	if err != nil {
		logs.Error("API[QueryALl]合约, 转换失败(proto.Marshal) ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	contractProtoStr := string(contractListProtoBytes)
	c.responseJsonBody(contractProtoStr, true, "API[Query]查询合约成功!")
	//c.responseJsonBody(contractProtoStr, true, "API[Query] success!")
}

func (c *ContractController) QueryLog() {
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractState, _ := requestParamMap["status"].(string)
	owner, _ := requestParamMap["owner"].(string)

	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		logs.Error("contractId type error")
	}
	contractName, _ := requestParamMap["contractName"].(string)
	/*------------------- requestParams end ------------------*/
	logs.Warn("Body: ", c.Ctx.Request.Body)
	//logs.Warn("Header: ", c.Ctx.Request.Header)

	logs.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, contractState=%s, contractId=%s, contractName=%s]",
		c.Ctx.Request.RequestURI, token, owner, contractState, contractId, contractName))
	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	contractOutputsModelStr, err := rethinkdb.GetContractsLogByMapCondition(requestParamMap)

	if err != nil {
		logs.Error("API[QueryLog]合约log(Id="+contractId+")查询错误: ", err)
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[QueryLog]查询错误!")
		return
	}

	if contractOutputsModelStr == "" {
		logs.Warn("API[QueryLog]合约log(Id=" + contractId + ")不存在: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[QueryLog](Id="+contractId+")不存在: ")
		return
	}
	//todo 需要过滤字段,只提取需要的字段!
	contractExecuteLogListProto, err := fromContractOutputsModelArrayStrToContractsForLog(contractOutputsModelStr)
	logs.Warn(contractExecuteLogListProto)
	if err != nil {
		logs.Error("API[QueryLog]合约(Id=" + contractId + "), 转换失败(fromContractOutputsModelArrayStrToContractsForLog)")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	contractExecuteLogListProtoBytes, err := proto.Marshal(&contractExecuteLogListProto)
	if err != nil {
		logs.Error("API[QueryLog]合约, 转换失败(proto.Marshal) ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	contractExecuteLogListProtoStr := string(contractExecuteLogListProtoBytes)
	c.responseJsonBody(contractExecuteLogListProtoStr, true, "API[QueryLog]查询成功!")
	//c.responseJsonBody(contractProtoStr, true, "API[Query] success!")
}

// @Title Update
// @Description update the contract
// @Param	body		body 	models.Contract	true		"body for contract content"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /update [post]
func (c *ContractController) Update() {
	_, contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(status, "", false, err.Error())
		return
	}

	contractModel := fromContractToContractModel(contract)
	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract Validate error")
		return
	}
	//TODO track contract 缺少测试合约方法
	//ok := rethinkdb.InsertContract(common.StructSerialize(contractModel))
	logs.Warn(c.Ctx.Request.RequestURI, "API[Update] 缺少测试合约方法!")
	c.responseJsonBody(contract.Id, false, "API[Update] 缺少合约更新方法!")
}

// @Title Test
// @Description test the contract
// @Param	cid		path 	string	true		"The uid you want to test"
// @Success 200 {string} test success!
// @Failure 403 cid is empty
// @router /test [post]
func (c *ContractController) Test() {
	_, _, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(status, "", false, err.Error())
		return
	}

	//TODO track contract 缺少测试合约方法
	logs.Warn(c.Ctx.Request.RequestURI, "API[Test] 缺少测试合约方法!")
	c.responseJsonBody(string(time.Now().Unix()), false, "API[Test] 缺少测试合约方法!")
}

// for press test [pressTest]
func (c *ContractController) PressTest() {
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	startTime, _ := requestParamMap["start"].(string)
	endTime, _ := requestParamMap["end"].(string)
	token, contract, err, status := c.parseProtoRequestBody()
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

	if err != nil {
		c.responseJsonBodyCode(status, "", false, err.Error())
		return
	}

	//todo 1. replace createTime, Signatures, owner, start and end time!

	//logs.Warn("Input contractDeserialize:\n", common.StructSerialize(contract))
	contractModel := fromContractToContractModel(contract)
	/*-------------------------- this for press test generate Id start---------------------*/
	// add random string
	randomString := common.GenerateUUID() + "_node" + c.Ctx.Request.RequestURI + "_token_" + token
	contractModel.ContractBody.Caption = randomString
	contractModel.ContractBody.Description = randomString

	contractOwnersLen := 3
	// 生成的合约签名人个数
	contractSignaturesLen := contractOwnersLen

	if contractSignaturesLen >= contractOwnersLen || contractSignaturesLen <= 0 {
		contractSignaturesLen = contractOwnersLen
	}

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
	contractSignatures := make([]*protos.ContractSignature, contractSignaturesLen)
	for i := 0; i < contractSignaturesLen; i++ {
		ownerPubkey := ownersPubkeys[i]
		privateKey := owners[ownerPubkey]
		contractSignatures[i] = &protos.ContractSignature{
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
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract 非法")
		logs.Debug("API[PressTest] token is", token)
		return
	}
	ok := core.WriteContract(contractModel)
	if !ok {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "API[PressTest] insert contract fail!")
		logs.Debug(c.Ctx.Request.RequestURI, "API[PressTest] insert contract fail!")
		return
	}
	logs.Warn("API[PressTest] InsertContract success!")
	c.responseJsonBody(contract.Id, true, "API[PressTest] insert contract Id "+contractModel.Id+"]")

}

//demo使用---------------------------------------------------------------------------------------------------------------
func (c *ContractController) QueryOutput() {
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	if len(token) == 0 {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId type is error!")
		return
	}
	if len(contractId) == 0 {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId is blank!")
		return
	}

	output, err := rethinkdb.QueryOutput(contractId)
	if err != nil {
		logs.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[Query]合约查询错误!")
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(output))))

}

func (c *ContractController) QueryOutputNum() {
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	if len(token) == 0 {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId type is error!")
		return
	}
	if len(contractId) == 0 {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId is blank!")
		return
	}

	count, err := rethinkdb.QueryOutputNum(contractId)
	if err != nil {
		logs.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[QueryOutputNum]合约查询错误!")
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"count":%d}`, count)))))
}

func (c *ContractController) QueryOutputDuration() {
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	if len(token) == 0 {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId type is error!")
		return
	}
	if len(contractId) == 0 {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId is blank!")
		return
	}

	startTime, err := rethinkdb.QueryContractStartTime(contractId)
	if err != nil {
		logs.Error("API[Query]合约(Id=" + contractId + ")查询错误: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[QueryOutputDuration]合约查询错误!")
		return
	}

	nowTime := common.GenTimestamp()

	start, err := strconv.Atoi(startTime)
	now, err := strconv.Atoi(nowTime)

	hours := ((now - start) / 1000) / 3600

	c.Ctx.ResponseWriter.WriteHeader(HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"duration":%d}`, hours)))))
}

func (c *ContractController) QueryAccountBalance() {
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	result, err := function.FuncQueryAccountBalance()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}

	data, _ := result.GetData().(string)

	c.Ctx.ResponseWriter.WriteHeader(HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(data))))
}

func (c *ContractController) QueryAmmeterBalance() {
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	result, err := function.FuncQueryAmmeterBalance()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}

	data, _ := result.GetData().(string)

	c.Ctx.ResponseWriter.WriteHeader(HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(data))))
}

func (c *ContractController) QueryRecords() {
	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	if len(token) == 0 {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	str, err := rethinkdb.GetTransactionRecords()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(base64.StdEncoding.EncodeToString([]byte(str))))
}

//demo使用---------------------------------------------------------------------------------------------------------------
