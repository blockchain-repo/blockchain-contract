package task

import (
	"testing"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/expression"
)

func TestInitEnquery(t *testing.T){
	e := *new(Enquiry)
	e.InitEnquriy()
	e.SetCname("test_enquiry")
	e.SetCaption("enquiry")
	e.SetDescription("enquiry is task")
	if e.GetCname() != "test_enquiry" {
		t.Error("Init Enquiry(name) Error!")
	}
	if e.GetCtype() != constdef.ComponentType[constdef.Component_Task] +"."+constdef.TaskType[constdef.Task_Enquiry] {
		t.Error("Init Enquiry(ctype) Error!")
	}
	if e.GetCaption() != "enquiry" {
		t.Error("Init Enquiry(caption) Error!")
	}
	if e.GetDescription() != "enquiry is task" {
		t.Error("Init Enquiry(description) Error!")
	}
}

func TestAddAndGetData(t *testing.T){
	e := new(Enquiry)
	e.InitEnquriy()
	e.SetCname("test_enquiry")
	e.SetCaption("enquiry")
	e.SetDescription("enquiry is task")
    //use method: interface implements use pointer
	//test AddData
	var data_exp_1 string ="data expression 1"
	var data_num_1 data.IntData = data.IntData{}
	data_num_1.InitIntData()
	data_num_1.SetCname("DATA_NUM_1")
	data_num_1.SetCaption("DATA_NUM_1")
	data_num_1.SetDescription("data type is int 1")
	e.AddData(&data_num_1, data_exp_1)

	var data_exp_2 string ="data expression 2"
	var data_num_2 data.IntData = *new(data.IntData)
	data_num_2.InitIntData()
	data_num_2.SetCname("DATA_NUM_2")
	data_num_2.SetCaption("DATA_NUM_2")
	data_num_2.SetDescription("data type is int 2")
	e.AddData(&data_num_2, data_exp_2)

	//test GetDataList
	if len(e.GetDataList()) != 2 {
		t.Error("Add 2 IntData, but length not 2!")
	}
	//test GetDataSetterValueExpression
	if len(e.GetDataValueSetterExpressionList()) != 2{
		t.Error("Add 2 IntData, but setterValueExpression not 2!")
	}
	// test GetData
	t_data,err := e.GetData("DATA_NUM_1")
	if err != nil || t_data == nil{
		t.Error("GetData test Error!")
		return
	}
	tt_data := t_data.(*data.IntData)
    if tt_data.GetName() != "DATA_NUM_1" {
		t.Error("GetData test Error,check name Error!")
	}
	if tt_data.GetCaption() != "DATA_NUM_1" {
		t.Error("GetCaption test Error,check caption Error!")
	}
	if tt_data.GetDescription() != "data type is int 1" {
		t.Error("GetDescription test Error,check descrition Error!")
	}
	//test GetDataExpression
	t_data_expression,err := e.GetDataExpression("DATA_NUM_1")
	if err != nil || t_data_expression == nil {
		t.Error("GetDataExpression test Error")
		return
	}
	test_expression := t_data_expression.(*expression.GeneralExpression)
	test_expression.InitExpression()
	if test_expression.GetExpressionStr() != "data expression 1" {
		t.Error("GetDataValueSeeterExpression test Error!")
	}
}

func TestAddAndGetData2(t *testing.T){
	e := new(Enquiry)
	e.InitEnquriy()
	e.SetCname("test_enquiry")
	e.SetCaption("enquiry")
	e.SetDescription("enquiry is task")
	//use method: interface implements use object
	//test AddData
	var data_exp_1 string ="data expression 1"
	var data_num_1 data.IntData = data.IntData{}
	data_num_1.InitIntData()
	data_num_1.SetCname("DATA_NUM_1")
	data_num_1.SetCaption("DATA_NUM_1")
	data_num_1.SetDescription("data type is int 1")
	e.AddData(data_num_1, data_exp_1)

	var data_exp_2 string ="data expression 2"
	var data_num_2 data.IntData = data.IntData{}
	data_num_2.InitIntData()
	data_num_2.SetCname("DATA_NUM_2")
	data_num_2.SetCaption("DATA_NUM_2")
	data_num_2.SetDescription("data type is int 2")
	e.AddData(data_num_2, data_exp_2)

	//test GetDataList
	if len(e.GetDataList()) != 2 {
		t.Error("Add 2 IntData, but data length not 2!")
	}
	if len(e.GetDataValueSetterExpressionList()) != 2 {
		t.Error("Add 2 IntData, but data expression length not 2!")
	}
	//test GetData
	t_data,err := e.GetData("DATA_NUM_1")
	if err != nil || t_data == nil{
		t.Error("GetData test Error!")
		return
	}
	//test GetDataExpression
	t_data_expression,err := e.GetDataExpression("DATA_NUM_1")
	if err != nil || t_data_expression == nil {
		t.Error("GetDataExpression test Error")
		return
	}
	//test RemoveData
	e.RemoveData("DATA_NUM_1")
	if len(e.GetDataList()) != 1 {
		t.Error("Remove Data Error,data length Error!")
	}
	if len(e.GetDataValueSetterExpressionList()) != 1 {
		t.Error("Remove Data Error,expression length Error!")
	}
	e.RemoveData("DATA_NUM_2")
	if len(e.GetDataList()) != 0 {
		t.Error("Remove Data Error,data length Error!")
	}
	if len(e.GetDataValueSetterExpressionList()) != 0 {
		t.Error("Remove Data Error,expression length Error!")
	}
}

func TestAddAndGetData3(t *testing.T){
	e := new(Enquiry)
	e.InitEnquriy()
	e.SetCname("test_enquiry")
	e.SetCaption("enquiry")
	e.SetDescription("enquiry is task")
	//use method: interface implements use pointer
	//test AddData
	var data_exp_1 string ="data expression 1"
	var data_num_1 *data.IntData = &data.IntData{}
	data_num_1.InitIntData()
	data_num_1.SetCname("DATA_NUM_1")
	data_num_1.SetCaption("DATA_NUM_1")
	data_num_1.SetDescription("data type is int 1")
	e.AddData(data_num_1, data_exp_1)
	var data_exp_2 string ="data expression 2"
	var data_num_2 *data.IntData = &data.IntData{}
	data_num_2.InitIntData()
	data_num_2.SetCname("DATA_NUM_2")
	data_num_2.SetCaption("DATA_NUM_2")
	data_num_2.SetDescription("data type is int 2")
	e.AddData(data_num_2, data_exp_2)

	//test GetDataList
	if len(e.GetDataList()) != 2 {
		t.Error("Add 2 IntData, but length not 2!")
	}
	//test GetData
	t_data,err := e.GetData("DATA_NUM_1")
	if err != nil || t_data == nil{
		t.Error("GetData test Error!")
		return
	}
}