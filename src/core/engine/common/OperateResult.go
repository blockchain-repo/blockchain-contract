package common

import (
	"bytes"
	"strconv"
)

const (
	SUCCESS = 200
	FAIL = 400
	ERROR = 500
)

type OperateResult struct {
	Code int  `json:"Code"`
	Message string  `json:"Message"`
	Data string `json:"Data"`
}

func (or *OperateResult) GetMessage()string{
	return or.Message
}
func (or *OperateResult) GetCode()int{
	return or.Code
}
func (or *OperateResult) GetData()string{
	return or.Data
}
func (or *OperateResult) SetMessage(p_message string){
	or.Message = p_message
}
func (or *OperateResult) SetCode(p_code int){
	or.Code = p_code
}
func (or *OperateResult) SetData(p_data string){
	or.Data = p_data
}
func (or *OperateResult)ToString()string{
	var r_buf bytes.Buffer = bytes.Buffer{}
	r_buf.WriteString("Result")
	r_buf.WriteString(": [Code]: " + strconv.Itoa(or.Code))
	r_buf.WriteString(", [Message]: " + or.Message)
	r_buf.WriteString(", [Data]: " + or.Data)
	return r_buf.String()
}