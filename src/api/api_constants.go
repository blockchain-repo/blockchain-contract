package api

var Constants_ContractState = map[string]string{
	"Contract_Unknown":    "未知",
	"Contract_Create":     "创建",
	"Contract_Signature":  "签约",
	"Contract_In_Process": "执行中",
	"Contract_Completed":  "完成",
	"Contract_Discarded":  "终止",
}

var Constants_TaskState = map[string]string{
	"TaskState_Dormant":     "休眠",
	"TaskState_In_Progress": "执行中",
	"TaskState_Completed":   "完成",
	"TaskState_Discard":     "终止",
}
