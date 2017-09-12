package execengine

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine"
	"unicontract/src/core/engine/common"
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
//Args: p_str_json      => 完整的contract Output结构体JSON
//      str_contractId  => contract的id，目前只用于日志记录
//说明：
//   反序列化回来的都是map类型， property_table中都是实际的struct
func (ce *ContractExecuter) Load(p_str_json, str_contractId string) error {
	fmt.Println("=================================================================================")
	uniledgerlog.Debug("contract json is :")
	uniledgerlog.Debug(p_str_json)
	fmt.Println("=================================================================================")

	var err error
	var r_buf bytes.Buffer
	r_buf.WriteString("Contract Executeor:Load.")
	if p_str_json == "" {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Param Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Param Error!")
	}
	//p_str_json为完整的合约交易(Output结构体)
	//0 识别Contract结构体和Relation结构体:
	var map_output_first interface{} = common.Deserialize(p_str_json)
	if map_output_first == nil {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Contract Map Deserialize Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Contract Map Deserialize Error!")
	}
	map_output_second, ok := map_output_first.(map[string]interface{})
	if !ok {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Assert Error!")
	}
	if map_output_second == nil || len(map_output_second) == 0 {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:[ContractOutput]Null Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("[ContractOutput]Null Error!")
	}
	if map_output_second["transaction"] == nil {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:[Transaction]Null Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("[Transaction]Null Error!")
	}
	map_transaction, ok := map_output_second["transaction"].(map[string]interface{})
	if !ok {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Transaction Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Transaction Assert Error!")
	}
	if map_transaction["Contract"] == nil {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:[Contract]Null Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("[Contract]Null Error!")
	}
	map_contract, ok := map_transaction["Contract"].(map[string]interface{})
	if !ok {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Contract Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Contract Assert Error!")
	}
	if map_transaction["Relation"] == nil {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:[Relation]Null Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("[Relation]Null Error!")
	}
	map_relation, ok := map_transaction["Relation"].(map[string]interface{})
	if !ok {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Relation Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Relation Assert Error!")
	}
	var str_json_contract string = common.Serialize(map_contract)
	if str_json_contract == "" {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Serialize Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Serialize Error!")
	}
	//l 反序列化
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, str_contractId, "deserialize to contract struct"))
	ret_contract, err := ce.contract_executer.Deserialize(str_json_contract)
	if err != nil {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Contract Struct Deserialize Error ," + err.Error())
		uniledgerlog.Error(r_buf.String())
		return err
	}
	ce.contract_executer, ok = ret_contract.(*contract.CognitiveContract)
	if !ok {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Contract Struct Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Contract Struct Assert Error!")
	}
	//2 Init初始化, 填充contract property_table
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, str_contractId, "struct init"))
	err = ce.contract_executer.InitCognitiveContract()
	if err != nil {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Error in Load(InitCognitiveContract) ," + err.Error())
		uniledgerlog.Error(r_buf.String())
		return err
	}
	r_buf.WriteString("[CName]:" + ce.contract_executer.GetCname() + "; ")
	r_buf.WriteString("[ContractId]:" + ce.contract_executer.GetContractId() + "; ")

	//3 Components填充 component_table 和 property_table
	//component_table: contract_component, task_component, data_component, expression_component
	//property_table: contract_property, task_property, data_property, expression_property
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, str_contractId, "contract struct load task"))
	for p_idx, p_component := range ce.contract_executer.GetContractComponents() {
		uniledgerlog.Debug("component[", p_idx, "]: ", p_component)
		err = loadTask(ce.contract_executer, p_component)
		if err != nil {
			r_buf.WriteString("[Result]:Load Fail;")
			r_buf.WriteString("[Error]:Error in Load(loadTask) ," + err.Error())
			uniledgerlog.Error(r_buf.String())
			return err
		}
	}
	ce.contract_executer.AddComponent(ce.contract_executer)
	//4 Check合约是否可以执行，并更新合约状态
	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, str_contractId, "check if can be executed and update contract state"))
	if !ce.contract_executer.CanExecute() {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Other Error - Contract can not execute!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Other Error - Contract can not execute!")
	}
	if !ce.contract_executer.UpdateContractState("") {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Other Error - Update ContractState Error!")
		uniledgerlog.Error(r_buf.String())
		return fmt.Errorf("Other Error - Update ContractState Error!")
	}
	err = ce.contract_executer.SetOrgTaskInfo(map_relation)
	if err != nil {
		r_buf.WriteString("[Result]:Load Fail;")
		r_buf.WriteString("[Error]:Error in Load(SetOrgTaskInfo) ," + err.Error())
		uniledgerlog.Error(r_buf.String())
		return err
	}

	uniledgerlog.Notice(fmt.Sprintf("[%s][The contract(%s) %s]",
		uniledgerlog.NO_ERROR, str_contractId, "load success"))
	r_buf.WriteString("[Result]:Load Success!")
	uniledgerlog.Info(r_buf.String())
	return err
}

