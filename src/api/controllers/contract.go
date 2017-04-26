package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	"time"
	"unicontract/src/common"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"
	"unicontract/src/core"
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
		beego.Debug("Request for contract[Content-type=" + contentType + "]")
		beego.Debug("contract content as follows:\n", contract)
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
		beego.Error("responseJsonBodyCode ",err.Error())
	}
	c.Ctx.ResponseWriter.WriteHeader(HTTP_STATUS_CODE_OK)
	c.Ctx.ResponseWriter.Write([]byte(body))
	beego.Error(responseData)
	//todo
	//c.ServeJSON()
}

func (c *ContractController) responseJsonBodyCode(status int, data string, ok bool, msg string) {
	responseData := new(protos.ResponseData)
	responseData.Ok = ok
	responseData.Msg = msg
	responseData.Data = data
	//body, _ := json.Marshal(responseData)
	body, err := proto.Marshal(responseData)
	if err != nil {
		beego.Error("responseJsonBodyCode ",err.Error())
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
		beego.Error("responseJsonBodyCode ",err.Error())
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

	beego.Error(common.SerializePretty(contractModel))

	return contractModel
}

// go rethink get contractModel string and transfer it to contract
func fromContractModelStrToContract(contractModelStr string) (protos.Contract, error) {
	// 1. to contractModel
	var contractModel model.ContractModel
	err := json.Unmarshal([]byte(contractModelStr), &contractModel)
	if err != nil {
		beego.Error("error fromContractModelStrToContract", err)
	}
	beego.Error("new ", common.SerializePretty(contractModel))

	// 2. to contract
	contract := contractModel.Contract

	return contract, err
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

	beego.Debug("Token is " + token)

	contractModel := fromContractToContractModel(*contract)
	//contractValid := contractModel.Validate() //todo
	//contractValid := true
	//if !contractValid {
	//	c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract error")
	//	return
	//}
	//beego.Warn(contractModel)
	//beego.Warn(contractModel.Id)
	//beego.Warn(contractModel.GenerateId())
	core.WriteContract(contractModel)
	//ok := rethinkdb.InsertContract(common.Serialize(contractModel))
	//response := make(map[string]interface{})

	beego.Info(c.Ctx.Request.RequestURI, "API Insert![Create Id:"+contractModel.Id+"]")
	//response["id"] = contractModel.Id
	c.responseJsonBody(contract.Id, true, "API Insert![Create Id:"+contractModel.Id+"]")

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

	beego.Debug("Token is " + token)

	contractModel := fromContractToContractModel(*contract)
	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract error")
		return
	}

	rethinkdb.InsertContract(common.Serialize(contractModel))
	beego.Info(c.Ctx.Request.RequestURI, "API Signature![Create Id:"+contractModel.Id+"]")

	c.responseJsonBody(contractModel.Id, true, "API Signature![Create Id:"+contractModel.Id+"]")
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
		beego.Error("合约(Id=" + contract.Id + ")不存在: ")
		c.responseJsonBody("", false, "合约终止失败!")
		return
	}

	beego.Debug("Token is " + token)

	//beego.Warn("缺少查询合约方法![Terminate]")
	beego.Warn(c.Ctx.Request.RequestURI, "缺少终止合约方法![Terminate]")
	beego.Warn("合约Id: " + contract.Id)

	beego.Info("合约(Id=" + contract.Id + ")存在: ")
	c.responseJsonBody(contract.Id, true, "合约终止成功!")
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
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, "查询错误!")
		return
	}

	if contractModelStr == "" {
		beego.Error("合约(Id=" + contract.Id + ")不存在: ")
	}

	contractProto, err := fromContractModelStrToContract(contractModelStr)
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	contractProtoBytes, err := proto.Marshal(&contractProto)
	if err != nil {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}
	beego.Error(contractProtoBytes)

	beego.Warn("合约Id: " + contract.Id)
	beego.Warn("合约: " + contractModelStr)
	c.responseJsonBody(contract.Id, true, "查询合约成功!")
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
		c.responseJsonBodyCode(HTTP_STATUS_CODE_OK, "", false, err.Error())
		return
	}

	var contractTasks []model.ContractTask
	json.Unmarshal([]byte(contractTasksStr), &contractTasks)

	if err != nil {
		fmt.Println("Unmarshal ContractTasks fail!")
	}

	//TODO 1. track contract 合约跟踪
	//response := make(map[string]interface{})

	beego.Warn(c.Ctx.Request.RequestURI, "缺少合约跟踪方法![Track]")
	c.responseJsonBody(contractTasksStr, true, "合约跟踪查询成功!")

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

	beego.Debug("Token is " + token)

	contractModel := fromContractToContractModel(*contract)
	contractValid := contractModel.Validate()
	if !contractValid {
		c.responseJsonBodyCode(HTTP_STATUS_CODE_BadRequest, "", false, "contract error")
		return
	}
	//beego.Warn(contractModel)
	//beego.Warn(contractModel.Id)
	//beego.Warn(contractModel.GenerateId())

	ok := rethinkdb.InsertContract(common.Serialize(contractModel))
	response := make(map[string]interface{})

	beego.Info(c.Ctx.Request.RequestURI, "API Insert![Create Id:"+contractModel.Id+"]")
	response["id"] = contractModel.Id

	beego.Warn(c.Ctx.Request.RequestURI, "缺少合约更新方法![Update]")
	c.responseJsonBody(contract.Id, ok, "API Insert![Create Id:"+contractModel.Id+"]")

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
	beego.Warn(c.Ctx.Request.RequestURI, "缺少测试合约方法![Test]")

	response := make(map[string]interface{})

	response["id"] = time.Now().Unix()
	c.responseJsonBody(contract.Id, true, "合约测试成功!")
}
