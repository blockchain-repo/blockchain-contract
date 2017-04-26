package core

import (
	"math/rand"
	"time"

	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

func WriteContract(contract model.ContractModel) {
	rand.Seed(time.Now().UnixNano())
	pubs := config.GetAllPublicKey()

	contract.ContractHead.MainPubkey = pubs[rand.Intn(len(pubs))]
	r.InsertContract(contract.ToString())
}
