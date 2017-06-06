package crypto

import (
	"testing"
)

func Test_AesEncryptToFile(t *testing.T) {
	err := AesEncryptToFile("../../../conf/unicontract", "./unicontract")
	if err != nil {
		t.Error(err)
	}
}

func Test_AesDecryptFromFile(t *testing.T) {
	origData, err := AesDecryptFromFile("./unicontract")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(origData))
	}
}
