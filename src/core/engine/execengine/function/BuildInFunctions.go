package function

import (
	"unicontract/src/core/engine/common"
	"errors"
	"bytes"
	"regexp"
	"strings"
	"fmt"
	"reflect"
)

var ContractFunctions map[string]func(arg...interface{})(common.OperateResult,error) = make(map[string]func(arg...interface{})(common.OperateResult,error) , 0)

//根据配置文件名称加载函数、方法集
//=====Common Method
func LoadFunctionsCommon()error{
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionsCommon fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		//TODO log
		v_err = errors.New("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	ContractFunctions["TestMethod"] = TestMethod

	return v_err
}

//=====TIANJS Method(天安金交所)
func LoadFunctionTIANJS()error{
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionsCommon fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		//TODO log
		v_err = errors.New("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	ContractFunctions["TestExample"] = TestExample

	return v_err
}


//函数运行
func RunFunction(p_function string) (common.OperateResult,error){
	var v_err error = nil
	var v_result common.OperateResult = common.OperateResult{}
	var r_buf bytes.Buffer = bytes.Buffer{}
	if p_function == ""{
		r_buf.WriteString("[Result]: RunFunction fail;")
		r_buf.WriteString("[Error]: param[p_function] is nil!")
		//TODO log
		v_err = errors.New(" param[p_function] is nil!")
		v_result = common.OperateResult{Code:400, Message:r_buf.String()}
		return v_result,v_err
	}
	//校验参数格式 xxxx(xxx, xxx, xxx)
	if v_bool,v_err := regexp.MatchString(`[a-zA-Z_0-9]+\(.*\)`, p_function);!v_bool || v_err != nil {
		r_buf.WriteString("[Result]: RunFunction fail;")
		r_buf.WriteString("[Error]: param[p_function] format error!")
		//TODO log
		v_err = errors.New(" param[p_function] format error!")
		v_result = common.OperateResult{Code:400, Message:r_buf.String()}
		return v_result,v_err
	}
   //正则匹配函数名
	name_reg := regexp.MustCompile(`\s*([^\(]+)`)
	func_name := strings.TrimSpace(name_reg.FindString(p_function))
	//TODO log
	fmt.Println("func_name:", func_name)
	//正则匹配函数的参数变量
	param_reg := regexp.MustCompile(`\((.*)\)`)
	func_param_str := strings.Trim(param_reg.FindString(p_function), "(|)")
	//TODO log
	fmt.Println("func_params:", func_param_str)
	//分割匹配的函数参数列表
	func_param_array := strings.Split(func_param_str, ",")
    //函数调用
	func_run := reflect.ValueOf(ContractFunctions[func_name])
	func_params := make([]reflect.Value, 0)
	for _,v_args := range func_param_array {
		//识别字符串，获取参数的值
		//TODO: Need add expression read
		v_arg_value := v_args
		func_params = append(func_params, reflect.ValueOf(v_arg_value))
	}
	func_result_arr := func_run.Call(func_params)
	return func_result_arr[0].Interface().(common.OperateResult),func_result_arr[1].Interface().(error)
}
