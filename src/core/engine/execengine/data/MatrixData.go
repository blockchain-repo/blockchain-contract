package data

import (
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/inf"
	"github.com/astaxie/beego/logs"
)

//TODO
type MatrixData struct {
	GeneralData
}

func NewMatrixData() *MatrixData{
	nd := &MatrixData{}
	return nd
}
//====================接口方法========================
func (nd MatrixData)GetName()string{
	return nd.GeneralData.GetName()
}

func (nd MatrixData) GetValue() interface{}{
	value_property := nd.PropertyTable[_Value].(property.PropertyT)
	return value_property.GetValue()
}

func (nd MatrixData)GetContract() inf.ICognitiveContract {
	return nd.GeneralComponent.GetContract()
}
func (nd MatrixData)SetContract(p_contract inf.ICognitiveContract) {
	nd.GeneralComponent.SetContract(p_contract)
}
func (nd MatrixData)GetCtype()string{
	return nd.GeneralData.GetCtype()
}
func (nd MatrixData) SetValue(p_Value interface{}){
	nd.GeneralData.SetValue(p_Value)
}
//====================描述态==========================


//====================运行态==========================
func (md *MatrixData) InitMatrixData()error{
    var err error = nil
	err = md.InitGeneralData()
	if err != nil {
		logs.Error("InitMatrixData fail["+err.Error()+"]")
		return err
	}
	md.SetCtype(constdef.ComponentType[constdef.Component_Data] + "." + constdef.DataType[constdef.Data_Matrix])
	return err
}
//=====运算
//TODO
func (md *MatrixData) Size()(int, error) {
	var r_err error = nil
	var r_size int = 0

	return r_size, r_err
}

