package api

// contract relation
const (
	REQUEST_FIELD_AUTH_APPID         = "appId"
	REQUEST_FIELD_AUTH_TOKEN         = "token"
	REQUEST_FIELD_AUTH_TIMESTAMP     = "timestamp"
	REQUEST_FIELD_AUTH_SIGN          = "sign"
	REQUEST_FIELD_CONTRACT_ID        = "contractId"
	REQUEST_FIELD_CONTRACT_OWNER     = "contractOwner"
	REQUEST_FIELD_CONTRACT_STATE     = "contractState"
	REQUEST_FIELD_CONTRACT_NAME      = "contractName"
	REQUEST_FIELD_CONTRACT_STARTTIME = "startTime"
	REQUEST_FIELD_CONTRACT_ENDTIME   = "endTime"
	REQUEST_FIELD_PAGE               = "page"
	REQUEST_FIELD_PAGE_SIZE          = "pageSize"
)

var ALLOW_REQUEST_PARAMETERS_ALL = map[string]bool{
	/*************** FOR API AUTH ************/
	"appId":     true,
	"token":     true,
	"timestamp": true,
	"sign":      true,
	/*************** FOR API QUERY ***********/
	//"format":   true, //sort=field1,field2,
	"sort":     true,
	"page":     true,
	"pagesize": true,
	/*************** FOR API MODEL QUERY ***********/
	"contractId":    true,
	"contractOwner": true,
	"contractState": true,
	"contractName":  true,
	"startTime":     true,
	"endTime":       true,
}

// filter sort invalid fields
var ALLOW_REQUEST_PARAMETERS_MODEL = map[string]bool{
	/*************** FOR API MODEL QUERY ***********/
	"contractId":    true,
	"contractOwner": true,
	"contractState": true,
	"contractName":  true,
	"startTime":     true,
	"endTime":       true,
}

// FOR API FILTER
var ALLOW_REQUEST_PARAMETERS_BASIC = map[string]bool{
	"appId":     true,
	"token":     true,
	"timestamp": true,
	"sign":      true,
}

var REQUEST_CONTRACT_STATE_MAP = map[string]bool{
	"Contract_Unknown":    true,
	"Contract_Create":     true,
	"Contract_Signature":  true,
	"Contract_In_Process": true,
	"Contract_Completed":  true,
	"Contract_Discarded":  true,
}
