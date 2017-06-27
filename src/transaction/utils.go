package transaction

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"unicontract/src/core/model"
)

func GenModelByExecStr(m string, r string, c string) (model.Metadata, model.Relation, model.ContractModel, error) {
	metadata, err := GenMetadataByExecStr(m)
	relation, err := GenRelationByExecStr(r)
	contract, err := GenContractByExecStr(c)
	return metadata, relation, contract, err
}

func GenMetadataByExecStr(m string) (met model.Metadata, err error) {
	metadata := model.Metadata{}
	err = json.Unmarshal([]byte(m), &metadata)
	return metadata, err
}

func GenRelationByExecStr(r string) (rela model.Relation, err error) {
	relation := model.Relation{}
	err = json.Unmarshal([]byte(r), &relation)
	return relation, err
}

func GenContractByExecStr(c string) (con model.ContractModel, err error) {
	logs.Info("===c: " + c)
	//contract := model.ContractModel{}
	var contract model.ContractModel
	err = json.Unmarshal([]byte(c), &contract)
	logs.Info("===contract: ", contract)
	return contract, err
}
