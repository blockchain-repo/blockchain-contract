package expressionutils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/component"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/function"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/task"
)

//Test Function:
//    ReflectComponent
//    ReflectVariable
//    RunFunction

//---------------------------------------------------------------------------
func CreateComponent() inf.IComponent {
	var v_map_1 map[string]string = make(map[string]string, 0)
	v_map_1["key_1"] = "value_1"
	v_map_1["key_2"] = "value_2"
	var v_category_1 []string = []string{"a", "b", "c", "d"}
	var v_options_1 map[string]int = map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	var v_general_data_1 data.GeneralData = data.GeneralData{Value: "test_data_1", Options: v_options_1, Category: v_category_1}
	var v_category_2 []string = []string{"e", "f", "g", "h"}
	var v_options_2 map[string]int = map[string]int{"5": 5, "6": 6, "7": 7, "8": 8}
	var v_general_data_2 data.GeneralData = data.GeneralData{Value: "test_data_2", Options: v_options_2, Category: v_category_2}

	var v_datarange_1 [2]int = [2]int{1, 2}
	var v_datarange_2 [2]int = [2]int{3, 4}
	var v_data_arr_1 []interface{} = make([]interface{}, 0)
	var data_1_1 data.IntData = data.IntData{GeneralData: v_general_data_1, DataRangeInt: v_datarange_1}
	var data_1_2 data.IntData = data.IntData{GeneralData: v_general_data_2, DataRangeInt: v_datarange_2}
	v_data_arr_1 = append(v_data_arr_1, data_1_1)
	v_data_arr_1 = append(v_data_arr_1, data_1_2)

	var v_expressionResult_1 common.OperateResult = common.OperateResult{Code: 200, Message: "success"}
	var v_expressionResult_2 common.OperateResult = common.OperateResult{Code: 500, Message: "error"}
	var v_condition_arr_1 []interface{} = make([]interface{}, 0)
	var exprssion_1_1 expression.GeneralExpression = expression.GeneralExpression{ExpressionResult: v_expressionResult_1}
	var exprssion_1_2 expression.GeneralExpression = expression.GeneralExpression{ExpressionResult: v_expressionResult_2}
	v_condition_arr_1 = append(v_condition_arr_1, exprssion_1_1)
	v_condition_arr_1 = append(v_condition_arr_1, exprssion_1_2)

	var v_general_comp_1 component.GeneralComponent = component.GeneralComponent{Cname: "test_enquiry", MetaAttribute: v_map_1}
	var v_general_task_1 task.GeneralTask = task.GeneralTask{GeneralComponent: v_general_comp_1, DataList: v_data_arr_1, PreCondition: v_condition_arr_1}
	var v_component_1 inf.IComponent = task.Enquiry{GeneralTask: v_general_task_1}
	return v_component_1
}

func ParseValue(p_variable string) interface{} {
	if p_variable == "" {
		return nil
	}

	v_variable_arr := strings.Split(p_variable, ".")
	var v_component reflect.Value = reflect.ValueOf(CreateComponent())

	v_variable_count := len(v_variable_arr)
	v_idx := 1
	var v_component_object reflect.Value = reflect.Value{}
	v_component_object = v_component
	var v_property_field reflect.Value = reflect.Value{}
	for v_idx < v_variable_count {
		//fmt.Println("++++++++++++++++++++++++++++++++++++")
		//fmt.Println(v_component_object.Kind(), "    ", v_component_object.Type(), "   ", v_component_object.Interface())
		v_property_field = v_component_object.FieldByName(v_variable_arr[v_idx])
		//fmt.Println(v_property_field.Kind(), "    ", v_property_field.Type(), "   ", v_property_field.Interface())
		switch v_property_field.Kind() {
		case reflect.Map:
			v_idx = v_idx + 1
			if v_idx >= v_variable_count {
				break
			}
			v_component_object = v_property_field.MapIndex(reflect.ValueOf(v_variable_arr[v_idx]))
			v_property_field = reflect.ValueOf(v_component_object)
		case reflect.Slice:
			v_idx = v_idx + 1
			if v_idx >= v_variable_count {
				break
			}
			v_arr_idx, _ := strconv.Atoi(v_variable_arr[v_idx])
			//TODO: 随着字段类型的增加，需要补充支持
			switch v_property_field.Interface().(type) {
			case []string:
				data_arr := v_property_field.Interface().([]string)
				v_component_object = reflect.ValueOf(data_arr[v_arr_idx])
			case []interface{}:
				data_arr := v_property_field.Interface().([]interface{})
				v_component_object = reflect.ValueOf(data_arr[v_arr_idx])
			}
			v_property_field = reflect.ValueOf(v_component_object)
		case reflect.Array:
			v_idx = v_idx + 1
			if v_idx >= v_variable_count {
				break
			}
			v_arr_idx, _ := strconv.Atoi(v_variable_arr[v_idx])
			//TODO：随着字段类型的增加，需要补充支持
			switch v_property_field.Interface().(type) {
			case [2]int:
				data_arr := v_property_field.Interface().([2]int)
				v_component_object = reflect.ValueOf(data_arr[v_arr_idx])
			default:
				v_component_object = reflect.ValueOf(v_property_field.Interface())
			}
			v_property_field = reflect.ValueOf(v_component_object)
		case reflect.Struct:
			v_struct_property := v_property_field.Interface()
			v_component_object = reflect.ValueOf(v_struct_property)
		default:
			break
		}
		v_idx = v_idx + 1
	}
	return v_property_field.Interface()
}

