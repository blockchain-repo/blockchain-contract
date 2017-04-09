package controllers

import (
	"encoding/json"
	"unicontract/src/api/models"

	"fmt"
	"github.com/astaxie/beego"

)

// Operations about Contract
type ContractController struct {
	beego.Controller
}


func (c *ContractController) Auth(signature string) bool {
	fmt.Println("hello Auth...")
	if signature == "" {
		return false
	}
	return true
}

// @Title CreateContract
// @Description create contract
// @Param	body		body 	models.Contract	true		"body for contract content"
// @Success 200 {int} models.Contract.Head.Id
// @Failure 403 body is empty
// @router /create [post]
func (c *ContractController) Create() {
	var contract models.Contract
	fmt.Println("input is ", contract)
	json.Unmarshal(c.Ctx.Input.RequestBody, &contract)
	types := c.Ctx.Input.Header("ReqData-Type")
	fmt.Println("request type is", types)
	fmt.Println("input is ", contract)
	cid := "00001test"
	c.Data["json"] = map[string]string{"cid": cid}
	c.ServeJSON()
}

// @Title Signature
// @Description signature the contract
// @Param	body		body 	interface{}	true		"body for contract id"
// @Success 200 {object} models.Contract
// @Failure 403 body is empty
// @router /signature [post]
func (c *ContractController) Signature() {
	var input map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)

	target := input["target"]
	desc := input["desc"]
	fmt.Println("input is ", input)

	fmt.Printf("target is %s, desc is %s\n", target, desc)

	json.Unmarshal(c.Ctx.Input.RequestBody, &input)
	json.Marshal(&input)
	fmt.Println("input is ", input)
	c.Data["json"] = input
	c.ServeJSON()
}

// @Title Stop
// @Description stop the contract
// @Param	body		body 	interface{}	true		"body for contract id"
// @Success 200 {string} stop success!
// @Failure 403 body is empty
// @router /stop [post]
func (c *ContractController) Stop() {
	var input map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)

	fmt.Println("input is ", input)
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)
	fmt.Println("input is ", input)
	c.Data["json"] = input
	c.ServeJSON()
}

// @Title Find
// @Description get contract by cid
// @Param	body		body 	interface{}	true			"The key for contract"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /find [post]
func (c *ContractController) Find() {
	var input map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)
	cid := input["cid"]

	fmt.Println("input is ", input["cid"])

	if cid != nil && cid != "" {
		//user, err := models.GetUser(uid)
		//if err != nil {
		//	c.Data["json"] = err.Error()
		//} else {
		//	c.Data["json"] = user
		//}
		fmt.Println("input is ", cid)
	}
	c.ServeJSON()
}

// @Title Track
// @Description track contract by uid
// @Param	cid		path 	string	true		"The key for contract"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /track [post]
func (c *ContractController) Track() {
	var input map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &input)

	cid := input["cid"]
	if cid != "" {
		//user, err := models.GetUser(uid)
		//if err != nil {
		//	c.Data["json"] = err.Error()
		//} else {
		//	c.Data["json"] = user
		//}
		fmt.Println("input is ", cid)
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the contract
// @Param	body		body 	models.Contract	true		"body for contract content"
// @Success 200 {object} models.Contract
// @Failure 403 cid is empty
// @router /update [post]
func (c *ContractController) Update() {
	var contract models.Contract
	fmt.Println("input is ", contract)

	json.Unmarshal(c.Ctx.Input.RequestBody, &contract)
	fmt.Println("contract.Id is ", contract.Id)
	cid := "00001test"
	c.Data["json"] = map[string]string{"cid": cid}
	c.ServeJSON()
}

// @Title Test
// @Description test the contract
// @Param	cid		path 	string	true		"The uid you want to test"
// @Success 200 {string} test success!
// @Failure 403 cid is empty
// @router /test [post]
func (c *ContractController) Test() {
	var contract models.Contract
	fmt.Println("contract is ", contract)
	json.Unmarshal(c.Ctx.Input.RequestBody, &contract)
	fmt.Println("Unmarshal contract is ", contract)

	cid := contract.Id

	fmt.Println("input is ", cid)
	c.Data["json"] = "delete success!"
	c.ServeJSON()
}
