package function

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

import (
	common2 "unicontract/src/common"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/model"
)

import (
	"github.com/astaxie/beego/logs"
	"math"
)

type PowerPlants struct {
	Key   string
	Type_ int
}

type MeterInfor struct {
	Electricity      float64 // 当前用量
	TotalElectricity float64 // 当月总量
	Key              string
}

// 电表余额
var mapMeterRemainMoney map[string]float64

// 电表key
var slMeterKey []MeterInfor

// 发电厂key
var slPowerPlantsKey []PowerPlants

// 个人key
var slPersonKey []string

// 运营商key
var slOperatorKey []string

// 起始时间
var strStartTime string
var strEndTime string

// 消耗电量
var fElectricity float64
var fTotalElectricity float64

var fMoney float64

var mapEnergy map[string]float64

func __init() {
	mapMeterRemainMoney = make(map[string]float64)
	mapEnergy = make(map[string]float64)

	// 获得个人、运营商、电表、发电厂的key
	// 个人
	keys, err := rethinkdb.GetRolePublicKey(0)
	if err != nil {
		logs.Error(err)
		// 数据库中没有个人角色，就插入一条
		strPublicKey, _ := common2.GenerateKeyPair()
		person1 := model.DemoRole{
			Id:          common2.GenerateUUID(),
			Name:        "个人",
			PublicKey:   strPublicKey,
			Infermation: "",
			Type:        0,
		}
		sldata, _ := json.Marshal(person1)
		rethinkdb.InsertEnergyTradingDemoRole(string(sldata))
		keys = append(keys, strPublicKey)
	}
	slPersonKey = append(slPersonKey, keys...)

	keys = nil

	// 电表
	keys, err = rethinkdb.GetRolePublicKey(1)
	if err != nil {
		logs.Error(err)
		// 数据库中没有电表角色，就插入一条
		strPublicKey, _ := common2.GenerateKeyPair()
		mapInformation := make(map[string]string)
		mapInformation["ownerPublicKey"] = slPersonKey[0]
		slInformation, _ := json.Marshal(mapInformation)
		electricityMeter1 := model.DemoRole{
			Id:            common2.GenerateUUID(),
			Name:          "个人电表",
			PublicKey:     strPublicKey,
			Infermation:   string(slInformation),
			LastTimestamp: common2.GenTimestamp(),
			Type:          1,
		}
		sldata, _ := json.Marshal(electricityMeter1)
		rethinkdb.InsertEnergyTradingDemoRole(string(sldata))
		keys = append(keys, strPublicKey)
	}
	for _, value := range keys {
		MeterInfor_ := MeterInfor{
			Key:              value,
			Electricity:      500,
			TotalElectricity: 245,
		}
		slMeterKey = append(slMeterKey, MeterInfor_)
	}

	keys = nil

	// 运营商
	keys, err = rethinkdb.GetRolePublicKey(2)
	if err != nil {
		logs.Error(err)
		// 数据库中没有运营商角色，就插入一条
		strPublicKey, _ := common2.GenerateKeyPair()
		operator1 := model.DemoRole{
			Id:          common2.GenerateUUID(),
			Name:        "运营商",
			PublicKey:   strPublicKey,
			Infermation: "",
			Type:        2,
		}
		sldata, _ := json.Marshal(operator1)
		rethinkdb.InsertEnergyTradingDemoRole(string(sldata))
		keys = append(keys, strPublicKey)
	}
	slOperatorKey = append(slOperatorKey, keys...)

	keys = nil

	// 发电厂
	mapBalabala := make(map[int]string)
	mapBalabala[3] = "风电"
	mapBalabala[4] = "光电"
	mapBalabala[5] = "火电"
	mapBalabala[6] = "国网"
	type_ := []int{3, 4, 5, 6}
	for _, v := range type_ {
		keys, err := rethinkdb.GetRolePublicKey(v)
		if err != nil {
			logs.Error(err)
			// 数据库中没有发电厂角色，就插入一条
			strPublicKey, _ := common2.GenerateKeyPair()
			R := model.DemoRole{
				Id:          common2.GenerateUUID(),
				Name:        mapBalabala[v],
				PublicKey:   strPublicKey,
				Infermation: "",
				Type:        v,
			}
			sldata, _ := json.Marshal(R)
			rethinkdb.InsertEnergyTradingDemoRole(string(sldata))
			keys = append(keys, strPublicKey)
		}

		for _, value := range keys {
			PowerPlants_ := PowerPlants{
				Key:   value,
				Type_: v,
			}
			slPowerPlantsKey = append(slPowerPlantsKey, PowerPlants_)
		}
	}

	// 获得电表余额
	money, _ := rethinkdb.GetMoneyFromEnergy(slPersonKey[0])
	mapMeterRemainMoney[slPersonKey[0]] = money

	// 如果电价表不存在数据，则插入数据
	_, err = rethinkdb.GetPrice()
	if err != nil {
		price1 := model.DemoPrice{
			Id:          common2.GenerateUUID(),
			Level:       1,
			Low:         1,
			High:        240,
			One:         1,
			Two:         2,
			Three:       3,
			Description: "第一阶梯",
		}
		sldata, _ := json.Marshal(price1)
		rethinkdb.InsertEnergyTradingDemoPrice(string(sldata))

		price2 := model.DemoPrice{
			Id:          common2.GenerateUUID(),
			Level:       2,
			Low:         241,
			High:        400,
			One:         2,
			Two:         3,
			Three:       4,
			Description: "第二阶梯",
		}
		sldata, _ = json.Marshal(price2)
		rethinkdb.InsertEnergyTradingDemoPrice(string(sldata))

		price3 := model.DemoPrice{
			Id:          common2.GenerateUUID(),
			Level:       3,
			Low:         401,
			High:        math.MaxFloat64,
			One:         3,
			Two:         4,
			Three:       5,
			Description: "第三阶梯",
		}
		sldata, _ = json.Marshal(price3)
		rethinkdb.InsertEnergyTradingDemoPrice(string(sldata))
	}
}

