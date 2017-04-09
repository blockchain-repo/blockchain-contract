package filters


import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

func responseWithStatusCode(ctx *context.Context, status int, output string) {
	ctx.Output.SetStatus(status)
	ctx.ResponseWriter.Write([]byte(output))
}

func ContractFilter(ctx *context.Context) {
	contentType := ctx.Input.Header("Content-Type")
	requestDataType := ctx.Input.Header("ReqData-Type")
	if contentType == "" || requestDataType == ""{
		responseWithStatusCode(ctx, 404,"error Headers")
		beego.Debug("ContractFilter contentType or requestDataType is empty!")
	}

	//fmt.Println("types is ",types)
	//fmt.Println("Content_Type is ",Content_Type)
	//if types == "proto"{
	//	fmt.Println("need use 数据")
	//}else if types == "json"{
	//	//ctx.Abort(401,"222")
	//	ctx.Output.SetStatus(404)
	//	ctx.ResponseWriter.Write([]byte("gogogo back"))
	//	fmt.Println("need use 数据")
	//}else{
	//	return
	//}
}