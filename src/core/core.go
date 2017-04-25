package core

import (
	"time"
	"math/rand"

	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
	"unicontract/src/config"
	"unicontract/src/core/protos"
)

func WriteContract(contract model.ContractModel) {
	rand.Seed(time.Now().UnixNano())
	pubs := config.GetAllPublicKey()
	contractHead := &protos.ContractHead{}

	contractHead.MainPubkey = pubs[rand.Intn(len(pubs))]
	contract.ContractHead = contractHead
	str := contract.ToString()
	r.Insert("Unicontract","Contracts",str)
}