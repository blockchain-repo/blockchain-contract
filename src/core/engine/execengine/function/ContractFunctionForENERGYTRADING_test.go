package function

import (
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
	money := 50
	result, err := FuncNoticeDeposit(publickey, money)
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
	result, err := FuncAutoPurchasingElectricity(user, operator, money)
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
