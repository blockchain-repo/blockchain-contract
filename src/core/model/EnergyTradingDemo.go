package model

// table [TABLE_ENERGYTRADINGDEMO_ROLE]
type DemoRole struct {
	Id            string `json:"id"`
	Name          string // 标识
	PublicKey     string
	Infermation   string // 如果是电表，记录所对应人的publickey
	LastTimestamp string // 仅当是电表时有效，代表上次查询耗电量的时间点
	Type          int    // 0：人 1：电表 2：运营商 3：风电 4：光电 5：火电 6：国网
}

// table [TABLE_ENERGYTRADINGDEMO_ENERGY]
type DemoEnergy struct {
	Id               string `json:"id"`
	PublicKey        string
	Timestamp        string  // 时间戳
	Electricity      float64 // 当前总电量（电）
	TotalElectricity float64 // 只当为电表时有效，当月总耗电量（电）
	Money            float64 // 只当为电表时有效，代表当前表内余额（钱）
	Type             int     // 0：电表 1：风电 2：光电 3：火电 4：国网
}

// table [TABLE_ENERGYTRADINGDEMO_TRANSACTION]
type DemoTransaction struct {
	Id            string  `json:"id"`
	BillId        string  // 对应的票据表id
	Timestamp     string  // 交易时间戳
	FromPublicKey string  // 付款方
	ToPublicKey   string  // 收款方
	Money         float64 // 金额
	Type          int     // 0：用户账户充值 1：用户购电充值 2：分张
}

// table [TABLE_ENERGYTRADINGDEMO_BILL]
type DemoBill struct {
	Id        string `json:"id"`
	PublicKey string
	Timestamp string
	Type      int // 0：用户账户充值 1：用户购电充值 2：电表耗电 3：分张
}

// table [TABLE_ENERGYTRADINGDEMO_MSGNOTICE]
type DemoMsgNotice struct {
	Id              string `json:"id"`
	NoticePublicKey string
	Timestamp       string
	Msg             string
	Type            int // 0：用户欠费通知
}

// table [TABLE_ENERGYTRADINGDEMO_PRICE]
type DemoPrice struct {
	Id          string  `json:"id"`
	Level       int     // 阶梯等级
	One         float64 // 波峰电价
	Two         float64 // 波平电价
	Three       float64 // 波谷电价
	Description string  // 阶梯电价描述
}
