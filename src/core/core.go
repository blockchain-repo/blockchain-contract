package core

import (
	"math/rand"
	"time"

	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
	"unicontract/src/common"
)

func WriteContract(contract model.ContractModel) bool {
	rand.Seed(time.Now().UnixNano())
	pubs := config.GetAllPublicKey()

	contract.ContractHead.MainPubkey = pubs[rand.Intn(len(pubs))]
	contract.ContractHead.Timestamp = common.GenTimestamp()
	ok := r.InsertContract(contract.ToString())
	return ok
}
