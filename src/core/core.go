package core

import (
	"math/rand"
	"time"

	"unicontract/src/common"
	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

func WriteContract(contract model.ContractModel) bool {
	rand.Seed(time.Now().UnixNano())
	pubs := config.GetAllPublicKey()

	contract.ContractHead.MainPubkey = pubs[rand.Intn(len(pubs))]
	contract.ContractHead.OperateTime = common.GenTimestamp()
	ok := r.InsertContract(contract.ToString())
	return ok
}
