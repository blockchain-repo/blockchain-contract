package model

import (
	"encoding/json"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/protos"
)

// transfer contractModel string to contract(proto)
func FromContractModelStrToContractProto(contractModelStr string) (*protos.Contract, error) {
	// 1. to contractModel
	var contractModel ContractModel
	err := json.Unmarshal([]byte(contractModelStr), &contractModel)
	if err != nil {
		uniledgerlog.Error("error Unmarshal contractModelStr", err)
		return nil, err
	}
	// 2. to contract
	contract, err := FromContractModelToContractProto(contractModel)
	if err != nil {
		uniledgerlog.Error("error FromContractModelStrToContractProto", err)
		return contract, err
	}

	return contract, nil
}

//transfer contractModel to ContractProto
func FromContractModelToContractProto(contract ContractModel) (*protos.Contract, error) {

	contractProto := &protos.Contract{}
	// contract ID
	contractProto.Id = contract.Id

	// contractHead
	model_contractHead := contract.ContractHead
	if model_contractHead != nil {
		contractProto.ContractHead = &protos.ContractHead{
			MainPubkey:      model_contractHead.MainPubkey,
			Version:         model_contractHead.Version,
			AssignTime:      model_contractHead.AssignTime,
			OperateTime:     model_contractHead.AssignTime,
			ConsensusResult: model_contractHead.ConsensusResult,
		}
	}

	/************************ contractBody start ***************************/
	model_contractBody := contract.ContractBody
	if model_contractBody == nil {
		return contractProto, nil
	}

	// contractBody.contractAssests
	model_contractAssets := contract.ContractBody.ContractAssets

	contractAssets := make([]*protos.ContractAsset, len(model_contractAssets))
	if model_contractAssets == nil {
		contractAssets = nil
	} else {
		for i := 0; i < len(model_contractAssets); i++ {
			contractAssets[i] = &protos.ContractAsset{
				AssetId:     model_contractAssets[i].AssetId,
				Name:        model_contractAssets[i].Name,
				Caption:     model_contractAssets[i].Caption,
				Description: model_contractAssets[i].Description,
				Unit:        model_contractAssets[i].Unit,
				Amount:      model_contractAssets[i].Amount,
				MetaData:    model_contractAssets[i].MetaData,
			}
		}
	}

	// contractBody.ContractSignatures
	model_contractSignatures := contract.ContractBody.ContractSignatures
	contractSignatures := make([]*protos.ContractSignature, len(model_contractSignatures))
	if model_contractSignatures == nil {
		contractSignatures = nil
	} else {
		for i := 0; i < len(model_contractSignatures); i++ {
			contractSignatures[i] = &protos.ContractSignature{
				OwnerPubkey:   model_contractSignatures[i].OwnerPubkey,
				Signature:     model_contractSignatures[i].Signature,
				SignTimestamp: model_contractSignatures[i].SignTimestamp,
			}
		}
	}

	/************************ contractBody.ContractComponents start ***************************/
	contractComponents := contractComponentConvertToProto(contract.ContractBody.ContractComponents)
	//******** contractBody.ContractComponents end ***************************/

	//contractBody
	contractProto.ContractBody = &protos.ContractBody{
		ContractId:         model_contractBody.ContractId,
		Cname:              model_contractBody.Cname,
		Ctype:              model_contractBody.Ctype,
		Caption:            model_contractBody.Caption,
		Description:        model_contractBody.Description,
		ContractState:      model_contractBody.ContractState,
		Creator:            model_contractBody.Creator,
		CreateTime:         model_contractBody.CreateTime,
		StartTime:          model_contractBody.StartTime,
		EndTime:            model_contractBody.EndTime,
		ContractOwners:     model_contractBody.ContractOwners,
		ContractAssets:     contractAssets,
		ContractSignatures: contractSignatures,
		ContractComponents: contractComponents,
		MetaAttribute:      model_contractBody.MetaAttribute,
		NextTasks:          model_contractBody.NextTasks,
		ContractProductId:  model_contractBody.ContractProductId,
		ContractTemplateId: model_contractBody.ContractTemplateId,
	}
	/************************ contractBody end ***************************/

	return contractProto, nil
}