// 1.模拟ELINK采集电表数据功能；2.模拟采集发电厂发电数据功能
func Simulate() {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(30 * time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for _ = range ticker.C {
			// 模拟采集电表数据
			logs.Info("模拟采集电表数据")
			for index, v := range slMeterKey {
				slMeterKey[index].Electricity += 1
				slMeterKey[index].TotalElectricity += 1
				electricityMeter1 := model.DemoEnergy{
					Id:               common2.GenerateUUID(),
					PublicKey:        v.Key,
					Timestamp:        common2.GenTimestamp(),
					Electricity:      slMeterKey[index].Electricity,
					TotalElectricity: slMeterKey[index].TotalElectricity,
					Money:            mapMeterRemainMoney[v.Key],
					Type:             0,
				}
				sldata, _ := json.Marshal(electricityMeter1)

				err := rethinkdb.InsertEnergyTradingDemoEnergy(string(sldata))
				if err != nil {
					logs.Error(err)
				}
			}

			// 模拟采集发电厂数据
			logs.Info("模拟采集发电厂数据")
			for _, v := range slPowerPlantsKey {
				electricityPowerPlant1 := model.DemoEnergy{
					Id:          common2.GenerateUUID(),
					PublicKey:   v.Key,
					Timestamp:   common2.GenTimestamp(),
					Electricity: float64(rand.Intn(1000)),
					Type:        v.Type_,
				}
				sldata, _ := json.Marshal(electricityPowerPlant1)

				err := rethinkdb.InsertEnergyTradingDemoEnergy(string(sldata))
				if err != nil {
					logs.Error(err)
				}
			}
		}
	}()
	wg.Wait()
}

