package function

import (
	"encoding/json"
	"fmt"
	"time"
)

import (
	common2 "unicontract/src/common"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/model"
)

func init() {

}

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//++++++++++合约机【ENERGYTRADING】能源交易专用扩展方法++++++++++
//账户说明： 银行账户：购电者银行账户、运营者银行账户、售电者银行账户
//           交易账户：购电者交易账户、运营者交易账户、售电者交易账户
//           电表账户：电表充值账户
//充值过程：用户账户充值：购电者通过手动支付方式由银行账户转账到交易账户中(银行交易)
//              用户交易账户充入100元（创建资产）
//              托管资金账户充入100元（银行转账）
//          用户购电充值，将用户交易账户中的钱转移到运营商交易账户中（账户交易）
//              用户交易账户转账给托管交易账户 50元
//              用户电表余额被修改为 50元
//          电表耗电（消耗记录交易）:
//              电表交易账户转账给运营商交易账户
//          合约分账：
//              运营商交易账户转账给风、光、火电、国网等交易账户
//          银行结算：
//              托管资金账户银行转账给风、光、火、运营商、国网等银行账户
//能源交易涉及：人员身份链、电力能源链、能源交易链、交易票据链
//   人员身份链：电表完整信息、电表用户完整信息、发电厂完整信息、运营商完整信息存储
//   电力能源链：通过Elink实时采集电表数据，推送到消息队列中，并同时完成入链操作
//   能源交易链：电表用户账户充值，用户资产创建；
//               电表用户购电，用户资产转移到托管交易账户中；
//                   同时，修改用户电表余额，生成购电充值票据（电表用户付款给运营商账户票据）；
//               电表耗电消费，运营商交易账户将钱转账到 交易风、光、水、火电账户；
//                   同时，修改用户电表余额，生成支付单据（各发电方生成票据）；
//   交易票据链：生成购电充值票据（电表用户付款给运营商账户票据）；
//               生成支付单据（运营商付款给各发电方票据）；
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//人员身份链： 建立身份表
//    人员身份信息
//    电表身份信息
//    发电厂身份信息
//    运营商身份信息
//电力能源链：建立能源表
//    用户电表能源
//    风发电电表能源
//    光发电电表能源
//    火发电电表能源
//能源交易链：建立交易表
//    交易用户充值交易
//	运营商分账交易
//交易票据链：建立票据表
//    用户电表充值后，提供充电票据
//    运营商分账给风电，提供分账票据
//    运营商分账给光电，提供分账票据
//    运营商分账给火电，提供分账票据

//++++++++++++++++自动购电合约++++++++++++++++++++++++++++++++++++++++++++++++++++
//查询电表余额
//    访问电力能源链，读取用户电表余额
//Args: User_A  string  电表公钥
func FuncQueryAmmeterBalance(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	if len(args) == 0 {
		v_result.SetMessage("param is null!")
		return v_result, v_err
	}

	publickey, ok := args[0].(string)
	if !ok {
		v_result.SetMessage("args[0].(string) is error!")
		return v_result, v_err
	}

	money, v_err := rethinkdb.GetMoneyFromEnergy(publickey)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	//构建返回值
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(fmt.Sprintf("{\"money\":%f}", money))
	return v_result, v_err
}

//查询用户账户余额
//    查询能源交易链,获取用户交易账户余额
//Args: UserA  string  用户公钥
func FuncQueryAccountBalance(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	if len(args) == 0 {
		v_result.SetMessage("param is null!")
		return v_result, v_err
	}

	publickey, ok := args[0].(string)
	if !ok {
		v_result.SetMessage("args[0].(string) is error!")
		return v_result, v_err
	}

	money, v_err := rethinkdb.GetUserMoneyFromTransaction(publickey)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(fmt.Sprintf("{\"money\":%f}", money))
	return v_result, v_err
}

