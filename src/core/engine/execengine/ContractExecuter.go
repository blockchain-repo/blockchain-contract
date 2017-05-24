package execengine

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"
	"strings"
	"unicontract/src/core/engine"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/contract"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

type ContractExecuter struct {
	contract_executer *contract.CognitiveContract
}

func NewContractExecuter() *ContractExecuter {
	ce := &ContractExecuter{}
	ce.contract_executer = &contract.CognitiveContract{}
	return ce
}

//合约执行要点：以task为最小粒度执行、共识、入链
//合约执行过程：
//     0.合约共识后将合约初始化到扫描监控表中（contractID0 描述态）
//     1.扫描监控表中(flag=0)(contractID0 描述态)加入队列，并且flag 0=>1
//     2.出队合约contractID0，并查询区块链获取合约contractID0
//           存在，则Load合约contractID0
//           不存在，更新扫描监控表 flag1 => 0, 重复执行2
//     3.启动合约contractID0 通过UpdateTasksState()从合约启动start开始判断执行
//     4.运行任务task1（contractID0 运行态）
//            dromant   任务态： 达不到执行条件（0）
//                               执行正常 inprocess态； 达不到完成执行条件，等待周期后，inprocess退出（-1）
//                                                      写入产出表（contractID1）, 写入正常complete态；  (contractID1 共识，入链)， 入链成功 digcard态 退出（1）
//                                                                                                                                  达不到入链条件 complete态（-1）
//                                                                                 写入异常退出（-1）
//                               执行异常 dromant退出(-1)
//     5.任务后处理
//           0 dromant态
//           1 digcard态
//          -1 dromant 执行异常
//             inprocess 完成条件未达到 或 写入产出失败
//             complete 入链失败

//====运行态周期方法
//====运行生命周期：Load(描述态到运行态) =》 Prepare =》 Run  =》 Stop
//将描述态加载到内存中，形成运行态（即初始化Contract、ComponentTable、PropertyTable）
func (ce *ContractExecuter) Load(p_str_json string) error {
	var err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Executeor:Load.")
	//l 反序列化
	ret_contract, err := ce.contract_executer.Deserialize(p_str_json)
	ce.contract_executer = ret_contract.(*contract.CognitiveContract)
	if err != nil {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Error in Load(Deserialize)," + err.Error())
		logs.Error(r_buf.String())
		return err
	}
	//2 Init初始化, 填充contract property_table
	err = ce.contract_executer.InitCognitiveContract()
	if err != nil {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Error in Load(InitCognitiveContract)," + err.Error())
		logs.Error(r_buf.String())
		return err
	}
	r_buf.WriteString("[CName]:" + ce.contract_executer.GetCname() + "; ")
	r_buf.WriteString("[ContractId]:" + ce.contract_executer.GetContractId() + "; ")

	//3 Components填充 component_table 和 property_table
	for p_idx, p_component := range ce.contract_executer.GetContractComponents() {
		fmt.Println("component[", p_idx, "]: ", p_component)
		err = loadTask(ce.contract_executer, p_component)
		if err != nil {
			r_buf.WriteString("[Result]:Load Fail;")
			r_buf.WriteString("[Error]:Error in Load(loadTask)," + err.Error())
			logs.Error(r_buf.String())
			return err
		}
	}
	r_buf.WriteString("[Result]:Load Success!")
	logs.Info(r_buf.String())
	return err
}

//合约运行生命周期：合约预处理
func (ce *ContractExecuter) Prepare() {
	//读取产品库函数配置
	var ExecuteEngineConf map[interface{}]interface{} = engine.UCVMConf["ExecuteEngine"].(map[interface{}]interface{})
	var product_func_str string = ExecuteEngineConf["function_source"].(string)
	//TODO 加载特定产品函数库
	for _, v_product := range strings.Split(product_func_str, ",") {
		switch v_product {
		case constdef.FunctionSource[constdef.FUNCTION_SRC_DEMO]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionDEMO()
		case constdef.FunctionSource[constdef.FUNCTION_SRC_TIANJS]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionTIANJS()
		case constdef.FunctionSource[constdef.FUNCTION_SRC_GUANGXIBIANMAO]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionGUANGXIBIAMAO()
		}
	}
}

