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
	"github.com/astaxie/beego/plugins/cors"
	"unicontract/src/api/controllers"
	"unicontract/src/api/filters"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true}))

	beego.InsertFilter("/*", beego.BeforeRouter, filters.MonitorFilter, false)
	// auth request app_id and app_key, return token
	beego.InsertFilter("/v1/unicontract/auth", beego.BeforeRouter, filters.APIAuthorizationFilter, true)
	//beego.InsertFilter("/*", beego.BeforeRouter, filters.APIContentTypeFilter, false)
	beego.InsertFilter("/*", beego.BeforeRouter, filters.APIAuthFilter, true)
	ns := beego.NewNamespace("/v1/unicontract",
		beego.NSNamespace("/contract",
			beego.NSRouter("/authSignature", &controllers.ContractController{}, "get:AuthSignature"),
			beego.NSRouter("/create", &controllers.ContractController{}, "post:Create"),
			//beego.NSRouter("/signature", &controllers.ContractController{}, "post:Signature"),
			//beego.NSRouter("/terminate", &controllers.ContractController{}, "post:Terminate"),
			beego.NSRouter("/queryPublishContract", &controllers.ContractController{}, "get:QueryPublishContract"),
			beego.NSRouter("/queryContractContent", &controllers.ContractController{}, "get:QueryContractContent"),
			beego.NSRouter("/query", &controllers.ContractController{}, "get:Query"),
			beego.NSRouter("/queryAll", &controllers.ContractController{}, "get:QueryAll"),
			beego.NSRouter("/queryLog", &controllers.ContractController{}, "get:QueryLog"),
			//beego.NSRouter("/update", &controllers.ContractController{}, "post:Update"),
			//beego.NSRouter("/test", &controllers.ContractController{}, "post:Test"),
			beego.NSRouter("/pressTest", &controllers.ContractController{}, "post:PressTest"),
			//demo使用---------------------------------------------------------------------------------------------------
			beego.NSRouter("/queryOutput", &controllers.ContractController{}, "get:QueryOutput"),
			beego.NSRouter("/queryOutputNum", &controllers.ContractController{}, "get:QueryOutputNum"),
			beego.NSRouter("/queryOutputDuration", &controllers.ContractController{}, "get:QueryOutputDuration"),
			beego.NSRouter("/queryAccountBalance", &controllers.ContractController{}, "get:QueryAccountBalance"),
			beego.NSRouter("/queryAmmeterBalance", &controllers.ContractController{}, "get:QueryAmmeterBalance"),
			beego.NSRouter("/queryRecords", &controllers.ContractController{}, "get:QueryRecords"),
			//demo使用---------------------------------------------------------------------------------------------------
		),
	)
	beego.AddNamespace(ns)
}
