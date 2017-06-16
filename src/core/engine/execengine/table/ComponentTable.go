package table

import (
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/inf"
)

type ComponentTable struct {
	//TODO: need sort struct
	//type: map[string][]property.PropertyT
	//        type:  Unknown, Data, Task, Expression
	CompTable map[string][]map[string]interface{} `json:"CompTable"`
}

func NewComponentTable() *ComponentTable {
	ct := &ComponentTable{}
	return ct
}

func (ct *ComponentTable) getComponentType(p_component interface{}) (string, string) {
	var r_type string = ""
	var r_name string = ""
	if p_component == nil {
		r_type = constdef.ComponentType[constdef.Component_Unknown]
		r_name = p_component.(inf.IComponent).GetName()
		return r_type, r_name
	}
	switch p_component.(type) {
	case *inf.IData:
		r_type = constdef.ComponentType[constdef.Component_Data]
		r_name = p_component.(inf.IData).GetName()
	case inf.IData:
		r_type = constdef.ComponentType[constdef.Component_Data]
		r_name = p_component.(inf.IData).GetName()
	case *inf.ITask:
		r_type = constdef.ComponentType[constdef.Component_Task]
		r_name = p_component.(inf.ITask).GetName()
	case inf.ITask:
		r_type = constdef.ComponentType[constdef.Component_Task]
		r_name = p_component.(inf.ITask).GetName()
	case *inf.IExpression:
		r_type = constdef.ComponentType[constdef.Component_Expression]
		r_name = p_component.(inf.IExpression).GetName()
	case inf.IExpression:
		r_type = constdef.ComponentType[constdef.Component_Expression]
		r_name = p_component.(inf.IExpression).GetName()
	default:
		r_type = constdef.ComponentType[constdef.Component_Unknown]
		r_name = p_component.(inf.IComponent).GetName()
	}
	return r_type, r_name
}

func (ct *ComponentTable) AddComponent(p_component interface{}) {
	if ct.CompTable == nil {
		ct.CompTable = make(map[string][]map[string]interface{})
	}
	c_type, c_name := ct.getComponentType(p_component)
	if _, ok := ct.CompTable[c_type]; !ok {
		ct.CompTable[c_type] = make([]map[string]interface{}, 0)
	}
	v_component := make(map[string]interface{})
	v_component[c_name] = p_component
	ct.CompTable[c_type] = append(ct.CompTable[c_type], v_component)
}

func (ct *ComponentTable) GetComponent(cname string, ctype string) interface{} {
	var g_component interface{}
	if ctype == "" {
		//总：map[string][]map[string]interface()
		for _, ct_value := range ct.CompTable {
			for _, ct_com := range ct_value {
				if _, ok := ct_com[cname]; ok {
					g_component = ct_com[cname]
				}
			}
		}
	} else {
		//总：map[string][]map[string]interface()
		//for: map[string]
		if v_comp_type, ok := ct.CompTable[ctype]; ok {
			//for: []
			for _, v_comp := range v_comp_type {
				//map[string]interface()
				if r_res, ok := v_comp[cname]; ok {
					g_component = r_res
				}
			}
		}
	}
	return g_component
}

func (ct *ComponentTable) GetTaskByID(cid string, ctype string) interface{} {
	var g_component interface{}
	if ctype == constdef.ComponentType[constdef.Component_Task] {
		//总：map[string][]map[string]interface()
		//for: map[string]
		if v_comp_type, ok := ct.CompTable[ctype]; ok {
			//for: []
			for _, v_comp := range v_comp_type {
				//map[string]interface()
				for _, v_value := range v_comp {
					if v_value == nil {
						continue
					}
					if v_value.(inf.ITask).GetTaskId() == cid {
						g_component = v_value
						break
					}
				}
			}
		}
	}
	return g_component
}

func (ct *ComponentTable) GetComponentByType(c_type string) []map[string]interface{} {
	if c_type == "" {
		return nil
	}
	if _, ok := ct.CompTable[c_type]; !ok {
		return nil
	}
	return ct.CompTable[c_type]
}

//更新PropertyTable中的元素
//Args: c_type string 组件类型
//      c_name string 组件名称
//      c_component interface{} 组件
func (ct *ComponentTable) UpdateComponent(c_type string, c_name string, c_component interface{}) error {
	var err error = nil
	//获取ctype对应的组件数组
	var new_task_component_array []map[string]interface{} = make([]map[string]interface{}, 0)
	task_component_array := ct.CompTable[c_type]
	for _, v_component_map := range task_component_array {
		//替换组件数组中对应的组件
		for v_key, v_value := range v_component_map {
			if v_key == c_name {
				v_component_map[v_key] = c_component
			} else {
				v_component_map[v_key] = v_value
			}
		}
		new_task_component_array = append(new_task_component_array, v_component_map)
	}
	//更新ctype对应的组件数组
	ct.CompTable[c_type] = new_task_component_array
	return err
}
