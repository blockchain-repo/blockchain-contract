package quartz

import (
	"errors"
	"time"

	"unicontract/src/chain"
	"unicontract/src/common"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/pipelines"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/model"
)

const (
	_DBName                  = "Unicontract"
	_TableContractOutputs    = "ContractOutputs"
	_TableSendFailingRecords = "SendFailingRecords"
)

func init() {
	go sendFailingDataTimer()
}

func sendFailingDataTimer() {
	timer := time.Tick(1 * time.Hour)
	for now := range timer {
		uniledgerlog.Info(now)
		idList, _ := rethinkdb.GetAllRecords(_DBName, _TableSendFailingRecords)
		for _, value := range idList {
			str := rethinkdb.Get(_DBName, _TableContractOutputs, value)

			result, err := chain.CreateContractTx(common.Serialize(str))
			if err != nil {
				uniledgerlog.Error(err.Error())
				continue
			}

			if result.Code != 200 {
				uniledgerlog.Error(errors.New("request send failed"))
				coModel := model.ContractOutput{}
				if err != nil {
					uniledgerlog.Error(err.Error())
					continue
				}
				pipelines.SaveOutputErrorData(_TableSendFailingRecords, coModel)
			}
		}
	}
}