//---------------------------------------------------------------------------
func TestOther(t *testing.T) {
	var v_str1 string = " abc "
	fmt.Println(strings.TrimSpace(v_str1))
	var v_str2 string = " a b c "
	fmt.Println(strings.TrimSpace(v_str2))
}

func TestReflectComponent(t *testing.T) {
	//TODO
}

func TestReflectVariable(t *testing.T) {
	//Test inf.IComponent => task.Enquiry
	var v_map_1 map[string]string = make(map[string]string, 0)
	v_map_1["key_1"] = "value_1"
	v_map_1["key_2"] = "value_2"

	var v_options_1 map[string]int = map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	var v_general_data_1 data.GeneralData = data.GeneralData{Value: "test_data_1", Options: v_options_1}
	var v_options_2 map[string]int = map[string]int{"5": 5, "6": 6, "7": 7, "8": 8}
	var v_general_data_2 data.GeneralData = data.GeneralData{Value: "test_data_2", Options: v_options_2}

	var v_datarange_1 [2]int = [2]int{1, 2}
	var v_datarange_2 [2]int = [2]int{3, 4}
	var v_data_arr_1 []interface{} = make([]interface{}, 0)
	var data_1_1 data.IntData = data.IntData{GeneralData: v_general_data_1, DataRangeInt: v_datarange_1}
	var data_1_2 data.IntData = data.IntData{GeneralData: v_general_data_2, DataRangeInt: v_datarange_2}
	v_data_arr_1 = append(v_data_arr_1, data_1_1)
	v_data_arr_1 = append(v_data_arr_1, data_1_2)

	var v_general_comp_1 component.GeneralComponent = component.GeneralComponent{Cname: "test_enquiry", MetaAttribute: v_map_1}
	var v_general_task_1 task.GeneralTask = task.GeneralTask{GeneralComponent: v_general_comp_1, DataList: v_data_arr_1}
	var v_component_1 inf.IComponent = task.Enquiry{GeneralTask: v_general_task_1}

	//2. property from property_table: string
	v_component_object := reflect.ValueOf(v_component_1)
	v_property_field := v_component_object.FieldByName("Cname")
	v_property_type := v_property_field.Kind()
	v_property_value := v_property_field.Interface()
	fmt.Println(v_property_type, v_property_value)
	//   property from property_table: map
	v_property_field = v_component_object.FieldByName("MetaAttribute")
	fmt.Println(v_property_field.Interface())
	if v_property_field.Kind() == reflect.Map {
		fmt.Println(v_property_field.MapIndex(reflect.ValueOf("key_1")))
	}
	//  property from property_table: struct array
	v_property_field = v_component_object.FieldByName("DataList")
	fmt.Println(v_property_field.Interface())
	if v_property_field.Kind() == reflect.Slice {
		fmt.Println(v_property_field.Index(0))
		fmt.Println()
		fmt.Println(v_property_field.Index(1))
		fmt.Println()

		data_arr := v_property_field.Interface().([]interface{})
		v_component_object = reflect.ValueOf(data_arr[0])
		v_property_field = v_component_object.FieldByName("Value")
		fmt.Println(v_property_field)

		v_property_field = v_component_object.FieldByName("DataRange")
		if v_property_field.Kind() == reflect.Array {
			fmt.Println(v_property_field.Index(0))
			fmt.Println(v_property_field.Index(1))
		}
		v_component_object = reflect.ValueOf(data_arr[1])
		v_property_field = v_component_object.FieldByName("Value")
		fmt.Println(v_property_field)

		v_property_field = v_component_object.FieldByName("DataRange")
		if v_property_field.Kind() == reflect.Array {
			fmt.Println(v_property_field.Index(0))
			fmt.Println(v_property_field.Index(1))
		}
	}
}

