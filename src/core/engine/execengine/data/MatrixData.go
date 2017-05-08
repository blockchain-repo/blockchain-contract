package data

import (
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/inf"
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
	if nd.PropertyTable[_Parent] != nil  {
		parent_property := nd.PropertyTable[_Parent].(property.PropertyT)
		if parent_property.GetValue() != nil {
			v_general_data := parent_property.GetValue().(inf.IData)
			if v_general_data.GetName() != "" {
				return v_general_data.GetName() + "." + nd.GetCname()
			} else {
				return nd.GetCname()
			}
		}
	}
	return nd.GetCname()
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
func (gc MatrixData)GetCtype()string{
	if gc.PropertyTable["_Ctype"] == nil {
		return ""
	}
	ctype_property := gc.PropertyTable["_Ctype"].(property.PropertyT)
	return ctype_property.GetValue().(string)
}
//====================描述态==========================


//====================运行态==========================
func (md *MatrixData) InitMatrixData()error{
    var err error = nil
	err = md.InitGeneralData()
	if err != nil {
		//TODO log
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

