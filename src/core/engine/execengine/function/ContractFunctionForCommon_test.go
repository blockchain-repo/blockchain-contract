package function

import (
	"github.com/astaxie/beego/logs"
	"strconv"
	"testing"
)

func TestTransferAsset(t *testing.T) {
	/*
	   var ownerBefore string = args[0].(string)
	   	var recipients [][2]interface{} = [][2]interface{}{
	   		[2]interface{}{"5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL", 100},
	   	}
	   	//executer provide
	   	var contractStr string = args[2].(string)
	   	var contractHashId string = args[3].(string)
	   	var contractId string = args[4].(string)
	   	var taskId string = args[5].(string)
	   	var taskIndex int = args[6].(int)
	   	var mainPubkey string = args[7].(string)
	*/
	FuncTransferAsset()
}

func TestBoo(t *testing.T) {
	var b bool = true
	var s string = strconv.FormatBool(b)
	logs.Info(b)
	logs.Info(s)
}
