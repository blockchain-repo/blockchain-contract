package controllers

import (
	"unicontract/src/api/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type ContractController struct {
	beego.Controller
}

// @Title CreateContract
// @Description create contract
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ContractController) Post() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	c.Data["json"] = map[string]string{"uid": uid}
	c.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (c *ContractController) GetAll() {
	users := models.GetAllUsers()
	c.Data["json"] = users
	c.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (c *ContractController) Get() {
	uid := c.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = user
		}
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (c *ContractController) Put() {
	uid := c.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(c.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = uu
		}
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (c *ContractController) Delete() {
	uid := c.GetString(":uid")
	models.DeleteUser(uid)
	c.Data["json"] = "delete success!"
	c.ServeJSON()
}

