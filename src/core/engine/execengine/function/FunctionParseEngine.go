package function

import (
	"bytes"
	"errors"
	"github.com/astaxie/beego/logs"
	"unicontract/src/core/engine/common"
)

type FunctionParseEngine struct {
	ContractFunctions map[string]func(arg ...interface{}) (common.OperateResult, error)
}

//---------------------------------------------------------------------------
//类构造方法
func NewFunctionParseEngine() *FunctionParseEngine {
	bif := &FunctionParseEngine{}
	bif.ContractFunctions = make(map[string]func(arg ...interface{}) (common.OperateResult, error), 0)
	return bif
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//根据配置文件名称加载函数、方法集
//=====Common Method
func (bif *FunctionParseEngine) LoadFunctionsCommon() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionsCommon fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		logs.Error(r_buf.String())
		v_err = errors.New("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	bif.ContractFunctions["FuncTestMethod"] = FuncTestMethod
	bif.ContractFunctions["FuncCreateAsset"] = FuncCreateAsset
	bif.ContractFunctions["FuncTransferAsset"] = FuncTransferAsset
	bif.ContractFunctions["FuncTransferAssetComplete"] = FuncTransferAssetComplete
	bif.ContractFunctions["FuncUnfreezeAsset"] = FuncUnfreezeAsset
	bif.ContractFunctions["FuncInterim"] = FuncInterim
	bif.ContractFunctions["FuncInterimComplete"] = FuncInterimComplete
	bif.ContractFunctions["FuncIsConPutInUnichian"] = FuncIsConPutInUnichian
	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====Demo Method(产品演示Demo)
func (bif *FunctionParseEngine) LoadFunctionDEMO() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionDEMO fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		logs.Error(r_buf.String())
		v_err = errors.New("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here

	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====TIANJS Method(天安金交所)
func (bif *FunctionParseEngine) LoadFunctionTIANJS() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionTIANJS fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		logs.Error(r_buf.String())
		v_err = errors.New("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	bif.ContractFunctions["FuncTIANJSExample"] = FuncTIANJSExample

	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====GUANGXIBIANMAO Method(广西边贸)
func (bif *FunctionParseEngine) LoadFunctionGUANGXIBIAMAO() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionGUANGXIBIAMAO fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		logs.Error(r_buf.String())
		v_err = errors.New("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	bif.ContractFunctions["FuncBIANMAOExample"] = FuncBIANMAOExample

	return v_err
}
