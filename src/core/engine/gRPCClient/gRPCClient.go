package gRPCClient

import (
	"encoding/json"
	"fmt"
)

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

import (
	log "unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	pb "unicontract/src/core/engine/gRPCClient/gRPCProto"
	"unicontract/src/core/engine/gRPCClient/paramProto"
)

//---------------------------------------------------------------------------
const (
	address = "localhost:50051"
)

//---------------------------------------------------------------------------
func FunctionRun(requestID, funcName, funcParams string) (common.OperateResult, error) {
	var result common.OperateResult
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
