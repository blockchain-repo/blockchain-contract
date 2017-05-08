package component

import (
	"testing"
	"fmt"
	"unicontract/src/core/engine/execengine/property"
	"unicontract/src/core/engine/execengine/constdef"
)

func TestGeneralComponent(t *testing.T) {
	g_component := &GeneralComponent{Cname:"GeneralComponent",Caption:"GeneralComponent",Description:"GeneralComponent Test"}
	g_component.InitGeneralComponent()

	if g_component.GetCname() != "GeneralComponent" || g_component.Cname != "GeneralComponent"{
		t.Error("InitGeneralComponent Error,GetName Error!")
	}

	if g_component.GetCtype() != constdef.ComponentType[constdef.Component_Unknown] || g_component.Ctype != constdef.ComponentType[constdef.Component_Unknown]{
		t.Error("InitGeneralComponent Error,GetCtype Error!")
	}
	if g_component.GetCaption() != "GeneralComponent" || g_component.Caption != "GeneralComponent"{
		t.Error("InitGeneralComponent Error,GetCaption Error!")
	}
	if g_component.GetDescription() != "GeneralComponent Test" || g_component.Description != "GeneralComponent Test"{
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
	if len(g_component.GetPropertyTable()) != 6{
		t.Error("AddProperty Error!")
	}
	g_component.AddProperty(g_component,"Int", 100)
	g_component.AddProperty(g_component,"String", "test property")
	if len(g_component.GetPropertyTable()) != 8{
		t.Error("AddProperty Error!")
	}
	test_property := g_component.GetProperty("Int").(property.PropertyT)
	if test_property.GetValue().(int) != 100 {
		t.Error("GetProperty Error!")
	}
}
