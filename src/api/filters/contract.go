package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"unicontract/src/common"
	//"github.com/golang/protobuf/proto"
	//"unicontract/src/core/protos"
	//"fmt"
)

func responseWithStatusCode(ctx *context.Context, status int, output string) {
	//ctx.Output.SetStatus(status) // invalid setStatus
	ctx.ResponseWriter.WriteHeader(status)
	ctx.ResponseWriter.Write([]byte(output))
}

func ContractFilter(ctx *context.Context) {
	contentType := ctx.Input.Header("Content-Type")
	requestDataType := ctx.Input.Header("RequestDataType")

	if contentType == "" || requestDataType == "" {
		result :=make(map[string]interface{})
		result["msg"] = "error Headers"
		result["status"] = 404
		responseWithStatusCode(ctx, 404, common.Serialize(result))
		beego.Error("ContractFilter contentType or requestDataType is empty!")

	} else if contentType == "application/json" && requestDataType == "json" {
		beego.Debug("RequestDataType is json!")
	} else if contentType == "application/octet-stream" && requestDataType == "proto" {
		//input := ctx.Input.RequestBody
		//
		//contract := &protos.ContractProto{}
		//err2 := proto.Unmarshal(input, contract)
		////err2 := json.Unmarshal(input, input2)
		//if err2 != nil {
		//	beego.Error("marshaling error2: ", err2)
		//}
		//fmt.Println("oring contract(application/octet-stream): \n", contract)
		//
		//fmt.Println(contract.Id)

	}else if contentType == "application/x-protobuf" && requestDataType == "proto" {
		//beego.Debug("RequestDataType is proto!")
		//input := ctx.Input.RequestBody
		//contract := &protos.ContractProto{}
		//err2 := proto.Unmarshal(input, contract)
		//if err2 != nil {
		//	beego.Error("marshaling error2: ", err2)
		//}
		//fmt.Println("oring contract(application/x-protobuf): \n", contract)
		//fmt.Println(contract.Id)
	}
}
