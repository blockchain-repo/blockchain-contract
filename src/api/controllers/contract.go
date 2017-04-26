package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	"time"
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

func (c *ContractController) parseProtoRequestBody() (token string, contract *protos.Contract, err error) {

	contentType := c.Ctx.Input.Header("Content-Type")
	requestDataType := c.Ctx.Input.Header("RequestDataType")
	token = c.Ctx.Input.Header("token")

	requestBody := c.Ctx.Input.RequestBody
	contract = &protos.Contract{}

	// return err init
	if requestDataType == "proto" && (contentType == "application/json" || contentType == "application/x-protobuf") {
		err := proto.Unmarshal(requestBody, contract)
		if err != nil {
			beego.Error("contract parseRequestBody unmarshal err ", err)
		}
		beego.Debug("[API] Request for contract[Content-type=" + contentType + "]")
		beego.Debug("[API] token is", token)
		beego.Debug("[API] contract content as follows:\n", contract)
		//fmt.Println(contract)
	}
	return token, contract, err
}

//response the json body
//serializeStrData
func (c *ContractController) responseJsonBody(data string, ok bool, msg string) {
	responseData := new(protos.ResponseData)
	responseData.Ok = ok
	responseData.Msg = msg
	responseData.Data = data
	//body, _ := json.Marshal(responseData)
	body, err := proto.Marshal(responseData)
	if err != nil {
		beego.Error("responseJsonBodyCode ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(body))
	//c.ServeJSON()
}

func (c *ContractController) responseJsonBodyCode(status int, data string, ok bool, msg string) {
	responseData := new(protos.ResponseData)
	responseData.Ok = ok
	responseData.Msg = msg
	responseData.Data = data
	body, err := proto.Marshal(responseData)
	if err != nil {
		beego.Error("responseJsonBodyCode ", err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(status)
	c.Ctx.ResponseWriter.Write([]byte(body))
}

func (c *ContractController) responseWithCode(status int, data string) {
	responseData := new(protos.ResponseData)
	responseData.Ok = true
	responseData.Msg = ""
	responseData.Data = data
	body, err := proto.Marshal(responseData)
	if err != nil {
		beego.Error("responseJsonBodyCode ", err.Error())
	}
	// last panic user string
	c.Ctx.ResponseWriter.WriteHeader(status)
	c.Ctx.ResponseWriter.Write([]byte(body))
}

// @Title AuthSignature
// @Description AuthSignature for contract
// @Param	body		body 	models.Contract	true		"body for contract content"
// @Success 200 {} models.Contract.Head.Id
// @Failure 403 body is empty
// @router /authSignature [post]
func (c *ContractController) AuthSignature() {
	token, contract, err := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "服务器拒绝请求")
		return
	}

	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	contractModel := fromContractToContractModel(*contract)
	signatureValid := contractModel.IsSignatureValid()
	if !signatureValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}
	//TODO func authSignature(...) 验证签名方法
	beego.Debug("Token is " + token)
	beego.Debug("contract signature is valid ", signatureValid)
	c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "验证签名 success")
}

// API receive and transfer it to contractModel
func fromContractToContractModel(contract protos.Contract) model.ContractModel {
	var contractModel model.ContractModel
	contractModel.Contract = contract
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
		beego.Error("error fromContractModelStrToContract", err)
		return contract, err
	}

	return contract, nil
}

// @Title CreateContract
// @Description create contract
// @Param	body		body 	models.Contract	true		"body for contract content"
// @Success 200 {int} models.Contract.Head.Id
// @Failure 403 body is empty
// @router /create [post]
func (c *ContractController) Create() {
	token, contract, err := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "服务器拒绝请求")
		return
	}

	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	if contract == nil || contract.Id == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract error!")
		return
	}

	contractModel := fromContractToContractModel(*contract)
	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract 非法")
		beego.Debug("API[Create] token is", token)
		return
	}
	ok := core.WriteContract(contractModel)
	if !ok {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "API[Create] insert contract fail!")
		beego.Debug(c.Ctx.Request.RequestURI, "API[Create] insert contract fail!")
		return
	}
	beego.Info(c.Ctx.Request.RequestURI, "API[Create] Id:"+contractModel.Id)
	c.responseJsonBody(contract.Id, true, "API[Create] insert contract Id "+contractModel.Id+"]")

}

// @Title Signature
// @Description signature the contract
// @Param	body		body 	interface{}	true		"body for contract id"
// @Success 200 {object} models.Contract
// @Failure 403 body is empty
// @router /signature [post]
func (c *ContractController) Signature() {
	token, contract, err := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "服务器拒绝请求")
		return
	}

	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	if contract == nil || contract.Id == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract error!")
		return
	}

	contractModel := fromContractToContractModel(*contract)
	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract 非法")
		beego.Debug("API[Signature] token is", token)
		return
	}
	ok := core.WriteContract(contractModel)
	if !ok {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "API[Signature] insert contract fail!")
		beego.Debug(c.Ctx.Request.RequestURI, "API[Signature] insert contract fail!")
		return
	}
	beego.Info(c.Ctx.Request.RequestURI, "API[Signature] Id:"+contractModel.Id)
	c.responseJsonBody(contract.Id, true, "API[Signature] insert contract Id "+contractModel.Id+"]")
}

