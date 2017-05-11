package transaction

import "unicontract/src/core/model"

func GenModelByExecStr(m string,r string,c string)(model.Metadata,model.Relation,model.ContractModel){
	metadata:=GenMetadataByExecStr(m)
	relation:=GenRelationByExecStr(r)
	contract:=GenContractByExecStr(c)

	return 	metadata,relation,contract
}


func GenMetadataByExecStr(m string) model.Metadata {
	metadata := model.Metadata{}
	//TODO

	return metadata
}

func GenRelationByExecStr(r string) model.Relation {
	relation := model.Relation{}
	//TODO

	return relation
}

func GenContractByExecStr(c string) model.ContractModel {
	contract := model.ContractModel{}
	//TODO

	return contract
}
