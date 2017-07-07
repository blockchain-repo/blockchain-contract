package transaction

import (
	"encoding/json"
	"unicontract/src/common/uniledgerlog"
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
	uniledgerlog.Info("===c1: " + c)
	//c = strings.Trim(c, "\"")
	uniledgerlog.Info("===c2: " + c)
	contract := model.ContractModel{}
	err = json.Unmarshal([]byte(c), &contract)
	uniledgerlog.Info("===contract: ", []byte(c))
	return contract, err
}
