/*************************************************************************
  > File Name: headNodeMonitor.go
  > Module:
  > Function: 对头节点进行监控，如果发现头节点出现问题，及时更换对应contract的头节点
  > Author: wangyp
  > Company:
  > Department:
  > Mail: wangyepeng87@163.com
  > Created Time: 2017.06.01
 ************************************************************************/
package headNodeMonitor

import (
	"encoding/json"
	"fmt"
	"time"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
func _HeadNodeMonitor() {
	for {
		ticker := time.NewTicker(time.Hour * (time.Duration)(headNodeMonitorConf["scan_time"].(int)))
		beegoLog.Debug("wait for scan contract table to monitor head node ...")
		select {
		case <-ticker.C:
			beegoLog.Debug("query no consensus contract")
			timePoint := time.Now().
				Add(-time.Hour*(time.Duration)(headNodeMonitorConf["interval_time"].(int))).
				UnixNano() / 1000000
			strNoConsensusContract, err :=
				rethinkdb.GetNoConsensusContracts(fmt.Sprintf("%d", timePoint), 0)
			if err != nil {
				beegoLog.Error(err)
				continue
			}

			if len(strNoConsensusContract) == 0 {
				beegoLog.Debug("no consensus contract")
				continue
			}

			var slContracts []model.ContractModel
			json.Unmarshal([]byte(strNoConsensusContract), &slContracts)

			beegoLog.Debug("delete old contract and insert new contract")
			for index, value := range slContracts {
				// 生成新的头节点
				index_new := _GenerateAnotherHeadNodeKey(value.ContractHead.GetMainPubkey())
				slContracts[index].ContractHead.MainPubkey = gslPublicKeys[index_new]

				// 删除老的contract对应的vote，如果删除出现问题，无关紧要，无需continue
				err := _DeleteOldVotes(value.Id)
				if err != nil {
					beegoLog.Error(err)
				}

				// 删除老的contract
				if !rethinkdb.DeleteContract(value.Id) {
					beegoLog.Error(err)
					continue
				}

				// 插入新的contract
				slData, err = json.Marshal(slContracts[index])
				if err != nil {
					beegoLog.Error(err)
					continue
				}
				if !rethinkdb.InsertContract(string(slData)) {
					beegoLog.Error(err)
				}
			}
		}
	}
	gwgHeadNodeMonitor.Done()
}

//---------------------------------------------------------------------------
func _GenerateAnotherHeadNodeKey(key string) int {
	var i int
	for index, value := range gslPublicKeys {
		if value == key {
			i = (index + 1) % gnPublicKeysNum
			break
		}
	}
	return i
}

//---------------------------------------------------------------------------
func _DeleteOldVotes(strContractId string) error {
	strVotes, err := rethinkdb.GetVotesByContractId(strContractId)
	if err != nil {
		return err
	}

	var slVote []model.Vote
	err = json.Unmarshal([]byte(strVotes), &slVote)
	if err != nil {
		return err
	}

	slID := make([]interface{}, 0)
	for index, _ := range slVote {
		slID = append(slID, slVote[index].Id)
	}

	deleteNum, err := rethinkdb.DeleteVotes(slID)
	if err != nil {
		return err
	}

	beegoLog.Info("delete [ %d ] votes for old contract", deleteNum)

	return nil
}

//---------------------------------------------------------------------------
