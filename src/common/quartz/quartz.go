package quartz

import (
	"time"
	"errors"
	"encoding/json"

	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/common"
	"unicontract/src/chain"
	"unicontract/src/pipelines"

	"github.com/astaxie/beego/logs"
)

const(
	_DBName = "Unicontract"
	_TableContractOutputs = "ContractOutputs"
	_TableSendFailingRecords = "SendFailingRecords"
)

func init() {
	go sendFailingDataTimer()
}

func sendFailingDataTimer()  {
	timer := time.Tick(1 * time.Hour)
	for now := range timer {
		logs.Info(now)
		idList,_ := rethinkdb.GetAllRecords(_DBName, _TableSendFailingRecords)
		for _,value := range idList {
			str := rethinkdb.Get(_DBName,_TableContractOutputs,value)

			result,err:= chain.CreateContractTx(common.Serialize(str))
			if err != nil{
				logs.Error(err.Error())
			}

			if result.Code != 200 {
				logs.Error(errors.New("request send failed"))
				strByte,_ := json.Marshal(str)
				pipelines.SaveOutputErrorData(_TableSendFailingRecords,strByte)
			}
		}
	}
}