func contractComponentConvertToProto(components []*ContractComponent) []*protos.ContractComponent {
	/************************ contractBody.ContractComponents start ***************************/
	contractComponents := make([]*protos.ContractComponent, len(components))
	if components == nil {
		contractComponents = nil
	} else {
		for i := 0; i < len(components); i++ {

			/************************ contractBody.ContractComponent.ComponentsExpression start ***************************/
			preConditions := componentExpressionConvertToProto(components[i].PreCondition)
			completeConditions := componentExpressionConvertToProto(components[i].CompleteCondition)
			discardConditions := componentExpressionConvertToProto(components[i].DiscardCondition)
			dataValueSetterExpressionList := componentExpressionConvertToProto(components[i].DataValueSetterExpressionList)
			/************************ contractBody.ContractComponent.ComponentsExpression end ***************************/

			/************************ contractBody.ContractComponent.ComponentData start ***************************/
			componentDataList := componentDataConvertToProto(components[i].DataList)
			/************************ contractBody.ContractComponent.ComponentData end ***************************/

			/************************ contractBody.ContractComponent start ***************************/
			candidateList := contractComponentSubConvertToProto(components[i].CandidateList)
			//decisionResult := contractComponentSubConvertToProto(components[i].DecisionResult)
			/************************ contractBody.ContractComponent end ***************************/

			/************************ contractBody.ContractComponent.SelectBranchExpression start ***************************/
			selectBranchExpressions := componentSelectBranchesConvertToProto(components[i].SelectBranches)
			/************************ contractBody.ContractComponent.SelectBranchExpression end ***************************/

			contractComponents[i] = &protos.ContractComponent{
				Cname:                         components[i].Cname,
				Ctype:                         components[i].Ctype,
				Caption:                       components[i].Caption,
				Description:                   components[i].Description,
				State:                         components[i].State,
				PreCondition:                  preConditions,
				CompleteCondition:             completeConditions,
				DiscardCondition:              discardConditions,
				NextTasks:                     components[i].NextTasks,
				DataList:                      componentDataList,
				CandidateList:                 candidateList,
				DataValueSetterExpressionList: dataValueSetterExpressionList,
				TaskList:                      components[i].TaskList,
				TaskExecuteIdx:                components[i].TaskExecuteIdx,
				TaskId:                        components[i].TaskId,
				SelectBranches:                selectBranchExpressions,
				MetaAttribute:                 components[i].MetaAttribute,
			}
		}
	}
	/************************ contractBody.ContractComponent end ***************************/
	return contractComponents
}

func contractComponentSubConvertToProto(componentSubs []*ContractComponentSub) []*protos.ContractComponentSub {
	/************************ contractBody.ContractComponents start ***************************/
	contractComponents := make([]*protos.ContractComponentSub, len(componentSubs))
	if componentSubs == nil {
		contractComponents = nil
	} else {
		for i := 0; i < len(componentSubs); i++ {

			/************************ contractBody.ContractComponentSub.ComponentsExpression start ***************************/
			preConditions := componentExpressionConvertToProto(componentSubs[i].PreCondition)
			completeConditions := componentExpressionConvertToProto(componentSubs[i].CompleteCondition)
			discardConditions := componentExpressionConvertToProto(componentSubs[i].DiscardCondition)
			dataValueSetterExpressionList := componentExpressionConvertToProto(componentSubs[i].DataValueSetterExpressionList)
			/************************ contractBody.ContractComponentSub.ComponentsExpression end ***************************/

			/************************ contractBody.ContractComponentSub.ComponentData start ***************************/
			componentDataList := componentDataConvertToProto(componentSubs[i].DataList)
			/************************ contractBody.ContractComponentSub.ComponentData end ***************************/

			/************************ contractBody.ContractComponentSub.SelectBranchExpression start ***************************/
			selectBranchExpressions := componentSelectBranchesConvertToProto(componentSubs[i].SelectBranches)
			/************************ contractBody.ContractComponentSub.SelectBranchExpression end ***************************/

			contractComponents[i] = &protos.ContractComponentSub{
				Cname:                         componentSubs[i].Cname,
				Ctype:                         componentSubs[i].Ctype,
				Caption:                       componentSubs[i].Caption,
				Description:                   componentSubs[i].Description,
				State:                         componentSubs[i].State,
				PreCondition:                  preConditions,
				CompleteCondition:             completeConditions,
				DiscardCondition:              discardConditions,
				NextTasks:                     componentSubs[i].NextTasks,
				DataList:                      componentDataList,
				DataValueSetterExpressionList: dataValueSetterExpressionList,
				TaskList:                      componentSubs[i].TaskList,
				SupportArguments:              componentSubs[i].SupportArguments,
				AgainstArguments:              componentSubs[i].AgainstArguments,
				Text:                          componentSubs[i].Text,
				TaskExecuteIdx:                componentSubs[i].TaskExecuteIdx,
				TaskId:                        componentSubs[i].TaskId,
				SelectBranches:                selectBranchExpressions,
				Result:                        componentSubs[i].Result,
				SupportNum:                    componentSubs[i].SupportNum,
				AgainstNum:                    componentSubs[i].AgainstNum,
			}
		}
	}
	/************************ contractBody.ContractComponentSub end ***************************/
	return contractComponents
}

