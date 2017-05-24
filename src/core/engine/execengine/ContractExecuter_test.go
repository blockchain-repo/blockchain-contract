package execengine

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/contract"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/property"
)

func ReadFile(p_filepath string) ([]byte, error) {
	f, err := os.Open(p_filepath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func PrintContract(v_contract *contract.CognitiveContract) {
	fmt.Println("===========================constract print=============================")
	fmt.Println("========constract ID=========")
	fmt.Println("Id: ", v_contract.Id)
	fmt.Println("  PropertyTable[_Id]: ", v_contract.GetPropertyTable()["_Id"])
	fmt.Println("========constract Head=========")
	fmt.Println("MainPubkey: ", v_contract.ContractHead.MainPubkey)
	fmt.Println("  PropertyTable[_MainPubkey]: ", v_contract.GetPropertyTable()["_MainPubkey"])
	fmt.Println("Timestamp: ", v_contract.ContractHead.Timestamp)
	fmt.Println("  PropertyTable[_Timestamp]: ", v_contract.GetPropertyTable()["_Timestamp"])
	fmt.Println("Version: ", v_contract.ContractHead.Version)
	fmt.Println("  PropertyTable[_Version]: ", v_contract.GetPropertyTable()["_Version"])
	fmt.Println("========constract Body=========")
	fmt.Println("ContractId: ", v_contract.ContractBody.ContractId)
	fmt.Println("  PropertyTable[_ContractId]: ", v_contract.GetPropertyTable()["_ContractId"])
	fmt.Println("Cname: ", v_contract.ContractBody.Cname)
	fmt.Println("  PropertyTable[_Cname]: ", v_contract.GetPropertyTable()["_Cname"])
	fmt.Println("Ctype: ", v_contract.ContractBody.Ctype)
	fmt.Println("  PropertyTable[_Ctype]: ", v_contract.GetPropertyTable()["_Ctype"])
	fmt.Println("Caption: ", v_contract.ContractBody.Caption)
	fmt.Println("  PropertyTable[_Caption]: ", v_contract.GetPropertyTable()["_Caption"])
	fmt.Println("Description: ", v_contract.ContractBody.Description)
	fmt.Println("  PropertyTable[_Description]: ", v_contract.GetPropertyTable()["_Description"])
	fmt.Println("ContractState: ", v_contract.ContractBody.ContractState)
	fmt.Println("  PropertyTable[_ContractState]: ", v_contract.GetPropertyTable()["_ContractState"])
	fmt.Println("Creator: ", v_contract.ContractBody.Creator)
	fmt.Println("  PropertyTable[_Creator]: ", v_contract.GetPropertyTable()["_Creator"])
	fmt.Println("CreateTime: ", v_contract.ContractBody.CreateTime)
	fmt.Println("  PropertyTable[_CreateTime]: ", v_contract.GetPropertyTable()["_CreateTime"])
	fmt.Println("StartTime: ", v_contract.ContractBody.StartTime)
	fmt.Println("  PropertyTable[_StartTime]: ", v_contract.GetPropertyTable()["_StartTime"])
	fmt.Println("EndTime: ", v_contract.ContractBody.EndTime)
	fmt.Println("  PropertyTable[_EndTime]: ", v_contract.GetPropertyTable()["_EndTime"])
	fmt.Println("ContractOwners: ", v_contract.ContractBody.ContractOwners)
	fmt.Println("  PropertyTable[_ContractOwners]: ", v_contract.GetPropertyTable()["_ContractOwners"])
	fmt.Println("  All Owners: ")
	for p_idx, p_owner := range v_contract.ContractBody.ContractOwners {
		fmt.Println("  owner[", p_idx, "]: ", p_owner)
	}
	fmt.Println("")
	fmt.Println("ContractAssets: ", v_contract.ContractBody.ContractAssets)
	fmt.Println("  PropertyTable[_ContractAssets]: ", v_contract.GetPropertyTable()["_ContractAssets"])
	fmt.Println("  All Assets: ")
	for p_idx, p_assert := range v_contract.ContractBody.ContractAssets {
		fmt.Println("  Asset.AssetId[", p_idx, "]: ", p_assert.AssetId)
		fmt.Println("  Asset.Name[", p_idx, "]: ", p_assert.Name)
		fmt.Println("  Asset.Caption[", p_idx, "]: ", p_assert.Caption)
		fmt.Println("  Asset.Description[", p_idx, "]: ", p_assert.Description)
		fmt.Println("  Asset.Unit[", p_idx, "]: ", p_assert.Unit)
		fmt.Println("  Asset.Amount[", p_idx, "]: ", p_assert.Amount)
		fmt.Println("  Asset.MetaData[", p_idx, "]: ", p_assert.MetaData)
	}
	fmt.Println("ContractSignatures: ", v_contract.ContractBody.ContractSignatures)
	fmt.Println("  PropertyTable[_ContractSignatures]: ", v_contract.GetPropertyTable()["_ContractSignatures"])
	fmt.Println("  All Signatures: ")
	for p_idx, p_signature := range v_contract.ContractBody.ContractSignatures {
		fmt.Println("  Signatures.OwnerPubkey[", p_idx, "]: ", p_signature.OwnerPubkey)
		fmt.Println("  Signatures.Signature[", p_idx, "]: ", p_signature.Signature)
		fmt.Println("  Signatures.SignTimestamp[", p_idx, "]: ", p_signature.SignTimestamp)
	}
	fmt.Println("MetaAttribute: ", v_contract.ContractBody.MetaAttribute)
	fmt.Println("  PropertyTable[_MetaAttribute]: ", v_contract.GetPropertyTable()["_MetaAttribute"])
	fmt.Println("  All MetaAttribute: ")
	for p_key, p_value := range v_contract.ContractBody.MetaAttribute {
		fmt.Println("  Attribute[", p_key, "]", p_value)
	}
	/*
		fmt.Println("Contract: ", v_contract.Contract.GetVersion())
		fmt.Println("Contract: ", v_contract.Contract.GetCopyRight())
	*/
	fmt.Println("ComponentTable: ", len(v_contract.ComponentTable.CompTable))
	fmt.Println("ComponentTable[Task]: ", len(v_contract.ComponentTable.GetComponentByType(constdef.ComponentType[constdef.Component_Task])))
	for p_key, p_value := range v_contract.ComponentTable.GetComponentByType(constdef.ComponentType[constdef.Component_Task]) {
		fmt.Println("Component[", p_key, "]", p_value)
	}
	fmt.Println("ComponentTable[Data]: ", len(v_contract.ComponentTable.GetComponentByType(constdef.ComponentType[constdef.Component_Data])))
	for p_key, p_value := range v_contract.ComponentTable.GetComponentByType(constdef.ComponentType[constdef.Component_Data]) {
		fmt.Println("Component[", p_key, "]", p_value)
	}
	fmt.Println("ComponentTable[Expression]: ", len(v_contract.ComponentTable.GetComponentByType(constdef.ComponentType[constdef.Component_Expression])))
	for p_key, p_value := range v_contract.ComponentTable.GetComponentByType(constdef.ComponentType[constdef.Component_Expression]) {
		fmt.Println("Component[", p_key, "]", p_value)
	}

	fmt.Println("Compontent Tasks State:")
	for _, p_value := range v_contract.ComponentTable.GetComponentByType(constdef.ComponentType[constdef.Component_Task]) {
		for m_key, m_value := range p_value {
			fmt.Println("Task[", m_key, "] : ", m_value.(inf.ITask).GetState())
		}
	}
}

func TestContractAllLife_New(t *testing.T) {
	//Read from file
	var file_path string = "./ContractDemo.json"
	v_byte, err := ReadFile(file_path)
	if err != nil {
		t.Error("Read File Error!")
	}
	fmt.Println(v_byte)
	//1. Test Load
	fmt.Println("=============Test Load========================================================")
	v_contract_execute := NewContractExecuter()
	err = v_contract_execute.Load(string(v_byte))
	if err != nil {
		t.Error("Load Error:", err)
		return
	}
	cname_property := v_contract_execute.contract_executer.PropertyTable["_Cname"].(property.PropertyT)
	if v_contract_execute.contract_executer.GetName() != "contract_mobilecallback" || cname_property.GetValue().(string) != "contract_mobilecallback" {
		t.Error("Load Error, GetName Error!")
	}
	ctype_property := v_contract_execute.contract_executer.PropertyTable["_Ctype"].(property.PropertyT)
	if v_contract_execute.contract_executer.GetCtype() != constdef.ComponentType[constdef.Component_Contract] || ctype_property.GetValue().(string) != constdef.ComponentType[constdef.Component_Contract] {
		t.Error("Load Error, GetCtype Error!")
	}
	/*
		if len(t_contract.GetContractComponents()) != 2 {
			t.Error("Load Error, ContractComponents[Describe] Error!")
		}

		task_component_table := t_contract.GetComponentByType(constdef.ComponentType[constdef.Component_Task])
		fmt.Println("component_table: ", task_component_table)
		if len(task_component_table) != 2 {
			t.Error("Load Error, Component Table[task] Error!")
		}

		property_table := t_contract.GetPropertyTable()
		fmt.Println("property_table: ", property_table)
		if len(property_table) < 10 {
			t.Error("Load Error, Property Table Error!")
		}
	*/
	PrintContract(v_contract_execute.contract_executer)

	//2. Test Export
	fmt.Println("=============Test Export Json========================================================")
	r_str_json, err := v_contract_execute.ExportToJson()
	if err != nil {
		t.Error("Export Error!")
	}
	fmt.Println("Export Json Result: ", r_str_json)

	//2. Test Export
	fmt.Println("=============Test Export Text========================================================")
	r_str_text, err := v_contract_execute.ExportToText()
	if err != nil {
		t.Error("Export Error!")
	}
	fmt.Println("Export Text Result: \n", r_str_text)
	PrintContract(v_contract_execute.contract_executer)

	//3. Test Start
	fmt.Println("=============Test Start ========================================================")
	_, err = v_contract_execute.Start()
	fmt.Println(err)
}

func TestContractAllLife_HasInprocess(t *testing.T) {
	//Read from file
	var file_path string = "./ContractDemo_Inprocess.json"
	v_byte, err := ReadFile(file_path)
	if err != nil {
		t.Error("Read File Error!")
	}
	fmt.Println(v_byte)
	//1. Test Load
	fmt.Println("=============Test Load========================================================")
	v_contract_execute := NewContractExecuter()
	err = v_contract_execute.Load(string(v_byte))
	if err != nil {
		t.Error("Load Error:", err)
		return
	}
	cname_property := v_contract_execute.contract_executer.PropertyTable["_Cname"].(property.PropertyT)
	if v_contract_execute.contract_executer.GetName() != "contract_mobilecallback" || cname_property.GetValue().(string) != "contract_mobilecallback" {
		t.Error("Load Error, GetName Error!")
	}
	ctype_property := v_contract_execute.contract_executer.PropertyTable["_Ctype"].(property.PropertyT)
	if v_contract_execute.contract_executer.GetCtype() != constdef.ComponentType[constdef.Component_Contract] || ctype_property.GetValue().(string) != constdef.ComponentType[constdef.Component_Contract] {
		t.Error("Load Error, GetCtype Error!")
	}
	PrintContract(v_contract_execute.contract_executer)

	//2. Test Export
	fmt.Println("=============Test Export Json========================================================")
	r_str_json, err := v_contract_execute.ExportToJson()
	if err != nil {
		t.Error("Export Error!")
	}
	fmt.Println("Export Json Result: ", r_str_json)

	//2. Test Export
	fmt.Println("=============Test Export Text========================================================")
	r_str_text, err := v_contract_execute.ExportToText()
	if err != nil {
		t.Error("Export Error!")
	}
	fmt.Println("Export Text Result: \n", r_str_text)
	PrintContract(v_contract_execute.contract_executer)

	//3. Test Start
	fmt.Println("=============Test Start ========================================================")
	_, err = v_contract_execute.Start()
	fmt.Println(err)
}

func TestContractAllLife_HasComplete(t *testing.T) {
	//Read from file
	var file_path string = "./ContractDemo_Complete.json"
	v_byte, err := ReadFile(file_path)
	if err != nil {
		t.Error("Read File Error!")
	}
	fmt.Println(v_byte)
	//1. Test Load
	fmt.Println("=============Test Load========================================================")
	v_contract_execute := NewContractExecuter()
	err = v_contract_execute.Load(string(v_byte))
	if err != nil {
		t.Error("Load Error:", err)
		return
	}
	cname_property := v_contract_execute.contract_executer.PropertyTable["_Cname"].(property.PropertyT)
	if v_contract_execute.contract_executer.GetName() != "contract_mobilecallback" || cname_property.GetValue().(string) != "contract_mobilecallback" {
		t.Error("Load Error, GetName Error!")
	}
	ctype_property := v_contract_execute.contract_executer.PropertyTable["_Ctype"].(property.PropertyT)
	if v_contract_execute.contract_executer.GetCtype() != constdef.ComponentType[constdef.Component_Contract] || ctype_property.GetValue().(string) != constdef.ComponentType[constdef.Component_Contract] {
		t.Error("Load Error, GetCtype Error!")
	}
	PrintContract(v_contract_execute.contract_executer)

	//2. Test Export
	fmt.Println("=============Test Export Json========================================================")
	r_str_json, err := v_contract_execute.ExportToJson()
	if err != nil {
		t.Error("Export Error!")
	}
	fmt.Println("Export Json Result: ", r_str_json)

	//2. Test Export
	fmt.Println("=============Test Export Text========================================================")
	r_str_text, err := v_contract_execute.ExportToText()
	if err != nil {
		t.Error("Export Error!")
	}
	fmt.Println("Export Text Result: \n", r_str_text)
	PrintContract(v_contract_execute.contract_executer)

	//3. Test Start
	fmt.Println("=============Test Start ========================================================")
	_, err = v_contract_execute.Start()
	fmt.Println(err)
}

func TestContractAllLife_HasDigcard(t *testing.T) {
	//Read from file
	var file_path string = "./ContractDemo_Digcard.json"
	v_byte, err := ReadFile(file_path)
	if err != nil {
		t.Error("Read File Error!")
	}
	fmt.Println(v_byte)
	//1. Test Load
	fmt.Println("=============Test Load========================================================")
	v_contract_execute := NewContractExecuter()
	err = v_contract_execute.Load(string(v_byte))
	if err != nil {
		t.Error("Load Error:", err)
		return
	}
	cname_property := v_contract_execute.contract_executer.PropertyTable["_Cname"].(property.PropertyT)
	if v_contract_execute.contract_executer.GetName() != "contract_mobilecallback" || cname_property.GetValue().(string) != "contract_mobilecallback" {
		t.Error("Load Error, GetName Error!")
	}
	ctype_property := v_contract_execute.contract_executer.PropertyTable["_Ctype"].(property.PropertyT)
	if v_contract_execute.contract_executer.GetCtype() != constdef.ComponentType[constdef.Component_Contract] || ctype_property.GetValue().(string) != constdef.ComponentType[constdef.Component_Contract] {
		t.Error("Load Error, GetCtype Error!")
	}
	PrintContract(v_contract_execute.contract_executer)

	//2. Test Export
	fmt.Println("=============Test Export Json========================================================")
	r_str_json, err := v_contract_execute.ExportToJson()
	if err != nil {
		t.Error("Export Error!")
	}
	fmt.Println("Export Json Result: ", r_str_json)

	//2. Test Export
	fmt.Println("=============Test Export Text========================================================")
	r_str_text, err := v_contract_execute.ExportToText()
	if err != nil {
		t.Error("Export Error!")
	}
	fmt.Println("Export Text Result: \n", r_str_text)
	PrintContract(v_contract_execute.contract_executer)

	//3. Test Start
	fmt.Println("=============Test Start ========================================================")
	_, err = v_contract_execute.Start()
	fmt.Println(err)
}
