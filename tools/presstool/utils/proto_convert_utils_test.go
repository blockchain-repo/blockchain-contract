package utils

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"testing"
	"unicontract/src/common"
	"unicontract/src/core/protos"
)

func Test_ReadJsonFromFile(t *testing.T) {
	file, _ := os.Getwd()
	//fmt.Println("current path:", file)
	file = GetParentDirectory(file) + "/data/pb.json"

	contractProto, err := ReadJsonToContractProto(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("contractProto", contractProto)

	contractProtoSerialize, err := proto.Marshal(&contractProto)
	if err != nil {
		fmt.Println("contractProtoSerialize ", err.Error())
	}
	fmt.Println("contractProtoSerialize", contractProtoSerialize)

	var contractProto2 protos.Contract
	proto.Unmarshal(contractProtoSerialize, &contractProto2)

	fmt.Println("contractProtoDeSerialize\n", common.StructSerializePretty(contractProto2))

}

//todo use this generate the binary file for post press test
func Test_ReadJsonToContractProtoSerialize(t *testing.T) {
	file, _ := os.Getwd()
	//fmt.Println("current path:", file)
	dir := GetParentDirectory(file)
	file = dir + "/data/contract_pb2.json"
	byteFile := dir + "/data/contract_pb2_bytes"

	contractProtoSerialize, err := ReadJsonToContractProtoSerialize(file)
	if err != nil {
		fmt.Println(err)
		return
	}


	ioutil.WriteFile(byteFile, contractProtoSerialize, 0644)
	fmt.Println("contractProtoSerialize", contractProtoSerialize)

	var contractProto2 protos.Contract
	proto.Unmarshal(contractProtoSerialize, &contractProto2)

	fmt.Println("contractProtoDeSerialize\n", common.StructSerializePretty(contractProto2))

}
