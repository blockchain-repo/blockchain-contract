package constdef

const (
	Data_Unknown = iota
	Data_Numeric_Uint
	Data_Numeric_Int
	Data_Numeric_Float
	Data_Text
	Data_Date
	Data_Array
	Data_Compound
	Data_Matrix
	Data_DecisionCandidate
	Data_OperateResultData
)

var DataType = map[int]string{
	Data_Unknown : "Data_Unknown",
	Data_Numeric_Uint: "Data_Numeric_Uint",
	Data_Numeric_Int: "Data_Numeric_Int",
	Data_Numeric_Float: "Data_Numeric_Float",
	Data_Text: "Data_Text",
	Data_Date: "Data_Date",
	Data_Array: "Data_Array",
	Data_Compound: "Data_Compound",
	Data_Matrix: "Data_Matrix",
	Data_DecisionCandidate: "Data_DecisionCandidate",
	Data_OperateResultData: "Data_OperateResultData",
}