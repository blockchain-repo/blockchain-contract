package function

import (
	"encoding/json"
	"testing"
	"unicontract/src/common"
)

func Test_FuncQueryAmmeterBalance(t *testing.T) {
	publickey := "5x1hxnPWpHRpvwR3tdo7ygPZ77sSUkywY56VhGaLpUm"
	result, err := FuncQueryAmmeterBalance(publickey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_FuncQueryAccountBalance(t *testing.T) {
	publickey := "64mDgEqY9KGp3NCfJPrrjiruL9hmuYiimmaD2234UYWd"
	result, err := FuncQueryAccountBalance(publickey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_FuncNoticeDeposit(t *testing.T) {
	publickey := "64mDgEqY9KGp3NCfJPrrjiruL9hmuYiimmaD2234UYWd"
	result, err := FuncNoticeDeposit(publickey, float64(60))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_FuncAutoPurchasingElectricity(t *testing.T) {
	user := "64mDgEqY9KGp3NCfJPrrjiruL9hmuYiimmaD2234UYWd"
	operator := "95b4DQfoNCh3o6jdy2k2AjCoZQrSUVubC5fFxEfRDpPH"
	money := 40
	result, err := FuncAutoPurchasingElectricity(user, operator, float64(money))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_FuncGetPowerPrice(t *testing.T) {
	result, err := FuncGetPowerPrice()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_FuncGetStartEndTime(t *testing.T) {
	result, err := FuncGetStartEndTime("64mDgEqY9KGp3NCfJPrrjiruL9hmuYiimmaD2234UYWd")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_FuncGetPowerConsumeParam(t *testing.T) {
	result, err := FuncGetPowerConsumeParam("64mDgEqY9KGp3NCfJPrrjiruL9hmuYiimmaD2234UYWd", "1497595656727", common.GenTimestamp())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_FuncCalcConsumeAmountAndMoney(t *testing.T) {
	result, err := FuncCalcConsumeAmountAndMoney("64mDgEqY9KGp3NCfJPrrjiruL9hmuYiimmaD2234UYWd", float64(20), float64(200), "1497675596000", common.GenTimestamp())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_FuncCalcAndSplitRatio(t *testing.T) {
	slKey := []string{
		"9Vqg4tSk9ocLfhwj2eeNgKgNR65oSV7WF9kYDu1HiwdM", // 风
		"3XmEh9ZtvDAcxtgiFL11cw9YAppCqhQaWQ6mrKxWhbom", // 光
		"H7tMDKFPMGsG2pV4Lpcic5MQiN1fKkqVaj6A15MMgNTQ", // 火
		"4nkFyWhLrUAGZxr1Ku5NreywPPA6HEKkpqV2hDgr1kLU", // 国网
	}
	slData, _ := json.Marshal(slKey)
	result, err := FuncCalcAndSplitRatio(string(slData), "1497679290000", "1497682890000")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_FuncAutoSplitAccount(t *testing.T) {
	str := `{"3XmEh9ZtvDAcxtgiFL11cw9YAppCqhQaWQ6mrKxWhbom":105,"4nkFyWhLrUAGZxr1Ku5NreywPPA6HEKkpqV2hDgr1kLU":150,"9Vqg4tSk9ocLfhwj2eeNgKgNR65oSV7WF9kYDu1HiwdM":120,"H7tMDKFPMGsG2pV4Lpcic5MQiN1fKkqVaj6A15MMgNTQ":190}`
	result, err := FuncAutoSplitAccount("95b4DQfoNCh3o6jdy2k2AjCoZQrSUVubC5fFxEfRDpPH", str, float64(200))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func Test_init(t *testing.T) {
	t.Log(slPowerPlantsKey)
	t.Log(slMeterKey)
}
