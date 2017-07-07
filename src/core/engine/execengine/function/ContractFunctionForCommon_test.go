package function

import (
	"unicontract/src/common/uniledgerlog"
	"strconv"
	"strings"
	"testing"
	"unicontract/src/config"
)

func init() {
	config.Init()
}
func TestBoo(t *testing.T) {
	var b bool = true
	var s string = strconv.FormatBool(b)
	uniledgerlog.Info(b)
	uniledgerlog.Info(s)
}

func TestSplit(t *testing.T) {
	s := strings.Split("a#100&b#200", "&")
	uniledgerlog.Info(s, len(s))
	var re [][2]interface{} = [][2]interface{}{}
	for i := 0; i < len(s); i++ {
		ss := strings.Split(s[i], "#")
		ownAfter := ss[0]
		amount, _ := strconv.ParseFloat(ss[1], 64)
		re = append(re, [2]interface{}{ownAfter, amount})
		uniledgerlog.Info(ss)
	}
	uniledgerlog.Info(re)
}

func TestFuncCreateAsset(t *testing.T) {

	conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	res, _ := FuncGetContracOutputtById(conId)

	var ownerBefore string = ""
	var recipients [][2]interface{} = [][2]interface{}{[2]interface{}{
		"EcWbt741xS8ytvKWEqCPtDu29sgJ1iHubHyoVvuAgc8W", 1000}}
	//executer provide
	var contractStr string = res.Data.(string)

	var contractHashId string = ""
	var contractId string = "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	var taskId string = "task_id"
	var taskIndex int = 1

	ou, err := FuncCreateAsset(ownerBefore, recipients, contractStr, contractHashId, contractId, taskId, taskIndex)
	uniledgerlog.Info(err)
	uniledgerlog.Info(ou.Data)
}

func TestFuncTransferAsset(t *testing.T) {
	//user provide
	var ownerBefore string = "EcWbt741xS8ytvKWEqCPtDu29sgJ1iHubHyoVvuAgc8W"
	var recipients [][2]interface{} = [][2]interface{}{
		[2]interface{}{"5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL", 100},
	}
	//executer provide
	conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	res, _ := FuncGetContracOutputtById(conId)
	var contractStr string = res.Data.(string)
	var contractHashId string = ""
	var contractId string = "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	var taskId string = "task_id"
	var taskIndex int = 1
	var mainPubkey string = "7mQXR8NY9M1Uj86VM4CHnacY8fpLPudfmn5DaJGgXDi9"
	//var metadataStr string = ""

	out, err := FuncTransferAsset(ownerBefore, recipients, contractStr, contractHashId, contractId, taskId, taskIndex, mainPubkey)
	uniledgerlog.Info(err)
	uniledgerlog.Info(out.Data)
}

func TestFuncTransferAssetComplete(t *testing.T) {
	//user provide
	var ownerBefore string = "EcWbt741xS8ytvKWEqCPtDu29sgJ1iHubHyoVvuAgc8W"
	var recipients [][2]interface{} = [][2]interface{}{
		[2]interface{}{"5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL", 100},
	}
	//executer provide
	conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	res, _ := FuncGetContracOutputtById(conId)
	var contractStr string = res.Data.(string)
	var contractHashId string = ""
	var contractId string = "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	var taskId string = "task_id"
	var taskIndex int = 1
	var mainPubkey string = "7mQXR8NY9M1Uj86VM4CHnacY8fpLPudfmn5DaJGgXDi9"
	//var metadataStr string = ""

	out, err := FuncTransferAsset(ownerBefore, recipients, contractStr, contractHashId, contractId, taskId, taskIndex, mainPubkey)
	if err != nil {
		uniledgerlog.Info("err-", err)
		return
	}
	outputStr := out.Data.(string)
	//var contractOutPut string = args[0].(string)
	var taskStatus string = "down"
	out, err = FuncTransferAssetComplete(outputStr, taskStatus)
	uniledgerlog.Info(err)
	uniledgerlog.Info(out.Data)
}

func TestFuncInterim(t *testing.T) {
	conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	res, _ := FuncGetContracOutputtById(conId)
	var contractStr string = res.Data.(string)
	var contractHashId string = ""
	var contractId string = "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	var taskId string = "task_id"
	var taskIndex int = 1
	out, err := FuncInterim(contractStr, contractHashId, contractId, taskId, taskIndex)
	uniledgerlog.Info(err)
	uniledgerlog.Info(out)
}

func TestFuncInterimComplete(t *testing.T) {
	conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	res, _ := FuncGetContracOutputtById(conId)
	var contractStr string = res.Data.(string)
	var contractHashId string = ""
	var contractId string = "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	var taskId string = "task_id"
	var taskIndex int = 1
	out, err := FuncInterim(contractStr, contractHashId, contractId, taskId, taskIndex)
	if err != nil {
		uniledgerlog.Info("err-", err)
		return
	}
	outputStr := out.Data.(string)
	//var contractOutPut string = args[0].(string)
	var taskStatus string = "down"
	out, err = FuncTransferAssetComplete(outputStr, taskStatus)
	uniledgerlog.Info(err)
	uniledgerlog.Info(out.Data)
}

func TestFuncUnfreezeAsset(t *testing.T) {
	var ownerBefore string = "EcWbt741xS8ytvKWEqCPtDu29sgJ1iHubHyoVvuAgc8W"
	var recipients [][2]interface{} = [][2]interface{}{}
	//executer provide
	conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	res, _ := FuncGetContracOutputtById(conId)
	var contractStr string = res.Data.(string)
	var contractHashId string = ""
	var contractId string = "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	var taskId string = "task_id"
	var taskIndex int = 1
	//var mainPubkey string = args[7].(string)
	//var metadataStr string = ""
	out, err := FuncUnfreezeAsset(ownerBefore, recipients, contractStr, contractHashId, contractId, taskId, taskIndex)

	uniledgerlog.Info(out)
	uniledgerlog.Info(err)

}

func TestFuncGetContracOutputtById(t *testing.T) {
	conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	res, _ := FuncGetContracOutputtById(conId)
	uniledgerlog.Info(res.Data)
}

func TestFuncIsConPutInUnichian(t *testing.T) {
	conhashId := "63841426ea1c501745d56ce47a4e7b93bf85841d54f2c77102ce488ac0ce8b51"
	res, _ := FuncIsConPutInUnichian(conhashId)
	uniledgerlog.Info(res.Code)
	uniledgerlog.Info(res.Data)
}

func TestFuncGetNowDateTimestamp(t *testing.T) {
	uniledgerlog.Info(FuncGetNowDateTimestamp())
}
