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

	// filter shouldn`t use the api log!
	beego.InsertFilter("/*", beego.BeforeRouter, filters.MonitorFilter, false)
	// auth request app_id and app_key, return token
	beego.InsertFilter("/*", beego.BeforeRouter, filters.APIContentTypeFilter, true)

	// if true, add the api filter
	api_auth := beego.AppConfig.DefaultBool("api_auth", false)
	// if true, add the api rate limit filter
	api_rate_limit := beego.AppConfig.DefaultBool("api_rate_limit", false)
	if api_auth {
		beego.InsertFilter("/v1/auth/getAccessKey", beego.BeforeRouter, filters.APIAuthorizationFilter, true)
		beego.InsertFilter("/v1/auth/getToken", beego.BeforeRouter, filters.APIGetTokenFilter, true)
		beego.InsertFilter("/*", beego.BeforeRouter, filters.APIAuthFilter, true)
		if api_rate_limit {
			beego.InsertFilter("/*", beego.BeforeRouter, filters.APIRateLimitFilter, true)
		}
	}

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/contract",
			beego.NSRouter("/create", &controllers.ContractController{}, "post:Create"),
			beego.NSRouter("/queryPublishContract", &controllers.ContractController{}, "post:QueryPublishContract"),
			beego.NSRouter("/queryContractContent", &controllers.ContractController{}, "post:QueryContractContent"),
			beego.NSRouter("/query", &controllers.ContractController{}, "post:Query"),
			beego.NSRouter("/queryAll", &controllers.ContractController{}, "post:QueryAll"),
			beego.NSRouter("/queryLog", &controllers.ContractController{}, "post:QueryLog"),
			beego.NSRouter("/pressTest", &controllers.ContractController{}, "post:PressTest"),
			//demo使用---------------------------------------------------------------------------------------------------
			beego.NSRouter("/queryOutput", &controllers.ContractController{}, "post:QueryOutput"),
			beego.NSRouter("/queryOutputNum", &controllers.ContractController{}, "post:QueryOutputNum"),
			beego.NSRouter("/queryOutputDuration", &controllers.ContractController{}, "post:QueryOutputDuration"),
			beego.NSRouter("/queryAccountBalance", &controllers.ContractController{}, "post:QueryAccountBalance"),
			beego.NSRouter("/queryAmmeterBalance", &controllers.ContractController{}, "post:QueryAmmeterBalance"),
			beego.NSRouter("/queryRecords", &controllers.ContractController{}, "post:QueryRecords"),
			//demo使用---------------------------------------------------------------------------------------------------
		),
	)
	beego.AddNamespace(ns)
}
