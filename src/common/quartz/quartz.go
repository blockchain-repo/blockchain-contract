package quartz

import (
	"errors"
	"time"

	"unicontract/src/chain"
	"unicontract/src/common"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/pipelines"

	"github.com/astaxie/beego/logs"
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
		logs.Info(now)
		idList, _ := rethinkdb.GetAllRecords(_DBName, _TableSendFailingRecords)
		for _, value := range idList {
			str := rethinkdb.Get(_DBName, _TableContractOutputs, value)

			result, err := chain.CreateContractTx(common.Serialize(str))
			if err != nil {
				logs.Error(err.Error())
				continue
			}

			if result.Code != 200 {
				logs.Error(errors.New("request send failed"))
				coModel := model.ContractOutput{}
				if err != nil {
					logs.Error(err.Error())
					continue
				}
				pipelines.SaveOutputErrorData(_TableSendFailingRecords, coModel)
			}
		}
	}
}
