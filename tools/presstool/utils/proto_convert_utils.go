package utils

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicontract/src/core/protos"
)

func readContractJsonFromFile(file string) ([]byte, error) {
	_, err := os.Stat(file)
	if err != nil {
		fmt.Println("file not exist", file)
		// todo byte
		return nil, err
	}

	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
		// todo byte
		return nil, err
	}

	jsonByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonByte, nil

}

func convertJsonToContractProto(bt []byte) (protos.Contract, error) {
	var contractProto protos.Contract
	err := json.Unmarshal(bt, &contractProto)
	if err != nil {
		fmt.Println(err)
		return contractProto, err
	}
	return contractProto, nil
}

func ReadJsonToContractProto(file string) (protos.Contract, error) {
	var contractProto protos.Contract
	jsonByte, err := readContractJsonFromFile(file)
	if err != nil {
		fmt.Println(err)
		return contractProto, err
	}
	contractProto, err = convertJsonToContractProto(jsonByte)
	if err != nil {
		fmt.Println(err)
		return contractProto, err
	}
	return contractProto, nil
}

func ReadJsonToContractProtoSerialize(file string) ([]byte, error) {
	var contractProto protos.Contract
	jsonByte, err := readContractJsonFromFile(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	contractProto, err = convertJsonToContractProto(jsonByte)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	contractProtoSerialize, err := proto.Marshal(&contractProto)
	if err != nil {
		fmt.Println("contractProtoSerialize ", err.Error())
	}
	return contractProtoSerialize, nil
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
