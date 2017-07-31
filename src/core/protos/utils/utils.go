package utils

import (
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/model"
	"unicontract/src/core/protos/api"
)

//transfer contract(proto) to contractModel
func FromProtoContractToContractModel(contract protos.Contract) model.ContractModel {
	uniledgerlog.Info("convert the protocol data to model <contract>")
	uniledgerlog.Info("deal start")
	uniledgerlog.Info("deal end")
	var contractModel model.ContractModel

	// contract ID
	contractModel.Id = contract.Id

	// contractHead
	proto_contractHead := contract.ContractHead
	if proto_contractHead != nil {
		contractModel.ContractHead = &model.ContractHead{
			MainPubkey:      proto_contractHead.MainPubkey,
			Version:         proto_contractHead.Version,
			AssignTime:      proto_contractHead.AssignTime,
			OperateTime:     proto_contractHead.AssignTime,
			ConsensusResult: proto_contractHead.ConsensusResult,
		}
	}

	/************************ contractBody start ***************************/
	proto_contractBody := contract.ContractBody
	if proto_contractBody == nil {
		return contractModel
	}

	// contractBody.contractAssests
	proto_contractAssets := contract.ContractBody.ContractAssets
	var contractAssets []*model.ContractAsset
	if proto_contractAssets == nil {
		contractAssets = nil
	} else {
		for i := 0; i < len(proto_contractAssets); i++ {
			contractAssets[i] = &model.ContractAsset{
				AssetId:     proto_contractAssets[i].AssetId,
				Name:        proto_contractAssets[i].Name,
				Caption:     proto_contractAssets[i].Caption,
				Description: proto_contractAssets[i].Description,
				Unit:        proto_contractAssets[i].Unit,
				Amount:      proto_contractAssets[i].Amount,
				MetaData:    proto_contractAssets[i].MetaData,
			}
		}
	}

	// contractBody.ContractSignatures
	proto_contractSignatures := contract.ContractBody.ContractSignatures
	var contractSignatures []*model.ContractSignature
	if proto_contractSignatures == nil {
		contractSignatures = nil
	} else {
		for i := 0; i < len(proto_contractSignatures); i++ {
			contractSignatures[i] = &model.ContractSignature{
				OwnerPubkey:   proto_contractSignatures[i].OwnerPubkey,
				Signature:     proto_contractSignatures[i].Signature,
				SignTimestamp: proto_contractSignatures[i].SignTimestamp,
			}
		}
	}

	/************************ contractBody.ContractComponents start ***************************/
	contractComponents := contractComponentConvertToModel(contract.ContractBody.ContractComponents)
	//******** contractBody.ContractComponents end ***************************/

	//contractBody
	contractModel.ContractBody = &model.ContractBody{
		ContractId:         proto_contractBody.ContractId,
		Cname:              proto_contractBody.Cname,
		Ctype:              proto_contractBody.Ctype,
		Caption:            proto_contractBody.Caption,
		Description:        proto_contractBody.Description,
		ContractState:      proto_contractBody.ContractState,
		Creator:            proto_contractBody.Creator,
		CreateTime:         proto_contractBody.CreateTime,
		StartTime:          proto_contractBody.StartTime,
		EndTime:            proto_contractBody.EndTime,
		ContractOwners:     proto_contractBody.ContractOwners,
		ContractAssets:     contractAssets,
		ContractSignatures: contractSignatures,
		ContractComponents: contractComponents,
		MetaAttribute:      proto_contractBody.MetaAttribute,
		NextTasks:          proto_contractBody.NextTasks,
	}
	/************************ contractBody end ***************************/

	return contractModel
}