//合约运行生命周期：合约启动
func (ce *ContractExecuter) Start() (int8, error) {
	var r_ret int8 = -1
	var r_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if ce.contract_executer == nil {
		logs.Error("Contract Start fail, Param[p_contract] is null!")
		r_err = errors.New("Param[p_contract] is null!")
		return r_ret, r_err
	}
	r_buf.WriteString("Contract Executeor:Run.")
	r_buf.WriteString("[CName]:" + ce.contract_executer.GetCname() + "; ")
	r_buf.WriteString("[ContractId]:" + ce.contract_executer.GetContractId() + "; ")
	r_buf.WriteString("[Result]:")
	//注意： 此处为整个合约的执行结果，只标记合约退出的状态
	//    1. ret=0 ： 合约执行过程中，某任务没有达到执行条件，暂时退出，等待下轮扫描再次加载执行
	//    2. ret=-1： 合约执行过程中，某任务执行失败，        暂时退出，等待下轮扫描再次加载执行
	//    3. ret=1 ： 合约执行完成
	r_ret, r_err = ce.contract_executer.UpdateTasksState()
	if r_ret == 0 {
		r_buf.WriteString("任务未达到执行条件,等待再次扫描执行;")
		if r_err != nil {
			r_buf.WriteString("[Error]:" + r_err.Error())
		}
		logs.Warning(r_buf.String())
	} else if r_ret == -1 {
		r_buf.WriteString("任务执行失败,等待再次扫描执行;")
		if r_err != nil {
			r_buf.WriteString("[Error]:" + r_err.Error())
		}
		logs.Warning(r_buf.String())
	} else {
		r_buf.WriteString("合约执行完成;")
		logs.Info(r_buf.String())
	}
	return r_ret, r_err
}

//合约运行生命周期： 合约停止（终止）
func (ce *ContractExecuter) Stop() {

}

//合约测试：空跑，测试合约逻辑是否可以到达
func (ce *ContractExecuter) Test() {

}

//将合约当前运行状态导出,形成字符描述态: JSON串
//Return: string -> json str
//        error  -> op error
func (ce *ContractExecuter) ExportToJson() (string, error) {
	var r_str_json string = ""
	var err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Executeor:ExportToJson.")
	if ce.contract_executer == nil {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("[Error]:Param[p_contrat] is null")
		logs.Warning(r_buf.String())
		err = errors.New("Param[p_contrat] is null!")
		return r_str_json, err
	}
	r_buf.WriteString("[CName]:" + ce.contract_executer.GetCname() + "; ")
	r_buf.WriteString("[ContractId]:" + ce.contract_executer.GetContractId() + "; ")
	//l 序列化
	r_str_json, err = ce.contract_executer.Serialize()
	if err != nil {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("[Error]:" + err.Error())
		logs.Warning(r_buf.String())
		return r_str_json, err
	}
	r_buf.WriteString("[Result]:Export Success;")
	logs.Info(r_buf.String())
	return r_str_json, err
}

