package function

import (
	"testing"
)

func Test_FuncQueryAmmeterBalance(t *testing.T) {
	publickey := "BLtALxedQpViBqQWSMDv6xbZpj1H6278CsLxAn2Na638"
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
