package component

import (
	"fmt"
	"reflect"
	"testing"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/property"
)

func TestGeneralComponent(t *testing.T) {
	g_component := &GeneralComponent{Cname: "GeneralComponent", Caption: "GeneralComponent", Description: "GeneralComponent Test"}
	g_component.InitGeneralComponent()

	if g_component.GetCname() != "GeneralComponent" || g_component.Cname != "GeneralComponent" {
		t.Error("InitGeneralComponent Error,GetName Error!")
	}

	if g_component.GetCtype() != constdef.ComponentType[constdef.Component_Unknown] || g_component.Ctype != constdef.ComponentType[constdef.Component_Unknown] {
		t.Error("InitGeneralComponent Error,GetCtype Error!")
	}
	if g_component.GetCaption() != "GeneralComponent" || g_component.Caption != "GeneralComponent" {
		t.Error("InitGeneralComponent Error,GetCaption Error!")
	}
	if g_component.GetDescription() != "GeneralComponent Test" || g_component.Description != "GeneralComponent Test" {
		t.Error("InitGeneralComponent Error,GetDescription Error!")
	}

	//Test SetName
	g_component.SetCname("TestComponent")
	if g_component.GetCname() != "TestComponent" {
		t.Error("InitGeneralComponent Error,GetName Error!")
	}
	//Test MetaAttribute
	var meta_test map[string]string = make(map[string]string)
	meta_test["version"] = "v1.0"
	meta_test["copyright"] = "uni-ledger"
	g_component.AddMetaAttribute(meta_test)
	if len(g_component.MetaAttribute) != 2 {
		t.Error("AddMetrAttribute Error!")
	}
	//Test GetMetaAttribute
	test_get_meta := g_component.GetMetaAttribute()
	if len(test_get_meta) != 2 {
		t.Error("GetMetaAttribute Error!")
	}
	if test_get_meta["version"] != "v1.0" {
		t.Error("MetaAttribute GetValue Error!")
	}

	fmt.Println("Component toString: ", g_component.ToString())

	//Test PropertyTable
	if len(g_component.GetPropertyTable()) != 6 {
		t.Error("AddProperty Error!")
	}
	common.AddProperty(g_component, g_component.PropertyTable, "Int", 100)
	common.AddProperty(g_component, g_component.PropertyTable, "String", "test property")
	if len(g_component.GetPropertyTable()) != 8 {
		t.Error("AddProperty Error!")
	}
	test_property := g_component.GetProperty("Int").(property.PropertyT)
	if test_property.GetValue().(int) != 100 {
		t.Error("GetProperty Error!")
	}
	//Test GetItem()
	v_cname := g_component.GetItem(_Cname)
	if v_cname.(string) != "TestComponent" {
		t.Error("GetItem[Cname] Error!")
	}
	v_ctype := g_component.GetItem(_Ctype)
	if v_ctype.(string) != "Component_Unknown" {
		t.Error("GetItem[Cname] Error!")
	}
	v_caption := g_component.GetItem(_Caption)
	if v_caption.(string) != "GeneralComponent" {
		t.Error("GetItem[Caption] Error!")
	}
	v_description := g_component.GetItem(_Description)
	if v_description.(string) != "GeneralComponent Test" {
		t.Error("GetItem[Description] Error!")
	}
	//Get Value By reflect
	fmt.Println(reflect.TypeOf(g_component))
	v_refl_object := reflect.ValueOf(g_component).Elem()
	v_refl_field := v_refl_object.FieldByName("MetaAttribute")
	fmt.Println(v_refl_field.Kind(), v_refl_field.Type(), v_refl_field.Interface())

	v_metadata := g_component.GetItem(_MetaAttribute)
	fmt.Println(v_metadata)

	fmt.Println("GetItem[version]", GetItem2(v_refl_field, "version"))
	fmt.Println("GetItem[copyright]", GetItem2(v_refl_field, "copyright"))
	fmt.Println(v_refl_field.Kind() == reflect.Map)
	fmt.Println(v_refl_field.MapKeys())
	fmt.Println(v_refl_field.MapIndex(reflect.ValueOf("version")))
	fmt.Println(v_refl_field.MapIndex(reflect.ValueOf("copyright")))
}

func GetItem2(p_object reflect.Value, p_item string) interface{} {
	switch p_object.Kind() {
	case reflect.Map:
		return (p_object.MapIndex(reflect.ValueOf(p_item)))
	}
	return nil
}
