package pipelines

import (
	"io"
	"encoding/json"
	"log"
	"time"

	"unicontract/src/core/model"
	"unicontract/src/core/db/rethinkdb"

	"github.com/astaxie/beego/logs"
)

const (
	MaxSizeTX                    = 16 * 1024
	_TableNameSendFailingRecords = "SendFailingRecords"
)

//bind函数主要是用来为pipe函数整合用的，通过将闭包将函数签名变成pipe所需的样子
//返回一个函数闭包，将一个函数字面量app和字符串slice 传入其中
func Bind(app func(in io.Reader, out io.Writer, args []string), args []string) func(in io.Reader, out io.Writer) {
	return func(in io.Reader, out io.Writer) {
		app(in, out, args)
	}
}

//将两个函数插入到管道的中间，调用者只需调用pipe返回的函数字面量，并传入管道的首尾两端，即可实现管道
//返回一个新的函数闭包
func Pipe(apps ...func(in io.Reader, out io.Writer)) func(in io.Reader, out io.Writer) {
	if len(apps) == 0 {
		return nil
	}
	app := apps[0]
	for i := 1; i < len(apps); i++ {
		app1, app2 := app, apps[i]
		app = func(in io.Reader, out io.Writer) {
			pr, pw := io.Pipe()
			defer pw.Close()
			go func() {
				defer pr.Close()
				app2(pr, out)
			}()
			app1(in, pw)
		}
	}
	return app
}

func SaveOutputErrorData(tableName string, t []byte) bool {
	coModel := model.ContractOutput{}
	err := json.Unmarshal(t, &coModel)
	if err != nil {
		log.Fatalf(err.Error())
	}
	dataId := coModel.Id

	res := rethinkdb.Get(_DBName, tableName, dataId)
	//TODO test nil data
	if res != nil {
		return true
	}

	//insert
	failTime := time.Now().String()
	dataJson := `{"id":"`+dataId+`","tableName":"`+tableName+`","failTime":"`+failTime+`","sendTime":"","numSend":"1","status":"unsend"}`
	logs.Info(dataJson)
	var result = rethinkdb.Insert(_DBName, tableName, dataJson)
	logs.Info(result)

	return true
}

//func init() {
//	//TODO log
//	go startContractVote()
//	go startContractElection()
//	go starttxElection()
//}
