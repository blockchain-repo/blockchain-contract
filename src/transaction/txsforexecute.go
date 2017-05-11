package transaction

import (
	"github.com/astaxie/beego/logs"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/common"
)

func ExecuteCreate(tx_signers []string, recipients [][2]interface{}, metadataStr string,
	relationStr string, contractStr string) {
	asset := GetAsset(tx_signers[0])
	metadata, relation, contract := GenModelByExecStr(metadataStr, relationStr, contractStr)

	output, _ := Create(tx_signers, recipients, &metadata, asset, relation, contract)
	output = NodeSign(output)
	b := rethinkdb.InsertContractOutput(common.StructSerialize(output))
	logs.Info(b)
	//TODO return
}

func ExecuteTransfer(operation string, ownerbefore string, recipients [][2]interface{},
	metadataStr string, relationStr string, contractStr string) {
	asset := GetAsset(ownerbefore)
	metadata, relation, contract := GenModelByExecStr(metadataStr, relationStr, contractStr)

	output, _ := Transfer(operation, ownerbefore, recipients, &metadata, asset, relation, contract)
	output = NodeSign(output)
	b := rethinkdb.InsertContractOutput(common.StructSerialize(output))
	logs.Info(b)
	//TODO return
}
