package core

import (
	"time"
	"math/rand"

	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
	"unicontract/src/config"
)

func WriteContract(contract model.ContractModel) {
	rand.Seed(time.Now().UnixNano())
	pubs := config.GetAllPublicKey()
	contract.MainPubkey = pubs[rand.Intn(len(pubs))]
	str := contract.ToString()
	r.Insert("Unicontract","Contracts",str)
}