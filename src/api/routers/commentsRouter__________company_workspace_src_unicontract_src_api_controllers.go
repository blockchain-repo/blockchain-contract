package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"],
		beego.ControllerComments{
			Method: "AuthSignature",
			Router: `/authSignature`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"],
		beego.ControllerComments{
			Method: "Signature",
			Router: `/signature`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"],
		beego.ControllerComments{
			Method: "Terminate",
			Router: `/terminate`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"],
		beego.ControllerComments{
			Method: "Query",
			Router: `/query`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"],
		beego.ControllerComments{
			Method: "Track",
			Router: `/track`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
