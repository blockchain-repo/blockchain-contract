package gRPCClient

import (
	"encoding/json"
	"fmt"
	"strconv"
)

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

import (
	log "unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine"
	"unicontract/src/core/engine/common"
	pb "unicontract/src/core/engine/gRPCClient/gRPCProto"
	"unicontract/src/core/engine/gRPCClient/paramProto"
)

//---------------------------------------------------------------------------
var (
	On     bool
	server string
	port   string
)

//---------------------------------------------------------------------------
func Init() {
	GRPCConf := engine.UCVMConf["GRPC"].(map[interface{}]interface{})
	server = GRPCConf["GRPCServer"].(string)
	tmp := GRPCConf["GRPCPort"].(int)
	port = strconv.Itoa(tmp)
	On = GRPCConf["GRPCon"].(bool)
}

//---------------------------------------------------------------------------
func QueryFuncType(funcName string) (int32, error) {
	address := server + ":" + port
	log.Debug(fmt.Sprintf("[%s][%s]", log.DEBUG_NO_ERROR, "GRPC server is "+address))
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error(fmt.Sprintf("[%s][%s]", log.CONNECTION_ERROR, err.Error()))
		return 0, err
	}
	defer conn.Close()
	c := pb.NewFunctionClient(conn)

	log.Info(fmt.Sprintf("[%s][Execute QueryFuncType for (%s)]", log.NO_ERROR, funcName))
	ret, err := c.QueryFuncType(context.Background(),
		&pb.QueryRequest{FunctionName: funcName})
	if err != nil {
		log.Error(fmt.Sprintf("[%s][%s]", log.OTHER_ERROR, err.Error()))
		return 0, err
	}
	return ret.Result, err
}

//---------------------------------------------------------------------------
func FunctionRun(requestID, funcName, funcParams string) (common.OperateResult, error) {
	var result common.OperateResult
	address := server + ":" + port
	log.Debug(fmt.Sprintf("[%s][%s]", log.DEBUG_NO_ERROR, "GRPC server is "+address))
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error(fmt.Sprintf("[%s][%s]", log.CONNECTION_ERROR, err.Error()))
		return result, err
	}
	defer conn.Close()
	c := pb.NewFunctionClient(conn)

	var requestParam paramProto.ReqParamStruct
	requestParam.RequestId = requestID
	requestParam.FuncName = funcName
	requestParam.FuncParams = funcParams
	slData, err := proto.Marshal(&requestParam)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][%s]", log.SERIALIZE_ERROR, err.Error()))
		return result, err
	}

	log.Info(fmt.Sprintf("[%s][requestID is %s, funcName is %s, funcParams is %s]",
		log.NO_ERROR, requestID, funcName, funcParams))
	r1, err := c.ExecuteFunc(context.Background(), &pb.ExecRequest{Params: string(slData)})
	if err != nil {
		log.Error(fmt.Sprintf("[%s][%s]", log.OTHER_ERROR, err.Error()))
		return result, err
	}

	var response paramProto.ResParamStruct
	err = proto.Unmarshal([]byte(r1.Result), &response)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][%s]", log.DESERIALIZE_ERROR, err.Error()))
		return result, err
	}

	err = json.Unmarshal([]byte(response.ResResult), &result)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][%s]", log.DESERIALIZE_ERROR, err.Error()))
	}
	return result, err
}

//---------------------------------------------------------------------------
