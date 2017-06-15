package function

import (
	"fmt"
	"strconv"
	"unicontract/src/core/engine/common"
)

//++++++++++++++++++++++++++++++++++++++++++++++++++++++
//++++++++++合约机【天安金交中心】专用扩展方法++++++++++
//++++++++++++++++++++++++++++++++++++++++++++++++++++++

//样例方法
func FuncTIANJSExample(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	var v_map_args map[string]interface{} = nil
	if len(args) != 0 {
		v_map_args = make(map[string]interface{}, 0)
	}
	//识别可变参数
	for v_idx, v_args := range args {
		tmp_arg := "v_arg_" + strconv.Itoa(v_idx)
		v_map_args[tmp_arg] = v_args
	}
	//调用参数
	for v_name, v_value := range v_map_args {
		fmt.Println(v_name, ":", v_value)
	}
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//===============理财产品申购================================================
//区分身份链、交易链（将理财看做钱份额，1元比1份）、理财产品链
//交易账户区分：客户交易账户、理财管理交易账户、托管管理交易账户
//交易区分  ：本金转移、收益转移
//理财产品链：
//交易链操作：
//     理财管理交易账户创建资产： 产品募集规模形成，创建理财管理资产

//理财产品表：募集期、募集期规模、份额封顶头寸

//A 【无需函数】读取合约产品参数(募集期)，并对募集期间进行判定(decision组件实现)
//B 【无需函数】在募集期间，获取输入的认购份额（合约属性中录入，直接读取即可）
//C 查询用户募集期内认购份额
//Args  ：User_A    string  用户公钥
//        Product_A string  产品ID
//Return：募集期内的  认购份额数
func FuncQueryUserTotalShareInPeriod(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	//查询交易链，获取已认购份额
	// product_A, user_A,  & timestamp in muji_begin--muji_end = acount
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//C2: 获取产品当期已募集总额
//Args： product_A  string  产品ID
//       user_B     stirng  运营方
//Return：募集期内的已购份额数
func FuncQueryProductTotalShareInPeriod(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	//查询交易链，获取已募集总份额数
	// product_A, user_B = acount
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//D 【无需函数】达到募集规模判定(decision组件实现)
//E 【无需函数】募集期间未达到募集规模：累加用户认购份额，计算份额是否达到封顶头寸判定(decision组件实现)
//F 募集期间达到募集规模：累加用户认购份额，询问管理员是否允许继续购买（外部输入）
//Args:  userB  string  运营方管理员
func FuncAskAdminIfContinue(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	//给消息队列推送针对管理员的询问消息

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//F2 获取管理员回复的是否继续购买信息
//Args:  userB  string  运营方管理员
func FuncGetAdminIfContinueReply(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	//给消息队列推送针对管理员的询问消息

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//G 达到封顶头寸  或 管理员不允许购买，认购失败
//任务退出：
func FuncExit(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	//认购失败，合约终止
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//H 未达到封顶头寸 或 管理员允许购买，认购成功
//任务完成
func FuncSuccess(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//I 【无需函数】募集期内有认购份额，新购份额不成立，同时判定是否达到募集规模(decision组件)
//J1:募集期外，没达到募集规模：获取用户的资金及利息
//Args: user_A    string  用户A
func FuncGetUserPrincipalAndInterest(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	//交易链中查询用户余额（本金和利息）
	// user_A
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//J 募集期外，没达到募集规模：原认购产品失败，退还资金及利息
//Args: user_A    string  用户A
//      user_B    string  中心账户B
//      amount    int     本金+利息
func FuncFailAndRefund(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//K 募集期外，达到募集规模：产品成立，原认购产品有效
//原产品有效，成功退出
//Args:  user_A  string
func FuncSuccessOrgProduct(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//L 签订认购协议
func FuncSignatureProduct(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil
	//新合约成立，自动签订新购买合约
	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//M 【无需函数】募集期内没有认购份额，判定是否达到募集规模（dicisison组件）
//N 未达到募集规模，产品失败，不可新购产品
func FuncFailNewProduct(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//O 达到募集规模，产品成立，可以新购产品，获取申购份额,认购成功
func FuncSuccessNewProduct(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

/*
//===============理财产品赎回================================================
//J   获取用户持有金额
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//K  获取用户持有期
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//L   持有期小于7天，赎回失败【不可赎回】
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//A.  持有期大于等于7天，获取用户赎回份额
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//B  全部赎回，终止合约，停止利息计算
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//C. 计算账户赎回总额度
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//E  大额赎回，通知管理员
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//F  获取上次大额赎回时间
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//G  是连续两天大额赎回，限制操作
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//H  不是连续两天大额赎回，确定赎回额度
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//I  小额赎回 或 大额赎回不受限 赎回转账
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//M 获取上一工作日理财计划的净产值
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//===============理财产品收益计算============================================
//A. 每天指定时间，查询截止当前的认购金额
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//B. 每天指定时间，查询截止当前的 账户余额
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//C. 判定上一工作日是否为募集期内
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//D 募集期内，理财合约未终止，查询人民币活期存款利率
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//E 募集期内，以人民币活期存款利率计算利息,并将利息转账给账户
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//G.募集期外，理财合约未终止，查询上一工作日实际年化收益率
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//H.募集期外，以上一工作日年化收益率计算利息,并将利息累加到认购金额中
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//F. 募集期外，计算用户理财实际收益，并将收益转入管理账户
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//I. 计算理财委托托管费，并由管理账户转账给托管人账户
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//O.计算用户预期收益，并由管理账户转账到用户账户
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//J.查询用户理财合约状态，是否为终止
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//K.募集期内，理财合约终止，停止计算利息
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//L 查询用户理财合约状态，是否为终止
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//M 募集期外，理财合约终止，停止计算利息
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//N.认购金额为0，停止计算利息
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//===============理财产品终止================================================
//A. 访问资源池查询终止条件1：不可抗力导致理财计划无法继续运行；
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//B. 访问资源池查询终止条件2：市场波动、异常风险事件发生；
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//C. 访问资源池查询终止条件3：付息人与管理人提前结束合作关系；
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//D. 访问资源池查询终止条件4：付息人违约；
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//E. 访问资源池查询终止条件5：申购资质账户少于2户；
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//F. 访问资源池查询终止条件6：理财规模低于100万份；
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//G. 判断终止条件，进行决策是否终止合约
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//H. 条件不满足，理财合约继续，sleep 5s
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//I. 条件满足，理财合约终止，记录终止日期
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//===============理财金额+利息 返还客户======================================
//A. 查询用户余额
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//B. 查询待付托管费用总额
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//C. 查询待付管理费用总额
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//D. 查询理财产品状态
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//E  产品状态为管理方终止：将用户账户余额（本金+预期利息），转账到银行账户
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//F  产品状态为正常终止：将用户账户余额（本金+预期利息），转账到银行账户
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//G 产品状态为运行中：查询当前日期
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//H  当前日期为季度末，将托管方账户余额转账到银行账户
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//J  当前日期不为季度末  或  托管账户余额转账成功，将管理费用转账到银行账户
func Func(args ...interface{}) (common.OperateResult, error) {
	var v_result common.OperateResult
	var v_err error = nil

	//构建返回值
	v_result = common.OperateResult{}
	v_result.SetCode(200)
	v_result.SetMessage("process success!")
	v_result.SetData("test success")
	return v_result, v_err
}

//===========================================================================
*/