// 传入时间戳，获得该时间戳所对应的电价级别
// in  ： 13位时间戳字符串             string
// out ： 级别（波谷：1 波平：2 波峰：3） int
func _GetPriceLevel(timeStamp string) (int, error) {
	format := "2006-01-02 15:04:05"

	now := time.Now()

	//--------------------------------------------------------------
	// 波峰时间段
	three_1_Start := fmt.Sprintf("%s 08:00:00", now.Format("2006-01-02"))
	three_1_End := fmt.Sprintf("%s 12:00:00", now.Format("2006-01-02"))
	three_2_Start := fmt.Sprintf("%s 19:00:00", now.Format("2006-01-02"))
	three_2_End := fmt.Sprintf("%s 23:00:00", now.Format("2006-01-02"))

	three_1_Start_T, _ := time.Parse(format, three_1_Start)
	three_1_End_T, _ := time.Parse(format, three_1_End)
	three_2_Start_T, _ := time.Parse(format, three_2_Start)
	three_2_End_T, _ := time.Parse(format, three_2_End)

	//--------------------------------------------------------------
	// 波平时间段
	two_1_Start := fmt.Sprintf("%s 07:00:00", now.Format("2006-01-02"))
	two_1_End := fmt.Sprintf("%s 08:00:00", now.Format("2006-01-02"))
	two_2_Start := fmt.Sprintf("%s 12:00:00", now.Format("2006-01-02"))
	two_2_End := fmt.Sprintf("%s 19:00:00", now.Format("2006-01-02"))

	two_1_Start_T, _ := time.Parse(format, two_1_Start)
	two_1_End_T, _ := time.Parse(format, two_1_End)
	two_2_Start_T, _ := time.Parse(format, two_2_Start)
	two_2_End_T, _ := time.Parse(format, two_2_End)

	//--------------------------------------------------------------
	// 波谷时间段
	one_1_Start := fmt.Sprintf("%s 23:00:00", now.Format("2006-01-02"))
	one_1_End := fmt.Sprintf("%s 00:00:00", now.Add(time.Hour*24).Format("2006-01-02"))
	one_2_Start := fmt.Sprintf("%s 00:00:00", now.Format("2006-01-02"))
	one_2_End := fmt.Sprintf("%s 07:00:00", now.Format("2006-01-02"))

	one_1_Start_T, _ := time.Parse(format, one_1_Start)
	one_1_End_T, _ := time.Parse(format, one_1_End)
	one_2_Start_T, _ := time.Parse(format, one_2_Start)
	one_2_End_T, _ := time.Parse(format, one_2_End)

	//--------------------------------------------------------------
	timeStamp_, _ := strconv.Atoi(timeStamp)
	tm := time.Unix(int64(timeStamp_)/1000, 0)
	timeTest, err := time.Parse(format, tm.Format(format))
	if err != nil {
		return 0, err
	}

	//--------------------------------------------------------------
	if (timeTest.After(three_1_Start_T) && timeTest.Before(three_1_End_T)) ||
		(timeTest.After(three_2_Start_T) && timeTest.Before(three_2_End_T)) {
		return 3, nil
	} else if (timeTest.After(two_1_Start_T) && timeTest.Before(two_1_End_T)) ||
		(timeTest.After(two_2_Start_T) && timeTest.Before(two_2_End_T)) {
		return 2, nil
	} else if (timeTest.After(one_1_Start_T) && timeTest.Before(one_1_End_T)) ||
		(timeTest.After(one_2_Start_T) && timeTest.Before(one_2_End_T)) {
		return 1, nil
	}

	return 0, nil
}