func TestReflectVariable_2(t *testing.T) {
	fmt.Println("===============================================")

	var str_variable_name_1 string = "test_enquiry.Cname"
	fmt.Println("test_enquiry.Cname:", "test_enquiry")
	fmt.Println(ParseValue(str_variable_name_1))

	var str_variable_name_2 string = "test_enquiry.MetaAttribute.key_1"
	fmt.Println("test_enquiry.MetaAttribute.key_1:", "value_1")
	fmt.Println(ParseValue(str_variable_name_2))

	var str_variable_name_3 string = "test_enquiry.DataList.0.Value"
	fmt.Println("test_enquiry.DataList.0.Value:", "test_data_1")
	fmt.Println(ParseValue(str_variable_name_3))

	var str_variable_name_4 string = "test_enquiry.DataList.1.Value"
	fmt.Println("test_enquiry.DataList.0.Value:", "test_data_2")
	fmt.Println(ParseValue(str_variable_name_4))

	var str_variable_name_5 string = "test_enquiry.DataList.0.DataRange"
	fmt.Println("test_enquiry.DataList.0.DataRange:", "[1,2]")
	fmt.Println(ParseValue(str_variable_name_5))

	var str_variable_name_6 string = "test_enquiry.DataList.0.DataRange.0"
	fmt.Println("test_enquiry.DataList.0.DataRange.0:", "1")
	fmt.Println(ParseValue(str_variable_name_6))

	var str_variable_name_7 string = "test_enquiry.DataList.1.DataRange"
	fmt.Println("test_enquiry.DataList.1.DataRange:", "[3,4]")
	fmt.Println(ParseValue(str_variable_name_7))

	var str_variable_name_8 string = "test_enquiry.DataList.1.DataRange.0"
	fmt.Println("test_enquiry.DataList.1.DataRange.0:", "3")
	fmt.Println(ParseValue(str_variable_name_8))

	var str_variable_name_9 string = "test_enquiry.DataList.0.Options"
	fmt.Println("test_enquiry.DataList.0.Options:", `{"1":1,"2":2,"3":3,"4":4}`)
	fmt.Println(ParseValue(str_variable_name_9))

	var str_variable_name_10 string = "test_enquiry.DataList.0.Options.3"
	fmt.Println("test_enquiry.DataList.0.Options.3:", "3")
	fmt.Println(ParseValue(str_variable_name_10))

	var str_variable_name_11 string = "test_enquiry.DataList.1.Options"
	fmt.Println("test_enquiry.DataList.1.Options:", `{"5":5,"6":6,"7":7,"8":8}`)
	fmt.Println(ParseValue(str_variable_name_11))

	var str_variable_name_12 string = "test_enquiry.DataList.1.Options.8"
	fmt.Println("test_enquiry.DataList.1.Options.8:", "8")
	fmt.Println(ParseValue(str_variable_name_12))

	var str_variable_name_13 string = "test_enquiry.DataList.0.Category.2"
	fmt.Println("test_enquiry.DataList.0.Category.2:", "c")
	fmt.Println(ParseValue(str_variable_name_13))

	var str_variable_name_14 string = "test_enquiry.DataList.1.Category.2"
	fmt.Println("test_enquiry.DataList.1.Category.2:", "g")
	fmt.Println(ParseValue(str_variable_name_14))

	var str_variable_name_15 string = "test_enquiry.DataList.0.Category"
	fmt.Println("test_enquiry.DataList.0.Category:", `{"a","b","c","d"}`)
	fmt.Println(ParseValue(str_variable_name_15))

	var str_variable_name_16 string = "test_enquiry.DataList.1.Category"
	fmt.Println("test_enquiry.DataList.1.Category:", `{"e","f","g","h"}`)
	fmt.Println(ParseValue(str_variable_name_16))

	var str_variable_name_17 string = "test_enquiry.PreCondition.0"
	fmt.Println("test_enquiry.PreCondition.0:", `.....`)
	fmt.Println(ParseValue(str_variable_name_17))

	var str_variable_name_18 string = "test_enquiry.PreCondition.1"
	fmt.Println("test_enquiry.PreCondition.1:", `.....`)
	fmt.Println(ParseValue(str_variable_name_18))

	var str_variable_name_19 string = "test_enquiry.PreCondition.0.ExpressionResult.Code"
	fmt.Println("test_enquiry.PreCondition.0.ExpressionResult.Code:", `200`)
	fmt.Println(ParseValue(str_variable_name_19))

	var str_variable_name_20 string = "test_enquiry.PreCondition.1.ExpressionResult.Code"
	fmt.Println("test_enquiry.PreCondition.1.ExpressionResult.Code:", `500`)
	fmt.Println(ParseValue(str_variable_name_20))

}

func TestRunFunction(t *testing.T) {
	v_express_parse := NewExpressionParseEngine()
	v_function_engine := function.NewFunctionParseEngine()
	v_function_engine.LoadFunctionsCommon()
	v_express_parse.SetFunctionEngine(v_function_engine)

	slTestRightStr := []string{
		`FuncGetNowDateTimestamp()`,
	}
	for _, value := range slTestRightStr {
		v_express_parse.RunFunction(value)
	}

	slTestErrorStr := []string{
		"Func",
		"func",
	}
	for index, value := range slTestErrorStr {
		if v_express_parse.IsExprFunction(value) {
			t.Errorf("index is [ %d ], value is [ %s ] is Function, Check Error!", index, value)
		}
	}
}