//短信提示账户：交易账户余额不足，自动购电操作无法完成；请及时充值50元到交易账户中
//Args: User_A  string 用户账户
//      50      int    充值额度
func FuncNoticeDeposit(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	if len(args) == 0 {
		v_result.SetMessage("param is null!")
		return v_result, v_err
	}

	publickey, ok := args[0].(string)
	if !ok {
		v_result.SetMessage("args[0].(string) is error!")
		return v_result, v_err
	}

	money, ok := args[1].(int)
	if !ok {
		v_result.SetMessage("args[1].(int) is error!")
		return v_result, v_err
	}

	var msgNotice model.DemoMsgNotice
	msgNotice.Id = common2.GenerateUUID()
	msgNotice.NoticePublicKey = publickey
	msgNotice.Timestamp = common2.GenTimestamp()
	msgNotice.Msg = fmt.Sprintf("请及时充值%d元到交易账户中，谢谢您的合作。", money)
	msgNotice.Type = 0

	slData, _ := json.Marshal(msgNotice)

	v_err = rethinkdb.InsertEnergyTradingDemoMsgNotice(string(slData))
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//电表自动购电50元（链上进行资产转移50给运营账户；同时访问电表接口，给电表充值50元）
//Args:  User_A   string  将用户账户中的钱50元，转到运营商账户
//       Ccount_D string
//       50       int     充值额度
func FuncAutoPurchasingElectricity(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	if len(args) == 0 {
		v_result.SetMessage("param is null!")
		return v_result, v_err
	}

	userPublicKey, ok := args[0].(string)
	if !ok {
		v_result.SetMessage("args[0].(string) is error!")
		return v_result, v_err
	}

	operatorPublicKey, ok := args[1].(string)
	if !ok {
		v_result.SetMessage("args[1].(string) is error!")
		return v_result, v_err
	}

	money, ok := args[2].(int)
	if !ok {
		v_result.SetMessage("args[2].(int) is error!")
		return v_result, v_err
	}

	// 用户->运营商
	bill1 := model.DemoBill{
		Id:        common2.GenerateUUID(),
		PublicKey: "",
		Timestamp: common.GenTimestamp(),
		Type:      1,
	}
	sldata, _ := json.Marshal(bill1)
	v_err = rethinkdb.InsertEnergyTradingDemoBill(string(sldata))
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	transaction1 := model.DemoTransaction{
		Id:            common2.GenerateUUID(),
		BillId:        bill1.Id,
		Timestamp:     common.GenTimestamp(),
		FromPublicKey: userPublicKey,
		ToPublicKey:   operatorPublicKey,
		Money:         float64(money),
		Type:          1,
	}
	sldata, _ = json.Marshal(transaction1)

	v_err = rethinkdb.InsertEnergyTradingDemoTransaction(string(sldata))
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	// 电表充值
	// TODO 修改电表余额

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//自动休眠1小时
//Args： SleepTime  int 单位是秒
func FuncAutoSleeping(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	if len(args) == 0 {
		v_result.SetMessage("param is null!")
		return v_result, v_err
	}

	sleeptime, ok := args[0].(int)
	if !ok {
		v_result.SetMessage("args[0].(string) is error!")
		return v_result, v_err
	}

	time.Sleep(time.Second * time.Duration(sleeptime))

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//++++++++++++++++耗电消耗 及 自动分账合约++++++++++++++++++++++++++++++++++++++++++
// 获取查询起始时间
// userKey string
func FuncGetStartEndTime(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	if len(args) == 0 {
		v_result.SetMessage("param is null!")
		return v_result, v_err
	}

	userPublicKey, ok := args[0].(string)
	if !ok {
		v_result.SetMessage("args[0].(string) is error!")
		return v_result, v_err
	}

	// 获得电表key
	meterKey, v_err := rethinkdb.GetMeterKeyByUserKey(userPublicKey)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	// 获得上次查询时间
	lastTime, v_err := rethinkdb.GetMeterQueryLastTime(meterKey)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	// 获得当前时间
	nowTime := common2.GenTimestamp()

	// 更新上次查询时间
	v_err = rethinkdb.UpdateMeterQueryLastTime(meterKey, nowTime)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(fmt.Sprintf("{\"start\":\"%s\",\"end\":\"%s\"}", lastTime, nowTime))
	return v_result, v_err
}

//A 根据电表账户从能源连上获取 上一时间点--—当前时间点的耗电量（分段时间点耗电量）、电表余额、当月截止当前总耗电量
//Args: ElecUser_A  string  电表交易用户
//      startTime   string
//      endTime     string
func FuncGetPowerConsumeParam(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	if len(args) == 0 {
		v_result.SetMessage("param is null!")
		return v_result, v_err
	}

	userPublicKey, ok := args[0].(string)
	if !ok {
		v_result.SetMessage("args[0].(string) is error!")
		return v_result, v_err
	}

	startTime, ok := args[1].(string)
	if !ok {
		v_result.SetMessage("args[1].(string) is error!")
		return v_result, v_err
	}

	endTime, ok := args[2].(string)
	if !ok {
		v_result.SetMessage("args[2].(string) is error!")
		return v_result, v_err
	}

	electricity, money, totalElectricity, v_err := rethinkdb.GetMeterInformation(userPublicKey, startTime, endTime)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(fmt.Sprintf("{\"electricity\":%f,\"money\":%f,\"totalElectricity\":%f}", electricity, money, totalElectricity))
	return v_result, v_err
}

//B. 消耗电量等于0，获取当前时间（作为下次循环评定的等待时间的起点）

//C. 消耗电量大于0时，获取对应的电价(峰谷平电价 和 阶梯的综合电价)
//获取电价信息（波峰平谷电价 & 阶梯电价）
func FuncGetPowerPrice(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	price, v_err := rethinkdb.GetPrice()
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(price)
	return v_result, v_err
}

//D. 根据用户耗电量对应出相应的电价，并计算用户消耗的电费、更新后的余额
//根据用户电表账户计算当前消耗的电量、将消耗对应电价计算消耗的金额、电表余额；
//Args: user_A              string  用户交易账户
//      electricity         string  当前耗电量
//      electricityTotal    string  当月总耗电量
//      startTime           string  采集的开始时间
//Return: consume_money   消耗金额
//        remain_money    电表余额
func FuncCalcConsumeAmountAndMoney(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//E. 打印分账票据，并记录到票据链上
//Args: user_B         string   运营商交易账户
//      other_users    string   合约分账的各用户
//      other_transfer string   各用户转账金额列表
func FuncTransferElecChargeToPlatform(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//F 修改电表余额
//Args：  user_A   string   elec_account
//        amount   int      电表余额
func FuncUpdateElecBalance(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	if len(args) == 0 {
		v_result.SetMessage("param is null!")
		return v_result, v_err
	}

	userPublicKey, ok := args[0].(string)
	if !ok {
		v_result.SetMessage("args[0].(string) is error!")
		return v_result, v_err
	}

	money, ok := args[1].(int)
	if !ok {
		v_result.SetMessage("args[1].(int)!")
		return v_result, v_err
	}

	// TODO 修改电表余额
	_ = userPublicKey
	_ = money

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//I. 获取各发电厂该时间段发电总电量，并计算发电比例（资源链中统计发电总量、各电厂发电总量，计算出各电厂发电比例； 统计用户全天用电总量，各时间段用电总量，计算出用电比例； ）
//Args: other_user  string   发电方交易账户字符串集
//      begin_time  string   统计起始时间串
//      end_time    string   统计结束时间串
//Return : split_percent string  各合约用户分账比例
func FuncCalcAndSplitRatio(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//J. 根据各发电厂全天比例、进行合约分账
//Args: user_B       string   运营商交易账户
//      other_users  string   发电厂及合约分账的各交易账户
//      split_percent string  合约分账各方分账的比例
func FuncAutoSplitAccount(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//K.获取当前时间（作为下次休眠判断的起始时间）
