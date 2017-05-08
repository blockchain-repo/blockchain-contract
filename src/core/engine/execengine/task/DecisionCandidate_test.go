package task

import (
	"testing"
	"unicontract/src/core/engine/execengine/property"
	"fmt"
)

func TestDecisionCandidate(t *testing.T){
	des := &DecisionCandidate{}
	//test InitDecisionCandidate
	des.InitDecisionCandidate()
	des.SetCname("DecisionCandidate")
	des.SetCaption("DecisionCandidate")
	des.SetDescription("test DecisionCandidate")
	if des.GetName() != "DecisionCandidate" {
		t.Error("DecisionCandidate Init Error, check name!")
	}
	if des.GetCaption() != "DecisionCandidate" {
		t.Error("DecisionCandidate Init Error, check caption!")
	}
	if des.GetDescription() != "test DecisionCandidate" {
		t.Error("DecisionCandidate Init Error, check description!")
	}
	//test AddText
	var test_text []string = []string{"add text test 1", "add text test 2", "add text test 3"}
	des.AddText(test_text)
	text_property := des.PropertyTable[_Text].(property.PropertyT)
	if len(text_property.GetValue().(map[string]string)) != 3{
		t.Error("AddText Error!")
	}
	map_text := text_property.GetValue().(map[string]string)
	if map_text["add text test 1"] != "add text test 1" {
		t.Error("AddText Error, text[0] value Error!")
	}
	//test ShowText
	des.ShowText()

	//test AddSupportArgument
	var support_1 string = "support1"
	var support_2 string = "support2"
	des.AddSupportArgument(support_1)
	des.AddSupportArgument(support_2)
	support_property := des.PropertyTable[_SupportArguments].(property.PropertyT)
	fmt.Println(support_property)
	if len(support_property.GetValue().(map[string]string)) != 2 {
		t.Error("AddSupportArgument Error!")
	}
	//test AddAgainstArgument
	var against_1 string = "against1"
	var against_2 string = "against2"
	var against_3 string = "against3"
	des.AddAgainstArgument(against_1)
	des.AddAgainstArgument(against_2)
	des.AddAgainstArgument(against_3)
	against_property := des.PropertyTable[_AgainstArguments].(property.PropertyT)
	fmt.Println(against_property)
	if len(against_property.GetValue().(map[string]string)) != 3 {
		t.Error("AddAgainstArgument Error!")
	}
}