func contractComponentConvertToModel(components []*protos.ContractComponent) []*model.ContractComponent {
	/************************ contractBody.ContractComponents start ***************************/
	var contractComponents []*model.ContractComponent
	if components == nil {
		contractComponents = nil
	} else {
		for i := 0; i < len(components); i++ {

			/************************ contractBody.ContractComponent.ComponentsExpression start ***************************/
			preConditions := componentExpressionConvertToModel(components[i].PreCondition)
			completeConditions := componentExpressionConvertToModel(components[i].CompleteCondition)
			discardConditions := componentExpressionConvertToModel(components[i].DiscardCondition)
			dataValueSetterExpressionList := componentExpressionConvertToModel(components[i].DataValueSetterExpressionList)
			/************************ contractBody.ContractComponent.ComponentsExpression end ***************************/

			/************************ contractBody.ContractComponent.ComponentData start ***************************/
			componentDataList := componentDatasConvertToModel(components[i].DataList)
			/************************ contractBody.ContractComponent.ComponentData end ***************************/

			/************************ contractBody.ContractComponent start ***************************/
			candidateList := contractComponentSubConvertToModel(components[i].CandidateList)
			decisionResult := contractComponentSubConvertToModel(components[i].DecisionResult)
			/************************ contractBody.ContractComponent end ***************************/

			/************************ contractBody.ContractComponent.SelectBranchExpression start ***************************/
			proto_selectBranchExpressions := components[i].SelectBranches
			var selectBranchExpressions []*model.SelectBranchExpression
			if proto_selectBranchExpressions == nil {
				selectBranchExpressions = nil
			} else {
				for j := 0; j < len(proto_selectBranchExpressions); j++ {
					selectBranchExpressions[j] = &model.SelectBranchExpression{
						BranchExpressionStr:   proto_selectBranchExpressions[j].BranchExpressionStr,
						BranchExpressionValue: proto_selectBranchExpressions[j].BranchExpressionValue,
					}
				}
			}
			/************************ contractBody.ContractComponent.SelectBranchExpression end ***************************/

			contractComponents[i] = &model.ContractComponent{
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
				DecisionResult:                decisionResult,
				DataValueSetterExpressionList: dataValueSetterExpressionList,
				TaskList:                      components[i].TaskList,
				SupportArguments:              components[i].SupportArguments,
				AgainstArguments:              components[i].AgainstArguments,
				Support:                       components[i].Support,
				Text:                          components[i].Text,
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

func contractComponentSubConvertToModel(componentSubs []*protos.ContractComponentSub) []*model.ContractComponentSub {
	/************************ contractBody.ContractComponents start ***************************/
	var contractComponents []*model.ContractComponentSub
	if componentSubs == nil {
		contractComponents = nil
	} else {
		for i := 0; i < len(componentSubs); i++ {

			/************************ contractBody.ContractComponentSub.ComponentsExpression start ***************************/
			preConditions := componentExpressionConvertToModel(componentSubs[i].PreCondition)
			completeConditions := componentExpressionConvertToModel(componentSubs[i].CompleteCondition)
			discardConditions := componentExpressionConvertToModel(componentSubs[i].DiscardCondition)
			dataValueSetterExpressionList := componentExpressionConvertToModel(componentSubs[i].DataValueSetterExpressionList)
			/************************ contractBody.ContractComponentSub.ComponentsExpression end ***************************/

			/************************ contractBody.ContractComponentSub.ComponentData start ***************************/
			componentDataList := componentDatasConvertToModel(componentSubs[i].DataList)
			/************************ contractBody.ContractComponentSub.ComponentData end ***************************/

			/************************ contractBody.ContractComponentSub.SelectBranchExpression start ***************************/
			proto_selectBranchExpressions := componentSubs[i].SelectBranches
			var selectBranchExpressions []*model.SelectBranchExpression
			if proto_selectBranchExpressions == nil {
				selectBranchExpressions = nil
			} else {
				for j := 0; j < len(proto_selectBranchExpressions); j++ {
					selectBranchExpressions[j] = &model.SelectBranchExpression{
						BranchExpressionStr:   proto_selectBranchExpressions[j].BranchExpressionStr,
						BranchExpressionValue: proto_selectBranchExpressions[j].BranchExpressionValue,
					}
				}
			}
			/************************ contractBody.ContractComponentSub.SelectBranchExpression end ***************************/

			contractComponents[i] = &model.ContractComponentSub{
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
				Support:                       componentSubs[i].Support,
				Text:                          componentSubs[i].Text,
				TaskExecuteIdx:                componentSubs[i].TaskExecuteIdx,
				TaskId:                        componentSubs[i].TaskId,
				SelectBranches:                selectBranchExpressions,
			}
		}
	}
	/************************ contractBody.ContractComponentSub end ***************************/
	return contractComponents
}

// convert the proto componentExpression to model.ComponentsExpression
func componentExpressionConvertToModel(expression []*protos.ComponentsExpression) []*model.ComponentsExpression {
	var componentExpressions []*model.ComponentsExpression
	if expression == nil {
		componentExpressions = nil
	} else {
		for j := 0; j < len(expression); j++ {
			/************************ contractBody ExpressionResult start ***************************/
			preExpressionResult := expression[j].ExpressionResult
			var expressionResult *model.ExpressionResult
			if preExpressionResult == nil {
				expressionResult = nil
			} else {
				expressionResult = &model.ExpressionResult{
					Message: preExpressionResult.Message,
					Code:    preExpressionResult.Code,
					Data:    preExpressionResult.Data,
					OutPut:  preExpressionResult.OutPut,
				}
			}
			/************************ contractBody ExpressionResult end ***************************/

			componentExpressions[j] = &model.ComponentsExpression{
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

func componentDataSubsConvertToModel(data *protos.ComponentDataSub) *model.ComponentDataSub {
	var componentDataSubs *model.ComponentDataSub
	if data == nil {
		componentDataSubs = nil
	} else {
		componentDataSubs = &model.ComponentDataSub{
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

func componentDatasConvertToModel(datas []*protos.ComponentData) []*model.ComponentData {
	var componentDatas []*model.ComponentData
	if datas == nil {
		componentDatas = nil
	} else {
		for i := 0; i < len(datas); i++ {

			parent := componentDataSubsConvertToModel(datas[i].Parent)
			componentDatas[i] = &model.ComponentData{
				Cname:              datas[i].Cname,
				Ctype:              datas[i].Ctype,
				Caption:            datas[i].Caption,
				Description:        datas[i].Description,
				ModifyDate:         datas[i].ModifyDate,
				HardConvType:       datas[i].HardConvType,
				Category:           datas[i].Category,
				Parent:             parent,
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
			}
		}
	}
	return componentDatas
}

func arrayConvert(elements []interface{}) []interface{} {
	var result []interface{}
	switch v := elements.(type) {
	case []string:
		result = append(result, v)
	case []int32:
		result = append(result, v)
	case int64:
		result = append(result, v)
	default:
		result = elements
	}
	return result
}