//合约运行生命周期：合约预处理
func (ce *ContractExecuter) Prepare() {
	var r_buf bytes.Buffer
	//读取产品库函数配置
	ExecuteEngineConf, ok := engine.UCVMConf["ExecuteEngine"].(map[interface{}]interface{})
	if !ok {
		r_buf.WriteString("[Result]:Prepare Fail;")
		r_buf.WriteString("[Error]:Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return
	}
	product_func_str, ok := ExecuteEngineConf["function_source"].(string)
	if !ok {
		r_buf.WriteString("[Result]:Prepare Fail;")
		r_buf.WriteString("[Error]:Assert Error!")
		uniledgerlog.Error(r_buf.String())
		return
	}
	//TODO 加载特定产品函数库
	for _, v_product := range strings.Split(product_func_str, ",") {
		switch v_product {
		case constdef.FunctionSource[constdef.FUNCTION_SRC_DEMO]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionDEMO()
		case constdef.FunctionSource[constdef.FUNCTION_SRC_AUTOSELLER]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionAUTOSELLER()
		case constdef.FunctionSource[constdef.FUNCTION_SRC_ENERGYTRADING]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionENERGYTRADING()
		case constdef.FunctionSource[constdef.FUNCTION_SRC_GUANGXIBIANMAO]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionGUANGXIBIANMAO()
		case constdef.FunctionSource[constdef.FUNCTION_SRC_HOUSETRANSFER]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionHOUSETRANSFER()
		case constdef.FunctionSource[constdef.FUNCTION_SRC_RENTPAYMENT]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionRENTPAYMENT()
		case constdef.FunctionSource[constdef.FUNCTION_SRC_TIANJS]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionTIANJS()
		case constdef.FunctionSource[constdef.FUNCTION_SRC_TRANSFER]:
			ce.contract_executer.FunctionParseEngine.LoadFunctionTRANSFER()
		}
	}
}

//合约运行生命周期：合约启动
func (ce *ContractExecuter) Start() (int8, error) {
	var r_ret int8 = -1
	var r_err error
	var r_buf bytes.Buffer
	if ce.contract_executer == nil {
		r_buf.WriteString("[Result]:Start Fail;")
		r_buf.WriteString("[Error]:[p_contract]Null Error!")
		uniledgerlog.Error("[p_contract]Null Error!!")
		return r_ret, fmt.Errorf("[p_contract]Null Error!!")
	}
	r_buf.WriteString("Contract Executeor:Run.")
	r_buf.WriteString("[CName]:" + ce.contract_executer.GetCname() + "; ")
	r_buf.WriteString("[ContractId]:" + ce.contract_executer.GetContractId() + "; ")
	r_buf.WriteString("[Result]:")
	//注意： 此处为整个合约的执行结果，只标记合约退出的状态
	//    1. ret=0 ： 合约执行过程中，某任务没有达到执行条件，暂时退出，等待下轮扫描再次加载执行
	//    2. ret=-1： 合约执行过程中，某任务执行失败，       暂时退出，等待下轮扫描再次加载执行
	//    3. ret=1 ： 合约执行完成
	r_ret, r_err = ce.contract_executer.UpdateTasksState()
	if r_ret == 0 {
		r_buf.WriteString("合约任务未达到执行条件,等待再次扫描执行;")
		if r_err != nil {
			r_buf.WriteString("[Result]:Start Fail;")
			r_buf.WriteString("[Error]:Error in Start(UpdateTasksState) ," + r_err.Error())
		}
	} else if r_ret == -1 {
		r_buf.WriteString("合约任务执行失败,等待再次扫描执行;")
		if r_err != nil {
			r_buf.WriteString("[Result]:Start Fail;")
			r_buf.WriteString("[Error]:Error in Start(UpdateTasksState) ," + r_err.Error())
		}
	} else {
		r_buf.WriteString("合约任务执行完成;")
	}
	uniledgerlog.Info(r_buf.String())
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
	var r_str_json string
	var err error
	r_buf := bytes.Buffer{}
	r_buf.WriteString("Contract Executeor:ExportToJson.")
	if ce.contract_executer == nil {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("[Error]:Param[p_contrat] is null")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Param[p_contrat] is null!")
		return r_str_json, err
	}
	r_buf.WriteString("[CName]:" + ce.contract_executer.GetCname() + "; ")
	r_buf.WriteString("[ContractId]:" + ce.contract_executer.GetContractId() + "; ")
	//l 序列化
	r_str_json, err = ce.contract_executer.Serialize()
	if err != nil {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("[Error]:" + err.Error())
		uniledgerlog.Warn(r_buf.String())
		return r_str_json, err
	}
	r_buf.WriteString("[Result]:Export Success;")
	uniledgerlog.Info(r_buf.String())
	return r_str_json, err
}

//将合约当前运行状态导出,形成文本描述态: Text文档
//Return: string -> text file path
//        error  -> op error
func (ce *ContractExecuter) ExportToText() (string, error) {
	var err error
	r_bytes := bytes.Buffer{}
	r_buf := bytes.Buffer{}
	r_buf.WriteString("Contract Executeor:ExportToText.")
	if ce.contract_executer == nil {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Param[p_contrat] is null!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Param[p_contrat] is null!")
		return r_bytes.String(), err
	}
	r_buf.WriteString("[CName]:" + ce.contract_executer.GetCname() + "; ")
	r_buf.WriteString("[ContractId]:" + ce.contract_executer.GetContractId() + "; ")
	//1 解析json串，映射成text文本
	contractId, ok := ce.contract_executer.PropertyTable["_ContractId"].(property.PropertyT)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	c_id, ok := contractId.GetValue().(string)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	r_bytes.WriteString(contractId.GetName() + ":" + c_id + "\n")

	contractOwner, ok := ce.contract_executer.PropertyTable["_ContractOwners"].(property.PropertyT)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	_, ok = contractOwner.GetValue().([]string)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	contractCreateTime, ok := ce.contract_executer.PropertyTable["_CreateTime"].(property.PropertyT)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	str, ok := contractCreateTime.GetValue().(string)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	r_bytes.WriteString(contractCreateTime.GetName() + ":" + str + "\n")
	contractStartTime, ok := ce.contract_executer.PropertyTable["_StartTime"].(property.PropertyT)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return "", fmt.Errorf("assert error")
	}
	str, ok = contractStartTime.GetValue().(string)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	r_bytes.WriteString(contractStartTime.GetName() + ":" + str + "\n")
	contractEndTime, ok := ce.contract_executer.PropertyTable["_EndTime"].(property.PropertyT)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	str, ok = contractEndTime.GetValue().(string)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	r_bytes.WriteString(contractEndTime.GetName() + ":" + str + "\n")
	contractAssets, ok := ce.contract_executer.PropertyTable["_ContractAssets"].(property.PropertyT)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	arr_contractAsset, ok := contractAssets.GetValue().([]contract.ContractAsset)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	r_bytes.WriteString(contractAssets.GetName() + ":" + "\n")
	for _, p_asset := range arr_contractAsset {
		r_bytes.WriteString("    _AssetCaption" + ":" + p_asset.GetCaption() + "\n")
		r_bytes.WriteString("    _AssetDescription" + ":" + p_asset.GetDescription() + "\n")
		r_bytes.WriteString("    _AssetUnit" + ":" + p_asset.GetUnit() + "\n")
		f, ok := p_asset.GetAmount().(float64)
		if !ok {
			r_buf.WriteString("[Result]:Export Fail;")
			r_buf.WriteString("Assert Error!")
			uniledgerlog.Warn(r_buf.String())
			err = fmt.Errorf("Assert Error!")
			return r_bytes.String(), err
		}
		r_bytes.WriteString("    _AssetAmount" + ":" + strconv.FormatFloat(f, 'f', 10, 64) + "\n")
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
					ttask, ok := v_value.(inf.ITask)
					if !ok {
						r_buf.WriteString("[Result]:Export Fail;")
						r_buf.WriteString("Assert Error!")
						uniledgerlog.Warn(r_buf.String())
						err = fmt.Errorf("Assert Error!")
						return r_bytes.String(), err
					}
					r_bytes.WriteString(v_key + ":" + ttask.GetDescription() + "\n")
				}
			}
		}
	}

	contractSignature, ok := ce.contract_executer.PropertyTable["_ContractSignatures"].(property.PropertyT)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	arr_signatures, ok := contractSignature.GetValue().([]contract.ContractSignature)
	if !ok {
		r_buf.WriteString("[Result]:Export Fail;")
		r_buf.WriteString("Assert Error!")
		uniledgerlog.Warn(r_buf.String())
		err = fmt.Errorf("Assert Error!")
		return r_bytes.String(), err
	}
	r_bytes.WriteString(contractSignature.GetName() + ":" + "\n")
	for _, p_signature := range arr_signatures {
		r_bytes.WriteString("    _OwnerPubkey" + ":" + p_signature.GetOwnerPubkey() + "\n")
		r_bytes.WriteString("    _Signature" + ":" + p_signature.GetSignature() + "\n")
		r_bytes.WriteString("    _SignTimestamp" + ":" + p_signature.GetSignTimestamp() + "\n")
	}
	r_buf.WriteString("[Result]:Export Success;")
	uniledgerlog.Info(r_buf.String())
	return r_bytes.String(), err
}

//合约销毁，合约退出时执行
func (ce *ContractExecuter) Destory() {
	//合约全局对象销毁
	if ce.contract_executer != nil {
		ce.contract_executer = nil
	}
}