// @Title Terminate
// @Description terminate the contract
// @Param	body		body 	interface{}	true		"body for contract id"
// @Success 200 {string} terminate success!
// @Failure 403 body is empty
// @router /terminate [post]
func (c *ContractController) Terminate() {
	token, contract, err := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "服务器拒绝请求")
		return
	}

	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	if contract == nil || contract.Id == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract error!")
		return
	}

	if contract.Id == "" {
		beego.Debug("API[Terminate]合约(Id=" + contract.Id + ")不存在: ")
		c.responseJsonBody("", false, "合约终止失败!")
		return
	}

	beego.Warn("Token is " + token)
	beego.Warn(c.Ctx.Request.RequestURI, "API[Signature]缺少终止合约方法!")
	beego.Warn("API[Terminate]合约Id: " + contract.Id)
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
	token, contract, err := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "服务器拒绝请求")
		return
	}

	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	if contract == nil || contract.Id == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "请输入合约Id!")
		return
	}

	beego.Debug("Token is " + token)

	contractModelStr, err := rethinkdb.GetContractById(contract.Id)
	if err != nil {
		beego.Error("API[Query]合约(Id=" + contract.Id + ")查询错误: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[Query]合约查询错误!")
		return
	}

	if contractModelStr == "" {
		beego.Error("API[Query]合约(Id=" + contract.Id + ")不存在: ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "API[Query]合约(Id="+contract.Id+")不存在: ")
		return
	}

	contractProto, err := fromContractModelStrToContract(contractModelStr)
	if err != nil {
		beego.Error("API[Query]合约(Id=" + contract.Id + "), 转换失败(fromContractModelStrToContract)")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	contractProtoBytes, err := proto.Marshal(&contractProto)
	if err != nil {
		beego.Error("API[Query]合约(Id=" + contract.Id + "), 转换失败(proto.Marshal) ")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	contractProtoStr := string(contractProtoBytes)
	c.responseJsonBody(contractProtoStr, true, "API[Query]查询合约成功!")
}

// @Title Track
// @Description track contract by contract Id
// @Param	cid		path 	string	true		"The key for contract"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /track [post]
func (c *ContractController) Track() {
	token, contract, err := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "服务器拒绝请求")
		return
	}

	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	if contract == nil || contract.Id == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "请输入合约Id!")
		return
	}

	beego.Debug("Token is " + token)

	contractId := contract.Id
	if contractId == "" {
		beego.Error("请输入合约Id")
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "错误请求!")
		return
	}
	contractTasksStr, err := rethinkdb.GetContractTasksByContractId(contractId)
	if err != nil {
		beego.Error("API[Track] GetContractTasksByContractId 查询失败 ", contractId)
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
	beego.Debug("Token is " + token)
	beego.Info(c.Ctx.Request.RequestURI, "API[Track] 缺少合约跟踪查询方法!")
	c.responseJsonBody(contract.Id, false, "API[Track] 缺少合约跟踪查询方法!")

}

// @Title Update
// @Description update the contract
// @Param	body		body 	models.Contract	true		"body for contract content"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /update [post]
func (c *ContractController) Update() {
	token, contract, err := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "服务器拒绝请求")
		return
	}

	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	if contract == nil || contract.Id == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract error!")
		return
	}

	contractModel := fromContractToContractModel(*contract)
	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract Validate error")
		return
	}

	//ok := rethinkdb.InsertContract(common.Serialize(contractModel))
	beego.Debug("Token is " + token)
	beego.Info(c.Ctx.Request.RequestURI, "API[Update] 缺少测试合约方法!")
	c.responseJsonBody(contract.Id, false, "API[Update] 缺少合约更新方法!")
}

// @Title Test
// @Description test the contract
// @Param	cid		path 	string	true		"The uid you want to test"
// @Success 200 {string} test success!
// @Failure 403 cid is empty
// @router /test [post]
func (c *ContractController) Test() {
	token, contract, err := c.parseProtoRequestBody()
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "服务器拒绝请求")
		return
	}

	if token == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_Forbidden, "", false, "服务器拒绝请求")
		return
	}

	if contract == nil || contract.Id == "" {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract error!")
		return
	}

	beego.Debug("Token is " + token)
	beego.Info(c.Ctx.Request.RequestURI, "API[Test] 缺少测试合约方法!")
	c.responseJsonBody(string(time.Now().Unix()), false, "API[Test] 缺少测试合约方法!")

}
