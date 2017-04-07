package routers

import (
	"github.com/astaxie/beego"
)

func init() {

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
			Method: "Stop",
			Router: `/stop`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ContractController"],
		beego.ControllerComments{
			Method: "Find",
			Router: `/find`,
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

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"] = append(beego.GlobalControllerRouter["unicontract/src/api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
