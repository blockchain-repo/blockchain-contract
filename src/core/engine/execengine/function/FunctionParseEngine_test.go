package function

import (
	"testing"
	"unicontract/src/core/engine/common"
	"reflect"
	"fmt"
	"regexp"
	"strings"
)

func TestRegex(t *testing.T){
	reg,_ := regexp.Compile(`[a-zA-Z_0-9]+\(.*\)`)
	v_function := "testMethod(p_arg_1, p_arg_2, p_arg_3)"
	v_bool,v_err := regexp.MatchString(`[a-zA-Z_0-9]+\(.*\)`, v_function)
	if !v_bool || v_err != nil {
		t.Error("param[" + v_function + "] function format error!")
	}
	fmt.Println(reg.FindString(v_function))

	v_function = " testMethod (p_arg)"
	v_bool,v_err = regexp.MatchString(`[a-zA-Z_0-9]+\(.*\)`, v_function)
	if v_bool {
		t.Error("param[" + v_function + "] function format error!")
	}
	fmt.Println(reg.FindString(v_function))

	v_function = "testMethod_1()"
	v_bool,v_err = regexp.MatchString(`[a-zA-Z_0-9]+\(.*\)`, v_function)
	if !v_bool || v_err != nil {
		t.Error("param[" + v_function + "] function format error!")
	}
	fmt.Println(reg.FindString(v_function))

	v_function = "testMethod_2"
	v_bool,v_err = regexp.MatchString(`[a-zA-Z_0-9]+\(.*\)`, v_function)
	if v_bool {
		t.Error("param[" + v_function + "] function format error!")
	}
	fmt.Println(reg.FindString(v_function))
}

//测试：可变参数方法
func TestArgsReflect(t *testing.T){
	var v_arg_0 int = 1
	var v_arg_1 string = "abc"
	var v_arg_2 map[string]string = map[string]string{"a":"aValue", "b":"bValue", "c":"cValue"}
	v_result,_ := FuncTestMethod(v_arg_0, v_arg_1, v_arg_2)
	fmt.Println("result: ", v_result)
	fmt.Println()
}

//测试：方法反射
func TestMethodReflect(t *testing.T){
	var func_set map[string]func(arg...interface{})(common.OperateResult,error) = make(map[string]func(arg...interface{})(common.OperateResult,error) , 0)
	func_set["TestMethod"] = FuncTestMethod

	//反射机制 1.1
	var func_name string = "TestMethod"
	v_result,_ := func_set[func_name](1, "abc", map[string]string{"a":"aValue", "b":"bValue", "c":"cValue"})
	fmt.Println("result: ", v_result)
	fmt.Println()
	v_result,_ = func_set[func_name]("test method reflect")
	fmt.Println("result: ", v_result)
	fmt.Println()

	//反射机制 1.2
	var func_info string = " TestMethod (1, \"abcdefg\", test(),`map[string]string{'a':'aValue', 'b':'bValue', 'c':'cValue'}`)"
	//正则匹配函数名
	name_reg := regexp.MustCompile(`\s*([^\(]+)`)
	func_name = strings.TrimSpace(name_reg.FindString(func_info))
	fmt.Println("func_name:", func_name)
	//正则匹配函数的参数变量
	param_reg := regexp.MustCompile(`\((.*)\)`)
	func_params := strings.Trim(param_reg.FindString(func_info), "(|)")
	fmt.Println("func_params:", func_params)
	//分割匹配的函数参数列表
	fmt.Println(strings.Split(func_params, ","))
	fmt.Println()

	//反射机制 2.1
	func_run := reflect.ValueOf(func_set["TestMethod"])
	params := make([]reflect.Value, 0)
	params = append(params, reflect.ValueOf(20))
	params = append(params, reflect.ValueOf("abc"))
	params = append(params, reflect.ValueOf(map[string]string{"a":"aValue", "b":"bValue", "c":"cValue"}))

	rs := func_run.Call(params)
	fmt.Println("result:",rs[0].Interface())
	fmt.Println("err:",rs[1].Interface())
}