// 计算电费
func _CalcElecPrice(electricity, electricityTotal float64, timeStamp string) (float64, error) {
	var price float64
	var err error

	// 获得各阶梯电价
	prices, err := rethinkdb.GetPrice()
	if err != nil {
		return price, err
	}

	// 判断属于哪个阶梯
	slPrices := make([]float64, 3) // 峰、平、谷电价
	for _, v := range prices {
		low, ok := v["Low"].(float64)
		if !ok {
			return price, fmt.Errorf("v[\"Low\"].(float64) is error")
		}
		high, ok := v["High"].(float64)
		if !ok {
			return price, fmt.Errorf("v[\"High\"].(float64) is error")
		}

		if electricityTotal >= low && low <= high {
			slPrices[0], ok = v["One"].(float64)
			if !ok {
				return price, fmt.Errorf("v[\"One\"].(float64) is error")
			}

			slPrices[1], ok = v["Two"].(float64)
			if !ok {
				return price, fmt.Errorf("v[\"Two\"].(float64) is error")
			}

			slPrices[2], ok = v["Three"].(float64)
			if !ok {
				return price, fmt.Errorf("v[\"Three\"].(float64) is error")
			}
			break
		}
	}

	// 判断波峰平谷
	flag, err := _GetPriceLevel(timeStamp)
	if err != nil {
		return price, err
	}

	if flag == 0 {
		return price, fmt.Errorf("timeStamp is error")
	}

	// 计算电价
	price = slPrices[flag-1] * electricity

	return price, nil
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
//查询电表余额FuncQueryAmmeterBalance(user_A)
//查询用户账户余额FuncQueryAccountBalance(user_A)
//短信提示账户充钱FuncNoticeDeposit(user_A, user_remainmoney)
//电表自动购电50元FuncAutoPurchasingElectricity(user_A, user_B, amount)
//自动休眠1小时FuncAutoSleeping(sleeptime)
//获取查询起始时间FuncGetStartEndTime(user_A)
//获取电表消耗电量等信息FuncGetPowerConsumeParam(user_A, stat_begintime, stat_endtime)
//获取查询起始时间FuncGetPowerPrice()
//计算用户消耗的电费FuncCalcConsumeAmountAndMoney(user_A, elec_amount, elec_month_tatalamount, stat_begintime, stat_endtime)
//打印分账票据FuncTransferElecChargeToPlatform(user_B, user_others, user_transfers)
//修改电表余额FuncUpdateElecBalance(user_A, elec_amount)
//计算合约分账比例FuncCalcAndSplitRatio(user_B, stat_begintime, stat_endtime)
//合约分账FuncAutoSplitAccount(user_B, Split_percent, money)

//查询电表余额
//    访问电力能源链，读取用户电表余额
//Args: User_A  string  电表公钥
func FuncQueryAmmeterBalance(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncQueryAmmeterBalance")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//publickey, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	publickey := slMeterKey[0].Key

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
	logs.Info("*********************************************************")
	logs.Info("FuncQueryAccountBalance")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//publickey, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	publickey := slPersonKey[0]

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
//Args: User_A  string     用户账户
//      50      float64    充值额度
func FuncNoticeDeposit(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncNoticeDeposit")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}

	//publickey, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//money, ok := args[1].(float64)
	//if !ok {
	//	v_result.SetMessage("args[1].(float64) is error!")
	//	return v_result, v_err
	//}

	publickey := slPersonKey[0]
	money := float64(50)

	var msgNotice model.DemoMsgNotice
	msgNotice.Id = common2.GenerateUUID()
	msgNotice.NoticePublicKey = publickey
	msgNotice.Timestamp = common2.GenTimestamp()
	msgNotice.Msg = fmt.Sprintf("请及时充值%0.2f元到交易账户中，谢谢您的合作。", money)
	msgNotice.Type = 0

	slData, _ := json.Marshal(msgNotice)

	v_err = rethinkdb.InsertEnergyTradingDemoMsgNotice(string(slData))
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	// 充值
	_Recharge()

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//账户充值
func _Recharge() {
	// bill
	/*
			type DemoBill struct {
			Id        string `json:"id"`
			PublicKey string
			Timestamp string
			Type      int // 0：用户账户充值 1：用户购电充值 2：分张
		}
	*/
	strPublicKey, _ := common2.GenerateKeyPair()
	bill1 := model.DemoBill{
		Id:        common2.GenerateUUID(),
		PublicKey: strPublicKey,
		Timestamp: common2.GenTimestamp(),
		Type:      0,
	}
	sldata, _ := json.Marshal(bill1)
	rethinkdb.InsertEnergyTradingDemoBill(string(sldata))

	// transaction
	/*
			type DemoTransaction struct {
			Id            string  `json:"id"`
			BillId        string  // 对应的票据表id
			Timestamp     string  // 交易时间戳
			FromPublicKey string  // 付款方
			ToPublicKey   string  // 收款方
			Money         float64 // 金额
			Type          int     // 0：用户账户充值 1：用户购电充值 2：分张
		}
	*/
	transaction1 := model.DemoTransaction{
		Id:            common2.GenerateUUID(),
		BillId:        bill1.Id,
		Timestamp:     common.GenTimestamp(),
		FromPublicKey: "",
		ToPublicKey:   slPersonKey[0],
		Money:         300,
		Type:          0,
	}
	sldata, _ = json.Marshal(transaction1)

	rethinkdb.InsertEnergyTradingDemoTransaction(string(sldata))
}

//电表自动购电50元（链上进行资产转移50给运营账户；同时访问电表接口，给电表充值50元）
//Args:  User_A   string      将用户账户中的钱50元，转到运营商账户
//       Ccount_D string
//       50       float64     充值额度
func FuncAutoPurchasingElectricity(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncAutoPurchasingElectricity")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//userPublicKey, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//operatorPublicKey, ok := args[1].(string)
	//if !ok {
	//	v_result.SetMessage("args[1].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//money, ok := args[2].(float64)
	//if !ok {
	//	v_result.SetMessage("args[2].(float64) is error!")
	//	return v_result, v_err
	//}
	userPublicKey := slPersonKey[0]
	operatorPublicKey := slOperatorKey[0]
	money := float64(50)

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
		Money:         money,
		Type:          1,
	}
	sldata, _ = json.Marshal(transaction1)

	v_err = rethinkdb.InsertEnergyTradingDemoTransaction(string(sldata))
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	// 修改电表余额
	meterKey, v_err := rethinkdb.GetMeterKeyByUserKey(userPublicKey)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}
	mapMeterRemainMoney[meterKey] += money

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//自动休眠1小时
//Args： SleepTime  int 单位是秒
func FuncAutoSleeping(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncAutoSleeping")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//sleeptime, ok := args[0].(int)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}

	sleeptime := 30

	time.Sleep(time.Second * time.Duration(sleeptime))
	logs.Info("*********************************************************")
	logs.Info("sleep complete")
	logs.Info("*********************************************************")
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("auto sleep process success!")
	return v_result, v_err
}

//++++++++++++++++耗电消耗 及 自动分账合约++++++++++++++++++++++++++++++++++++++++++
// 获取查询起始时间
// userKey string
func FuncGetStartEndTime(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncGetStartEndTime")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//userPublicKey, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	userPublicKey := slPersonKey[0]

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
	strStartTime = lastTime

	// 获得当前时间
	nowTime := common2.GenTimestamp()
	strEndTime = nowTime

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
//Args: ElecUser_A  string  用户key
//      startTime   string
//      endTime     string
func FuncGetPowerConsumeParam(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncGetPowerConsumeParam")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//userPublicKey, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//startTime, ok := args[1].(string)
	//if !ok {
	//	v_result.SetMessage("args[1].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//endTime, ok := args[2].(string)
	//if !ok {
	//	v_result.SetMessage("args[2].(string) is error!")
	//	return v_result, v_err
	//}
	FuncGetStartEndTime()
	userPublicKey := slPersonKey[0]
	startTime := strStartTime
	endTime := strEndTime

	electricity, money, totalElectricity, v_err := rethinkdb.GetMeterInformation(userPublicKey, startTime, endTime)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}
	fElectricity = electricity
	fTotalElectricity = totalElectricity

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
	logs.Info("*********************************************************")
	logs.Info("FuncGetPowerPrice")
	logs.Info("*********************************************************")
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
	slData, _ := json.Marshal(price)
	v_result.SetData(string(slData))
	return v_result, v_err
}

//D. 根据用户耗电量对应出相应的电价，并计算用户消耗的电费、更新后的余额
//根据用户电表账户计算当前消耗的电量、将消耗对应电价计算消耗的金额、电表余额；
//Args: user_A              string   用户key
//      electricity         float64  当前耗电量
//      electricityTotal    float64  当月总耗电量
//      startTime           string   采集的开始时间
//      endTime             string   采集的终止时间
//Return: consume_money   消耗金额
//        remain_money    电表余额
func FuncCalcConsumeAmountAndMoney(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncCalcConsumeAmountAndMoney")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//userPublicKey, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//electricity, ok := args[1].(float64)
	//if !ok {
	//	v_result.SetMessage("args[1].(float64) is error!")
	//	return v_result, v_err
	//}
	//
	//electricityTotal, ok := args[2].(float64)
	//if !ok {
	//	v_result.SetMessage("args[2].(float64) is error!")
	//	return v_result, v_err
	//}
	//
	//startTime, ok := args[3].(string)
	//if !ok {
	//	v_result.SetMessage("args[3].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//endTime, ok := args[4].(string)
	//if !ok {
	//	v_result.SetMessage("args[4].(string) is error!")
	//	return v_result, v_err
	//}
	//_ = endTime
	userPublicKey := slPersonKey[0]
	electricity := fElectricity
	electricityTotal := fTotalElectricity
	startTime := strStartTime

	// 计算电价
	prices, v_err := _CalcElecPrice(electricity, electricityTotal, startTime)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	// 获得电表key
	meterKey, v_err := rethinkdb.GetMeterKeyByUserKey(userPublicKey)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}

	// 获得电表余额
	money, v_err := rethinkdb.GetMoneyFromEnergy(meterKey)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}
	fMoney = money

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData(fmt.Sprintf("{\"consume\":%f,\"money\":%f}", prices, money-prices))
	return v_result, v_err
}