func componentSelectBranchesConvertToProto(selectBranches []*SelectBranchExpression) []*protos.SelectBranchExpression {
	selectBranchExpressions := make([]*protos.SelectBranchExpression, len(selectBranches))
	if selectBranches == nil {
		selectBranchExpressions = nil
	} else {
		for j := 0; j < len(selectBranches); j++ {
			selectBranchExpressions[j] = &protos.SelectBranchExpression{
				BranchExpressionStr:   selectBranches[j].BranchExpressionStr,
				BranchExpressionValue: selectBranches[j].BranchExpressionValue,
			}
		}
	}
	return selectBranchExpressions
}

// convert the model.ComponentsExpression to proto componentExpression
func componentExpressionConvertToProto(expression []*ComponentsExpression) []*protos.ComponentsExpression {
	componentExpressions := make([]*protos.ComponentsExpression, len(expression))
	if expression == nil {
		componentExpressions = nil
	} else {
		for j := 0; j < len(expression); j++ {
			/************************ contractBody ExpressionResult start ***************************/
			model_preExpressionResult := expression[j].ExpressionResult
			var expressionResult *protos.ExpressionResult
			if model_preExpressionResult == nil {
				expressionResult = nil
			} else {
				expressionResult = &protos.ExpressionResult{
					Message: model_preExpressionResult.Message,
					Code:    model_preExpressionResult.Code,
					Data:    model_preExpressionResult.Data,
					OutPut:  model_preExpressionResult.OutPut,
				}
			}
			/************************ contractBody ExpressionResult end ***************************/

			componentExpressions[j] = &protos.ComponentsExpression{
				Cname:            expression[j].Cname,
				Ctype:            expression[j].Ctype,
				Caption:          expression[j].Caption,
				Description:      expression[j].Description,
				ExpressionStr:    expression[j].ExpressionStr,
				ExpressionResult: expressionResult,
				LogicValue:       expression[j].LogicValue,
				MetaAttribute:    expression[j].MetaAttribute,
			}
		}
	}
	return componentExpressions
}

func componentDataSubsConvertToProto(data *ComponentDataSub) *protos.ComponentDataSub {
	var componentDataSubs *protos.ComponentDataSub
	if data == nil {
		componentDataSubs = nil
	} else {
		componentDataSubs = &protos.ComponentDataSub{
			Cname:              data.Cname,
			Ctype:              data.Ctype,
			Caption:            data.Caption,
			Description:        data.Description,
			ModifyDate:         data.ModifyDate,
			HardConvType:       data.HardConvType,
			Category:           data.Category,
			Mandatory:          data.Mandatory,
			Unit:               data.Unit,
			Options:            data.Options,
			Format:             data.Format,
			ValueInt:           data.ValueInt,
			ValueUint:          data.ValueUint,
			ValueFloat:         data.ValueFloat,
			ValueString:        data.ValueString,
			DefaultValueInt:    data.DefaultValueInt,
			DefaultValueUint:   data.DefaultValueUint,
			DefaultValueFloat:  data.DefaultValueFloat,
			DefaultValueString: data.DefaultValueString,
			DataRangeInt:       data.DataRangeInt,
			DataRangeUint:      data.DataRangeUint,
			DataRangeFloat:     data.DataRangeFloat,
		}
	}
	return componentDataSubs
}

func componentDataConvertToProto(datas []*ComponentData) []*protos.ComponentData {
	componentData := make([]*protos.ComponentData, len(datas))
	if datas == nil {
		componentData = nil
	} else {
		for i := 0; i < len(datas); i++ {
			//parent := componentDataSubsConvertToProto(datas[i].Parent)
			componentData[i] = &protos.ComponentData{
				Cname:        datas[i].Cname,
				Ctype:        datas[i].Ctype,
				Caption:      datas[i].Caption,
				Description:  datas[i].Description,
				ModifyDate:   datas[i].ModifyDate,
				HardConvType: datas[i].HardConvType,
				Category:     datas[i].Category,
				//Parent:             parent,
				Mandatory:          datas[i].Mandatory,
				Unit:               datas[i].Unit,
				Options:            datas[i].Options,
				Format:             datas[i].Format,
				ValueInt:           datas[i].ValueInt,
				ValueUint:          datas[i].ValueUint,
				ValueFloat:         datas[i].ValueFloat,
				ValueString:        datas[i].ValueString,
				DefaultValueInt:    datas[i].DefaultValueInt,
				DefaultValueUint:   datas[i].DefaultValueUint,
				DefaultValueFloat:  datas[i].DefaultValueFloat,
				DefaultValueString: datas[i].DefaultValueString,
				DataRangeInt:       datas[i].DataRangeInt,
				DataRangeUint:      datas[i].DataRangeUint,
				DataRangeFloat:     datas[i].DataRangeFloat,
				Value:              datas[i].Value,
				DefaultValue:       datas[i].DefaultValue,
				ValueBool:          datas[i].ValueBool,
				DefaultValueBool:   datas[i].DefaultValueBool,
			}
		}
	}
	return componentData
}
