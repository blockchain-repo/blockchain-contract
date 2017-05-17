package function

import (
	"fmt"
	"strconv"
	"unicontract/src/core/engine/common"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//++++++++++合约机【天安金交中心】专用扩展方法++++++++++
//++++++++++++++++++++++++++++++++++++++++++++++++++++++

//样例方法
func FuncTIANJSExample(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	var v_map_args map[string]interface{} = nil
	if len(args) != 0 {
		v_map_args = make(map[string]interface{}, 0)
	}
	//识别可变参数
	for v_idx, v_args := range args {
		tmp_arg := "v_arg_" + strconv.Itoa(v_idx)
		v_map_args[tmp_arg] = v_args
	}
	//调用参数
	for v_name, v_value := range v_map_args {
		fmt.Println(v_name, ":", v_value)
	}
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}
