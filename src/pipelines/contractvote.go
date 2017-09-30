package pipelines

import (
	"encoding/json"
	"sync"

	"unicontract/src/common"
	"unicontract/src/common/monitor"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	r "unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"

	"github.com/astaxie/beego"
)

func cvValidateContract(arg interface{}) interface{} {
	bs, err := json.Marshal(arg)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return nil
	}
	mod := model.ContractModel{}
	err = json.Unmarshal(bs, &mod)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return nil
	}
	v := model.Vote{}
	contract_validate_time := monitor.Monitor.NewTiming()
	if mod.Validate() {
		//vote true
		v.VoteBody.IsValid = true
	} else {
		//vote flase
		v.VoteBody.IsValid = false
	}
	contract_validate_time.Send("contract_validate")
	v.VoteBody.VoteFor = mod.Id
	//uniledgerlog.Debug("-------cvValidateContract:", common.Serialize(v))
	return v
}

func cvVote(arg interface{}) interface{} {
	v := arg.(model.Vote)

	v.NodePubkey = config.Config.Keypair.PublicKey
	v.VoteBody.Timestamp = common.GenTimestamp()
	v.VoteBody.VoteType = "Contract"
	v.Id = v.GenerateId()
	v.Signature = v.SignVote()
	//uniledgerlog.Debug("-------cvVote:", common.Serialize(v))
	return v

}

func cvWriteVote(arg interface{}) interface{} {
	v := arg.(model.Vote)
	vote_write_time := monitor.Monitor.NewTiming()
	res := r.Insert("Unicontract", "Votes", v.ToString())
	uniledgerlog.Debug("-------cvWriteVote:", common.Serialize(res))
	vote_write_time.Send("vote_write")
	return v
}

func getcvChangefeed() *ChangeFeed {
	change := &ChangeFeed{
		db:        r.DBNAME,
		table:     r.TABLE_CONTRACTS,
		operation: INSERT,
		//operation: INSERT | UPDATE,
	}
	go change.runForever()
	return change
}

func createcvPip() (cvPip Pipeline) {
	cvNodeSlice := make([]*Node, 0)
	NodeGoroutineNum, err := beego.AppConfig.Int("PipelineNodeGoroutineNum")
	if err != nil {
		uniledgerlog.Error(err)
		NodeGoroutineNum = 1
	}
	cvNodeSlice = append(cvNodeSlice, &Node{target: cvValidateContract, routineNum: NodeGoroutineNum, name: "cvValidateContract"})
	cvNodeSlice = append(cvNodeSlice, &Node{target: cvVote, routineNum: NodeGoroutineNum, name: "cvVote"})
	cvNodeSlice = append(cvNodeSlice, &Node{target: cvWriteVote, routineNum: NodeGoroutineNum, name: "cvWriteVote"})
	cvPip = Pipeline{
		nodes: cvNodeSlice,
	}
	return cvPip
}

func startContractVote() {
	cvPip := createcvPip()
	changefeed := getcvChangefeed()
	cvPip.setup(&changefeed.node)
	cvPip.start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}
