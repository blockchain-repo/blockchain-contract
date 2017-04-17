// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"unicontract/src/api/controllers"
	"net/http"
	"fmt"
	//"unicontract-back/src/core/model"
	"github.com/astaxie/beego/context"
	"unicontract/src/api/filters"
	//"unicontract-back/src/core/model"
)

func msgHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HandlerFunc!")
	fmt.Println("heel")
	return
}

var UrlManager = func(ctx *context.Context) {
	// 数据库读取全部的 url mapping 数据
	fmt.Println("heel")
	fmt.Println("heel")
	fmt.Println("heel")
	types := ctx.Input.Header("ReqData-Type")
	Content_Type := ctx.Input.Header("Content-Type")
	fmt.Println("types is ",types)
	fmt.Println("Content_Type is ",Content_Type)
	if types == "proto"{
		fmt.Println("need use 数据")
	}else if types == "json"{
		//ctx.Abort(401,"222")
		ctx.Output.SetStatus(404)
		ctx.ResponseWriter.Write([]byte("gogogo back"))
		fmt.Println("need use 数据")
	}else{
		return
	}
}

func init() {

	//beego.Handler("/v1/contract/*", http.HandlerFunc(msgHandler))
	//beego.NSHandler("/v1/contract/", http.HandlerFunc(msgHandler))
	beego.InsertFilter("/*",beego.BeforeRouter, filters.ContractFilter, false)
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/contract",
			beego.NSInclude(
				&controllers.ContractController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
