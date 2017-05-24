package function

import (
	"github.com/astaxie/beego/logs"
	"strconv"
	"testing"
)

func TestBoo(t *testing.T) {
	var b bool = true
	var s string = strconv.FormatBool(b)
	logs.Info(b)
	logs.Info(s)
}

func TestFuncCreateAsset(t *testing.T) {

	//conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	//res, _ := FuncGetContracOutputtById(conId)
	//
	//var ownerBefore string = ""
	//var recipients [][2]interface{} = [][2]interface{}{[2]interface{}{"EcWbt741xS8ytvKWEqCPtDu29sgJ1iHubHyoVvuAgc8W", 200}}
	////executer provide
	//var contractStr string = res.Data
	//
	//var contractHashId string = ""
	//var contractId string = "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	//var taskId string = "task_id"
	//var taskIndex int = 1

}

func TestFuncTransferAsset(t *testing.T) {

}

func TestFuncTransferAssetComplete(t *testing.T) {

}

func TestFuncUnfreezeAsset2(t *testing.T) {

}

func TestFuncGetContracOutputtById(t *testing.T) {
	conId := "feca0672-4ad7-4d9a-ad57-83d48db2269b"
	res, _ := FuncGetContracOutputtById(conId)
	logs.Info(res.Data)
}

func TestFuncIsConPutInUnichian(t *testing.T) {

}