//E. 打印分账票据，并记录到票据链上
//Args: user_B         string   运营商交易账户
//      other_users    string   合约分账的各用户
//      other_transfer string   各用户转账金额列表
func FuncTransferElecChargeToPlatform(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncTransferElecChargeToPlatform")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	// 空操作

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	return v_result, v_err
}

//F 修改电表余额
//Args：  user_A   string   elec_account
//       amount   float64      电表余额
func FuncUpdateElecBalance(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncUpdateElecBalance")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//userPublicKey, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//money, ok := args[1].(float64)
	//if !ok {
	//	v_result.SetMessage("args[1].(float64)!")
	//	return v_result, v_err
	//}
	userPublicKey := slPersonKey[0]
	money := fMoney

	// 修改电表余额
	meterKey, v_err := rethinkdb.GetMeterKeyByUserKey(userPublicKey)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}
	mapMeterRemainMoney[meterKey] = money

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
	logs.Info("*********************************************************")
	logs.Info("FuncCalcAndSplitRatio")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//strPowerPlants, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//// 反序列化各发电厂key
	//var slPowerPlants []string
	//v_err = json.Unmarshal([]byte(strPowerPlants), &slPowerPlants)
	//if v_err != nil {
	//	v_result.SetMessage(v_err.Error())
	//	return v_result, v_err
	//}
	//
	//if len(slPowerPlants) == 0 {
	//	if !ok {
	//		v_result.SetMessage("power plant key is null!")
	//		return v_result, v_err
	//	}
	//}
	//
	//startTime, ok := args[1].(string)
	//if !ok {
	//	v_result.SetMessage("args[1].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//endTime, ok := args[2].(string)
	//if !ok {
	//	v_result.SetMessage("args[2].(string) is error!")
	//	return v_result, v_err
	//}

	var slPowerPlants []string
	for _, v := range slPowerPlantsKey {
		slPowerPlants = append(slPowerPlants, v.Key)
	}
	startTime := strStartTime
	endTime := strEndTime

	// 获得各个发电厂此时段的发电量
	energys, v_err := rethinkdb.GetPowerPlantEnergy(slPowerPlants, startTime, endTime)
	if v_err != nil {
		v_result.SetMessage(v_err.Error())
		return v_result, v_err
	}
	mapEnergy = energys

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	slData, _ := json.Marshal(energys)
	v_result.SetData(string(slData))
	return v_result, v_err
}

