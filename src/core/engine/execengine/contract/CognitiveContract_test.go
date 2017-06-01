package contract

import (
	"testing"
)

func TestContract(t *testing.T) {
	v_contract := CognitiveContract{}
	//Test Version
	if v_contract.GetUCVMVersion() != "v1.0" {
		t.Error("GetVersion Erro!")
	}

	//Test CopyRight
	if v_contract.GetUCVMCopyRight() != "uni-ledger" {
		t.Error("GetCopyRight Erro!")
	}

}

func TestRun(t *testing.T) {
	/*
		step 1: 创建合约
		//metaAttribute: version, copyright
		metaAttribute
		//合约属性： Cname,Ctype,Caption,Description,metaAttribute
		//           合约创建时间、合约状态、合约创建人
		//           合约签名
		propertyTable map[string] interface{}
		//任务组件列表：enquiry,action,decision,decisioncandicate,plan,data,
		//              合约主体, 合约资产
		componentTable table.ComponentTable
		//事件
		eventQueue event.EventQueue
		//事件 handle
		eventHandlerPool event.EventHandlerPool

		step 2: 合约转化（可视化==>json描述）
		==> contract json
		  		meta
		  		property
				component: task  data   object  assert

		step 3: 合约执行存储

		step 4: 合约执行json加载， 合约解析(json描述==>code描述)
		 	通过执行引擎，初始化成contract对象

		step 5: 合约执行任务识别， 执行状态判断

		step 6: 合约任务执行

		step 7: 合约产出共识
	*/
}
