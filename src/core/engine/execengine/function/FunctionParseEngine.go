package function

import (
	"bytes"
	"fmt"

	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
)

type FunctionParseEngine struct {
	ContractFunctions map[string]func(arg ...interface{}) (common.OperateResult, error)
}

//---------------------------------------------------------------------------
//类构造方法
func NewFunctionParseEngine() *FunctionParseEngine {
	bif := &FunctionParseEngine{}
	bif.ContractFunctions = make(map[string]func(arg ...interface{}) (common.OperateResult, error), 0)
	return bif
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//根据配置文件名称加载函数、方法集
//=====Common Method
func (bif *FunctionParseEngine) LoadFunctionsCommon() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionsCommon fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		uniledgerlog.Error(r_buf.String())
		v_err = fmt.Errorf("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	bif.ContractFunctions["FuncGetNowDay"] = FuncGetNowDay
	bif.ContractFunctions["FuncGetNowDate"] = FuncGetNowDate
	bif.ContractFunctions["FuncSleepTime"] = FuncSleepTime
	bif.ContractFunctions["FuncGetNowDateTimestamp"] = FuncGetNowDateTimestamp
	bif.ContractFunctions["FuncTestMethod"] = FuncTestMethod
	bif.ContractFunctions["FuncTestMethod1"] = FuncTestMethod1
	bif.ContractFunctions["FuncCreateAsset"] = FuncCreateAsset
	bif.ContractFunctions["FuncTransferAsset"] = FuncTransferAsset
	bif.ContractFunctions["FuncTransferAssetComplete"] = FuncTransferAssetComplete
	bif.ContractFunctions["FuncUnfreezeAsset"] = FuncUnfreezeAsset
	bif.ContractFunctions["FuncInterim"] = FuncInterim
	bif.ContractFunctions["FuncInterimComplete"] = FuncInterimComplete
	bif.ContractFunctions["FuncIsConPutInUnichian"] = FuncIsConPutInUnichian
	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====Demo Method(产品演示Demo)
func (bif *FunctionParseEngine) LoadFunctionDEMO() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionDEMO fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		uniledgerlog.Error(r_buf.String())
		v_err = fmt.Errorf("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	bif.ContractFunctions["FuncGetBalance"] = FuncGetBalance
	bif.ContractFunctions["FuncTanferMoney"] = FuncTanferMoney
	bif.ContractFunctions["FuncDeposit"] = FuncDeposit
	bif.ContractFunctions["FuncQueryMonthConsumption"] = FuncQueryMonthConsumption
	bif.ContractFunctions["FuncBackTelephoneFare"] = FuncBackTelephoneFare
	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====TIANJS Method(天安金交所)
func (bif *FunctionParseEngine) LoadFunctionTIANJS() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionTIANJS fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		uniledgerlog.Error(r_buf.String())
		v_err = fmt.Errorf("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	// 购买合约
	bif.ContractFunctions["FuncQueryUserTotalShareInPeriod"] = FuncQueryUserTotalShareInPeriod
	bif.ContractFunctions["FuncQueryProductTotalShareInPeriod"] = FuncQueryProductTotalShareInPeriod
	bif.ContractFunctions["FuncAskAdminIfContinue"] = FuncAskAdminIfContinue
	bif.ContractFunctions["FuncGetAdminIfContinueReply"] = FuncGetAdminIfContinueReply
	bif.ContractFunctions["FuncPurchaseExit"] = FuncPurchaseExit
	bif.ContractFunctions["FuncPurchaseSuccess"] = FuncPurchaseSuccess
	bif.ContractFunctions["FuncGetUserPrincipalAndInterest"] = FuncGetUserPrincipalAndInterest
	bif.ContractFunctions["FuncPurchaseFailAndRefund"] = FuncPurchaseFailAndRefund
	bif.ContractFunctions["FuncPurchaseSuccessOrgProduct"] = FuncPurchaseSuccessOrgProduct
	bif.ContractFunctions["FuncSignatureProduct"] = FuncSignatureProduct
	bif.ContractFunctions["FuncPurchaseFailNewProduct"] = FuncPurchaseFailNewProduct
	bif.ContractFunctions["FuncPurchaseSuccessNewProduct"] = FuncPurchaseSuccessNewProduct
	// 赎回合约
	bif.ContractFunctions["FuncGetUserTotalShare"] = FuncGetUserTotalShare
	bif.ContractFunctions["FuncGetUserHoldPeriod"] = FuncGetUserHoldPeriod
	bif.ContractFunctions["FuncRedeemFail"] = FuncRedeemFail
	bif.ContractFunctions["FuncRedeemAllProcess"] = FuncRedeemAllProcess
	bif.ContractFunctions["FuncCalcTotalAmount"] = FuncCalcTotalAmount
	bif.ContractFunctions["FuncRedeemLargeProcess"] = FuncRedeemLargeProcess
	bif.ContractFunctions["FuncGetLastRedeemLargeTime"] = FuncGetLastRedeemLargeTime
	bif.ContractFunctions["FuncRedeemLimit"] = FuncRedeemLimit
	bif.ContractFunctions["FuncRedeemSmallProcess"] = FuncRedeemSmallProcess
	bif.ContractFunctions["FuncTotalOutValueLastDay"] = FuncTotalOutValueLastDay
	// 收益计算
	bif.ContractFunctions["FuncGetUserPurchase"] = FuncGetUserPurchase
	bif.ContractFunctions["FuncGetUserBalance"] = FuncGetUserBalance
	bif.ContractFunctions["FuncCheckLastDayInRaisePeriod"] = FuncCheckLastDayInRaisePeriod
	bif.ContractFunctions["FuncGetDepositRate"] = FuncGetDepositRate
	bif.ContractFunctions["FuncCalcAndTransferInterest"] = FuncCalcAndTransferInterest
	bif.ContractFunctions["FuncGetYearYieldRateOfLastDay"] = FuncGetYearYieldRateOfLastDay
	bif.ContractFunctions["FuncCalcUserRealIncome"] = FuncCalcUserRealIncome
	bif.ContractFunctions["FuncCalcAndTransferTrusteeTee"] = FuncCalcAndTransferTrusteeTee
	bif.ContractFunctions["FuncCalcAndTransferExpectIncome"] = FuncCalcAndTransferExpectIncome
	bif.ContractFunctions["FuncQueryContractState"] = FuncQueryContractState
	bif.ContractFunctions["FuncQueryContractState"] = FuncQueryContractState
	bif.ContractFunctions["FuncTerminateContract"] = FuncTerminateContract
	bif.ContractFunctions["FuncStopCalcInterest"] = FuncStopCalcInterest
	// 合约终止
	bif.ContractFunctions["FuncGetConditionState"] = FuncGetConditionState
	bif.ContractFunctions["FuncAbnormalEnd"] = FuncAbnormalEnd
	// 账务结算
	bif.ContractFunctions["FuncUserTotalRemain"] = FuncUserTotalRemain
	bif.ContractFunctions["FuncPayTotalTrustFee"] = FuncPayTotalTrustFee
	bif.ContractFunctions["FuncPayTotalManageFee"] = FuncPayTotalManageFee
	bif.ContractFunctions["FuncGetProductState"] = FuncGetProductState
	bif.ContractFunctions["FuncBankTransfer"] = FuncBankTransfer
	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====GUANGXIBIANMAO Method(广西边贸)
func (bif *FunctionParseEngine) LoadFunctionGUANGXIBIANMAO() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionGUANGXIBIAMAO fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		uniledgerlog.Error(r_buf.String())
		v_err = fmt.Errorf("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	bif.ContractFunctions["FuncBIANMAOExample"] = FuncBIANMAOExample

	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====ENEGYTRADING [能源交易]
func (bif *FunctionParseEngine) LoadFunctionENERGYTRADING() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionENEGYTRADING fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		uniledgerlog.Error(r_buf.String())
		v_err = fmt.Errorf("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	bif.ContractFunctions["FuncBIANMAOExample"] = FuncBIANMAOExample
	bif.ContractFunctions["FuncQueryAmmeterBalance"] = FuncQueryAmmeterBalance
	bif.ContractFunctions["FuncQueryAccountBalance"] = FuncQueryAccountBalance
	bif.ContractFunctions["FuncNoticeDeposit"] = FuncNoticeDeposit
	bif.ContractFunctions["FuncAutoPurchasingElectricity"] = FuncAutoPurchasingElectricity
	bif.ContractFunctions["FuncAutoSleeping"] = FuncAutoSleeping
	bif.ContractFunctions["FuncGetStartEndTime"] = FuncGetStartEndTime
	bif.ContractFunctions["FuncGetPowerConsumeParam"] = FuncGetPowerConsumeParam
	bif.ContractFunctions["FuncGetPowerPrice"] = FuncGetPowerPrice
	bif.ContractFunctions["FuncCalcConsumeAmountAndMoney"] = FuncCalcConsumeAmountAndMoney
	bif.ContractFunctions["FuncTransferElecChargeToPlatform"] = FuncTransferElecChargeToPlatform
	bif.ContractFunctions["FuncUpdateElecBalance"] = FuncUpdateElecBalance
	bif.ContractFunctions["FuncCalcAndSplitRatio"] = FuncCalcAndSplitRatio
	bif.ContractFunctions["FuncAutoSplitAccount"] = FuncAutoSplitAccount
	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====TRANSFER [转账支付]
func (bif *FunctionParseEngine) LoadFunctionTRANSFER() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionTRANSFER fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		uniledgerlog.Error(r_buf.String())
		v_err = fmt.Errorf("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	//TODO: when add method in ContractFunctionForCommon.go，must add it here
	bif.ContractFunctions["FuncAutoTransferAssetAtTime"] = FuncAutoTransferAssetAtTime

	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====RENTPAYMENT 【房租自动缴纳合约】
func (bif *FunctionParseEngine) LoadFunctionRENTPAYMENT() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionRENTPAYMENT fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		uniledgerlog.Error(r_buf.String())
		v_err = fmt.Errorf("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	bif.ContractFunctions["FuncIfContinueToPayNextMonth"] = FuncIfContinueToPayNextMonth
	bif.ContractFunctions["FuncContractExitForComplete"] = FuncContractExitForComplete
	bif.ContractFunctions["FuncQueryUserBalance"] = FuncQueryUserBalance
	bif.ContractFunctions["FuncTransferMoney"] = FuncTransferMoney
	bif.ContractFunctions["FuncPrintReceipt"] = FuncPrintReceipt
	bif.ContractFunctions["FuncRemindAccount"] = FuncRemindAccount
	bif.ContractFunctions["FuncNoAction"] = FuncNoAction

	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====HOUSETRANSFER 【房屋交易】
func (bif *FunctionParseEngine) LoadFunctionHOUSETRANSFER() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionHOUSETRANSFER fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		uniledgerlog.Error(r_buf.String())
		v_err = fmt.Errorf("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	bif.ContractFunctions["FuncQueryHouse"] = FuncQueryHouse
	bif.ContractFunctions["FuncExitForNoHouse"] = FuncExitForNoHouse
	bif.ContractFunctions["FuncUserBalance"] = FuncUserBalance
	bif.ContractFunctions["FuncTransferHouseFees"] = FuncTransferHouseFees
	bif.ContractFunctions["FuncNoticeRecharge"] = FuncNoticeRecharge
	bif.ContractFunctions["FuncQueryHouseFeesResult"] = FuncQueryHouseFeesResult
	bif.ContractFunctions["FuncExitForTransferFail"] = FuncExitForTransferFail
	bif.ContractFunctions["FuncTransferHouse"] = FuncTransferHouse
	bif.ContractFunctions["FuncQueryHouseResult"] = FuncQueryHouseResult
	bif.ContractFunctions["FuncExitForSuccess"] = FuncExitForSuccess
	bif.ContractFunctions["FuncExitForHouseTransferFail"] = FuncExitForHouseTransferFail

	return v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//=====AUTOSELLER 【自动售卖机交易】
func (bif *FunctionParseEngine) LoadFunctionAUTOSELLER() error {
	var v_err error = nil
	var r_buf bytes.Buffer = bytes.Buffer{}
	if bif.ContractFunctions == nil {
		r_buf.WriteString("[Result]: LoadFunctionAUTOSELLER fail;")
		r_buf.WriteString("[Error]: ContractFunctions is nil!")
		uniledgerlog.Error(r_buf.String())
		v_err = fmt.Errorf("ContractFunctions is nil!")
		return v_err
	}
	//Add Common Method,Here
	bif.ContractFunctions["FuncGetUserSelectedStyle"] = FuncGetUserSelectedStyle
	bif.ContractFunctions["FuncGetUserSelectedCount"] = FuncGetUserSelectedCount
	bif.ContractFunctions["FuncQueryRemainingCount"] = FuncQueryRemainingCount
	bif.ContractFunctions["FuncExitForNoRemaining"] = FuncExitForNoRemaining
	bif.ContractFunctions["FuncCalculatedCost"] = FuncCalculatedCost
	bif.ContractFunctions["FuncWaitPayMoney"] = FuncWaitPayMoney
	bif.ContractFunctions["FuncQueryUserPayCount"] = FuncQueryUserPayCount
	bif.ContractFunctions["FuncSupplyGoods"] = FuncSupplyGoods
	bif.ContractFunctions["FuncQueryRemainingMoney"] = FuncQueryRemainingMoney
	bif.ContractFunctions["FuncOddChange"] = FuncOddChange
	bif.ContractFunctions["FuncExitForSuccess"] = FuncExitForSuccess
	bif.ContractFunctions["FuncExitForTerminal"] = FuncExitForTerminal
	return v_err
}
