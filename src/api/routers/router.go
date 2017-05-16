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
	"unicontract/src/api/filters"
)

func init() {

	beego.InsertFilter("/*", beego.BeforeRouter, filters.ContractFilter, false)
	beego.InsertFilter("/*", beego.BeforeRouter, filters.MonitorFilter, false)
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/contract",
			beego.NSRouter("/authSignature", &controllers.ContractController{}, "post:AuthSignature"),
			beego.NSRouter("/create", &controllers.ContractController{}, "post:Create"),
			beego.NSRouter("/signature", &controllers.ContractController{}, "post:Signature"),
			beego.NSRouter("/terminate", &controllers.ContractController{}, "post:Terminate"),
			beego.NSRouter("/query", &controllers.ContractController{}, "post:Query"),
			beego.NSRouter("/queryAll", &controllers.ContractController{}, "post:QueryAll"),
			beego.NSRouter("/track", &controllers.ContractController{}, "post:Track"),
			beego.NSRouter("/update", &controllers.ContractController{}, "post:Update"),
			beego.NSRouter("/test", &controllers.ContractController{}, "post:Test"),
			beego.NSRouter("/pressTest", &controllers.ContractController{}, "post:PressTest"),
		),
	)
	beego.AddNamespace(ns)
}
