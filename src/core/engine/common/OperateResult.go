package common

import (
	"bytes"
	"strconv"
)

const (
	SUCCESS = 200
	FAIL    = 400
	ERROR   = 500
)

type OperateResult struct {
	Code    int         `json:"Code"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
	Output  interface{} `json:"Output"`
}

func NewOperateResult() *OperateResult {
	or := &OperateResult{}
	return or
}

func (or *OperateResult) GetMessage() string {
	return or.Message
}
func (or *OperateResult) GetCode() int {
	return or.Code
}
func (or *OperateResult) GetData() interface{} {
	return or.Data
}
func (or *OperateResult) GetOutput() interface{} {
	return or.Output
}
func (or *OperateResult) SetMessage(p_message string) {
	or.Message = p_message
}
func (or *OperateResult) SetCode(p_code int) {
	or.Code = p_code
}
func (or *OperateResult) SetData(p_data interface{}) {
	or.Data = p_data
}
func (or *OperateResult) SetOutput(p_output interface{}) {
	or.Output = p_output
}
func (or *OperateResult) ToString() string {
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Result")
	r_buf.WriteString(": [Code]: " + strconv.Itoa(or.Code))
	r_buf.WriteString(", [Message]: " + or.Message)
	r_buf.WriteString(", [Data]: ")
	r_buf.WriteString(or.Data.(string))
	return r_buf.String()
}
