package function

import (
	"github.com/astaxie/beego/logs"
	"strconv"
	"testing"
	"unicontract/src/config"
)

func init() {
	config.Init()
}
func TestBoo(t *testing.T) {
	var b bool = true
	var s string = strconv.FormatBool(b)
	logs.Info(b)
	logs.Info(s)
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

	ou, _ := FuncCreateAsset(ownerBefore, recipients, contractStr, contractHashId, contractId, taskId, taskIndex)

	logs.Info(ou.Data)

	//var ownerBefore string = args[0].(string)
	//var recipients [][2]interface{} = [][2]interface{}{}
	////executer provide
	//var contractStr string = args[2].(string)
	//var contractHashId string = args[3].(string)
	//var contractId string = args[4].(string)
	//var taskId string = args[5].(string)
	//var taskIndex int = args[6].(int)
	////var mainPubkey string = args[7].(string)
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

	out, _ := FuncTransferAsset(ownerBefore, recipients, contractStr, contractHashId, contractId, taskId, taskIndex, mainPubkey)
	logs.Info(out.Data)
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
		logs.Info("err-", err)
	}
	outputStr := out.Data.(string)
	//var contractOutPut string = args[0].(string)
	var taskStatus string = "down"
	out, _ = FuncTransferAssetComplete(outputStr, taskStatus)
	logs.Info(out.Data)
}

func TestFuncInterim(t *testing.T) {

}

func TestFuncInterimComplete(t *testing.T) {

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

	logs.Info(out)
	logs.Info(err)

}

func TestFuncGetContracOutputtById(t *testing.T) {
	conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	res, _ := FuncGetContracOutputtById(conId)
	logs.Info(res.Data)
}

func TestFuncIsConPutInUnichian(t *testing.T) {

}
