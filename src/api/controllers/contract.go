package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/golang/protobuf/proto"
	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/core"
	"unicontract/src/core/db/rethinkdb"
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
	//requestDataType := c.Ctx.Input.Header("RequestDataType")
	token = c.Ctx.Input.Header("token")
	if token == "" {
		err = fmt.Errorf("服务器拒绝请求")
		status = HTTP_STATUS_CODE_Forbidden
		return
	}

	requestBody := c.Ctx.Input.RequestBody
	contract = &protos.Contract{}
	// return err init
	if contentType == "application/x-protobuf" {
		err = proto.Unmarshal(requestBody, contract)
		if err != nil {
			logs.Error("contract parseRequestBody unmarshal err ", err)
			err = fmt.Errorf("服务器拒绝请求")
			status = HTTP_STATUS_CODE_BadRequest
			return
		}

		if contract == nil || contract.Id == "" {
			err = fmt.Errorf("contract error")
			status = HTTP_STATUS_CODE_BadRequest
			return
		}
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
	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract 非法")
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

// @Title Query
// @Description get contract by cid
// @Param	body		body 	interface{}	true			"The key for contract"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /query [post]
func (c *ContractController) Query() {
	// java client StringEntity JsonStr
	// c.Ctx.Request.Body //java client httpEntity
	// c.Ctx.Request.Form //postman choose this
	//logs.Warn("Form: ", c.Ctx.Request.Form) // here map[]

	var requestParamMap map[string]interface{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &requestParamMap)

	token := c.Ctx.Request.Header.Get("token")
	/*------------------- requestParams start ------------------*/
	contractId, _ := requestParamMap["contractId"].(string)
	owner, _ := requestParamMap["owner"].(string)
	signatureStatus, _ := requestParamMap["signatureStatus"].(string)
	flag, _ := requestParamMap["flag"].(string)
	/*------------------- requestParams end ------------------*/
	logs.Warn("Body: ", c.Ctx.Request.Body)
	//logs.Warn("Header: ", c.Ctx.Request.Header)

	logs.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, signatureStatus=%s, flag=%s, contractId=%s]",
		c.Ctx.Request.RequestURI, token, owner, signatureStatus, flag, contractId))
	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}
	if contractId == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractId is blank!")
		return
	}
	//todo
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
	contractId, ok := requestParamMap["contractId"].(string)
	if !ok {
		logs.Error("contractId type error")
	}
	owner, _ := requestParamMap["owner"].(string)
	signatureStatus, _ := requestParamMap["signatureStatus"].(string)
	name, _ := requestParamMap["name"].(string)
	/*------------------- requestParams end ------------------*/
	logs.Warn("Body: ", c.Ctx.Request.Body)
	//logs.Warn("Header: ", c.Ctx.Request.Header)

	logs.Warn(fmt.Sprintf("[API] match |%s [token =%s, owner =%s, signatureStatus=%s, name=%s, contractId=%s]",
		c.Ctx.Request.RequestURI, token, owner, signatureStatus, name, contractId))
	// todo
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

// @Title Track
// @Description track contract by contract Id
// @Param	cid		path 	string	true		"The key for contract"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /track [post]
func (c *ContractController) Track() {
	_, contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(status, "", false, err.Error())
		return
	}

	contractId := contract.Id
	if contractId == "" {
		logs.Error("请输入合约Id")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "错误请求!")
		return
	}
	contractTasksStr, err := rethinkdb.GetContractTasksByContractId(contractId)
	if err != nil {
		logs.Error("API[Track] GetContractTasksByContractId 查询失败 ", contractId)
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	if contractTasksStr == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "contractTask 不存在")
		return
	}

	var contractTasks []model.ContractTask
	err = json.Unmarshal([]byte(contractTasksStr), &contractTasks)

	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "Unmarshal ContractTasks fail!")
		return
	}

	//TODO track contract 合约跟踪
	logs.Warn(c.Ctx.Request.RequestURI, "API[Track] 缺少合约跟踪查询方法!")
	c.responseJsonBody(contract.Id, false, "API[Track] 缺少合约跟踪查询方法!")
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
	token, contract, err, status := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(status, "", false, err.Error())
		return
	}

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
