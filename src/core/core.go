package core

import (
	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

func WriteContract(contract model.ContractModel) {
	//TODO keyrings
	contract.MainPubkey = "EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet"
	str := contract.ToString()
	r.Insert("Unicontract","Contracts",str)
}