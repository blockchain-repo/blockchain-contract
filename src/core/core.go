package core

import (
	"math/rand"
	"time"

	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

func WriteContract(contract model.ContractModel) bool {
	contract_write_time := monitor.Monitor.NewTiming()
	rand.Seed(time.Now().UnixNano())
	pubs := config.GetAllPublicKey()
	contract.ContractHead.MainPubkey = pubs[rand.Intn(len(pubs))]
	contract.ContractHead.AssignTime = common.GenTimestamp()
	ok := r.InsertContract(contract.ToString())
	contract_write_time.Send("contract_write")
	return ok
}
