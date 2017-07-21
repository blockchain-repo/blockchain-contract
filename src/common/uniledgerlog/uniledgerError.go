package uniledgerlog

// example:
// [API_ERROR][Token is error!]

type ErrorType int

const (
	API_ERROR        ErrorType = iota
	DB_ERROR                   // DB operate error
	CONNECTION_ERROR           // connect to server error
	ASSET_ERROR
	PARAM_ERROR
	ASSERT_ERROR
	SERIALIZE_ERROR
	DESERIALIZE_ERROR
	NULL_ERROR
	OTHER_ERROR
	DEBUG_NO_ERROR
)

func (this ErrorType) String() string {
	switch this {
	case API_ERROR:
		return "API_Error"
	case DB_ERROR:
		return "DB_Error"
	case CONNECTION_ERROR:
		return "Connect_Error"
	case ASSET_ERROR:
		return "Asset_Error"
	case PARAM_ERROR:
		return "Param_Error"
	case ASSERT_ERROR:
		return "Assert_Error"
	case SERIALIZE_ERROR:
		return "Serialize_Error"
	case DESERIALIZE_ERROR:
		return "Deserialize_Error"
	case NULL_ERROR:
		return "Null_Error"
	case OTHER_ERROR:
		return "Other_Error"
	case DEBUG_NO_ERROR:
		return "Debug_No_Error"
	default:
		return "Unknow_Error"
	}
}
