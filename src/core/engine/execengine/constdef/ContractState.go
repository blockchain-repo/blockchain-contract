package constdef

const (
	Contract_Unknown = iota
	Contract_Create
	Contract_Signature
	Contract_In_Process
	Contract_Completed
	Contract_Discarded
)

var ContractState = map[int]string{
	Contract_Unknown:    "Contract_Unknown",
	Contract_Create:     "Contract_Create",
	Contract_Signature:  "Contract_Signature",
	Contract_In_Process: "Contract_In_Process",
	Contract_Completed:  "Contract_Completed",
	Contract_Discarded:  "Contract_Discarded",
}