//将合约当前运行状态导出,形成文本描述态: Text文档
//Return: string -> text file path
//        error  -> op error
func (ce *ContractExecuter) ExportToText() (string, error) {
	var r_bytes bytes.Buffer = bytes.Buffer{}
	var err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Contract Executeor:ExportToText.")
	if ce.contract_executer == nil {
		r_buf.WriteString("[Result]:Export Fail;")
		logs.Warning(r_buf.String())
		err = errors.New("Param[p_contrat] is null!")
		return r_bytes.String(), err
	}
	r_buf.WriteString("[CName]:" + ce.contract_executer.GetCname() + "; ")
	r_buf.WriteString("[ContractId]:" + ce.contract_executer.GetContractId() + "; ")
	//1 解析json串，映射成text文本
	contractId := ce.contract_executer.PropertyTable["_ContractId"].(property.PropertyT)
	r_bytes.WriteString(contractId.GetName() + ":" + contractId.GetValue().(string) + "\n")

	contractOwner := ce.contract_executer.PropertyTable["_ContractOwners"].(property.PropertyT)
	r_bytes.WriteString(contractOwner.GetName() + ":" + contractOwner.GetValue().([]string)[0] + "  " + contractOwner.GetValue().([]string)[1] + "\n")

	contractCreateTime := ce.contract_executer.PropertyTable["_CreateTime"].(property.PropertyT)
	r_bytes.WriteString(contractCreateTime.GetName() + ":" + contractCreateTime.GetValue().(string) + "\n")
	contractStartTime := ce.contract_executer.PropertyTable["_StartTime"].(property.PropertyT)
	r_bytes.WriteString(contractStartTime.GetName() + ":" + contractStartTime.GetValue().(string) + "\n")
	contractEndTime := ce.contract_executer.PropertyTable["_EndTime"].(property.PropertyT)
	r_bytes.WriteString(contractEndTime.GetName() + ":" + contractEndTime.GetValue().(string) + "\n")

	contractAssets := ce.contract_executer.PropertyTable["_ContractAssets"].(property.PropertyT)
	arr_contractAsset := contractAssets.GetValue().([]contract.ContractAsset)
	r_bytes.WriteString(contractAssets.GetName() + ":" + "\n")
	for _, p_asset := range arr_contractAsset {
		r_bytes.WriteString("    _AssetCaption" + ":" + p_asset.GetCaption() + "\n")
		r_bytes.WriteString("    _AssetDescription" + ":" + p_asset.GetDescription() + "\n")
		r_bytes.WriteString("    _AssetUnit" + ":" + p_asset.GetUnit() + "\n")
		r_bytes.WriteString("    _AssetAmount" + ":" + strconv.FormatFloat(p_asset.GetAmount().(float64), 'f', 10, 64) + "\n")
		for m_key, m_value := range p_asset.GetMetaData() {
			r_bytes.WriteString("  " + m_key + ":" + m_value + "\n")
		}

	}
	//map[string][]map[string]interface{}
	contractComponents := ce.contract_executer.ComponentTable.CompTable
	for v_component_type, v_component_arr := range contractComponents {
		for _, a_component_map := range v_component_arr {
			for v_key, v_value := range a_component_map {
				if v_component_type == constdef.ComponentType[constdef.Component_Task] {
					r_bytes.WriteString(v_key + ":" + v_value.(inf.ITask).GetDescription() + "\n")
				}
			}
		}
	}

	contractSignature := ce.contract_executer.PropertyTable["_ContractSignatures"].(property.PropertyT)
	arr_signatures := contractSignature.GetValue().([]contract.ContractSignature)
	r_bytes.WriteString(contractSignature.GetName() + ":" + "\n")
	for _, p_signature := range arr_signatures {
		r_bytes.WriteString("    _OwnerPubkey" + ":" + p_signature.GetOwnerPubkey() + "\n")
		r_bytes.WriteString("    _Signature" + ":" + p_signature.GetSignature() + "\n")
		r_bytes.WriteString("    _SignTimestamp" + ":" + p_signature.GetSignTimestamp() + "\n")
	}
	r_buf.WriteString("[Result]:Export Success;")
	logs.Info(r_buf.String())
	return r_bytes.String(), err
}

//合约销毁，合约退出时执行
func (ce *ContractExecuter) Destory() {
	//合约全局对象销毁
	if ce.contract_executer != nil {
		ce.contract_executer = nil
	}
}