//J. 根据各发电厂全天比例、进行合约分账
//Args: user_B        string   运营商交易账户
//      split_percent string   合约分账各方分账的比例
//      money         float    要分帐的金额
func FuncAutoSplitAccount(args ...interface{}) (common.OperateResult, error) {
	logs.Info("*********************************************************")
	logs.Info("FuncAutoSplitAccount")
	logs.Info("*********************************************************")
	var v_result common.OperateResult
	v_result.SetCode(500)
	var v_err error

	//if len(args) == 0 {
	//	v_result.SetMessage("param is null!")
	//	return v_result, v_err
	//}
	//
	//operatorKey, ok := args[0].(string)
	//if !ok {
	//	v_result.SetMessage("args[0].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//strPowerPlants, ok := args[1].(string)
	//if !ok {
	//	v_result.SetMessage("args[1].(string) is error!")
	//	return v_result, v_err
	//}
	//
	//// 反序列化发电厂分帐比例
	//var mapPowerPlants map[string]float64
	//v_err = json.Unmarshal([]byte(strPowerPlants), &mapPowerPlants)
	//if v_err != nil {
	//	v_result.SetMessage(v_err.Error())
	//	return v_result, v_err
	//}
	//
	//money, ok := args[2].(float64)
	//if !ok {
	//	v_result.SetMessage("args[2].(float64) is error!")
	//	return v_result, v_err
	//}
	operatorKey := slOperatorKey[0]
	mapPowerPlants := make(map[string]float64)
	mapPowerPlants = mapEnergy
	money := fMoney

	v_err = _AutoSplitAccount(money, operatorKey, mapPowerPlants)
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

func _AutoSplitAccount(money float64, strOperatorKey string, mapPowerPlants map[string]float64) (err error) {
	var totalElectricity float64

	for _, v := range mapPowerPlants {
		totalElectricity += v
	}

	if totalElectricity == 0 {
		return fmt.Errorf("totalElectricity is 0")
	}

	// 运营商留下20%
	money = money * 0.8

	var count int
	var money1 float64
	for key, value := range mapPowerPlants {
		count++
		// bill
		strPublicKey, _ := common2.GenerateKeyPair()
		bill1 := model.DemoBill{
			Id:        common2.GenerateUUID(),
			PublicKey: strPublicKey,
			Timestamp: common2.GenTimestamp(),
			Type:      2,
		}
		sldata, _ := json.Marshal(bill1)
		err = rethinkdb.InsertEnergyTradingDemoBill(string(sldata))
		if err != nil {
			return
		}

		// transaction
		moneySplit := money * (value / totalElectricity)
		if count == len(mapPowerPlants) {
			moneySplit = money - money1
		} else {
			money1 += moneySplit
		}
		transaction1 := model.DemoTransaction{
			Id:            common2.GenerateUUID(),
			BillId:        bill1.Id,
			Timestamp:     common.GenTimestamp(),
			FromPublicKey: strOperatorKey,
			ToPublicKey:   key,
			Money:         moneySplit,
			Type:          2,
		}
		sldata, _ = json.Marshal(transaction1)

		err = rethinkdb.InsertEnergyTradingDemoTransaction(string(sldata))
		if err != nil {
			return
		}
	}

	return nil
}

//K.获取当前时间（作为下次休眠判断的起始时间）
