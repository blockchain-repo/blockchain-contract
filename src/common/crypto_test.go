package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_HashData(t *testing.T) {
	//hash := HashData("hello unichain 2017")
	ddd := "hello unichain >= 2+ &2017"
	//ddd := `{"Caption":"购智能手机返话费合约产品协议","Cname":"contract_mobilecallback","ContractId":"contract test from js->java->go-python","ContractOwners":["3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ","BbfY4Dc5s8dTP1Z1yixnetezRKYREHqwbP8GQGh3WyVS"],"StartTime":"1496535120000"}`
	t.Log(ddd)
	aaa := ddd
	aaa = Serialize(ddd)
	t.Log(aaa)
	hash := HashData(aaa)
	data := "2c27fa14ff62005acda2b1845bb335f5c139ff252670d6d86cf9801617120037"
	if hash != data {
		t.Error("Test_HashData error")
	}
	fmt.Println("----------------------data:", "hello unichain 2017")
	fmt.Println("----------------------hash:", hash)
}

func Test_GenerateKeyPair(t *testing.T) {
	publicKeyBase58, privateKeyBase58 := GenerateKeyPair()
	fmt.Println("----------------------pub:", publicKeyBase58)
	fmt.Println("----------------------pri:", privateKeyBase58)
}

func Test_GetPubByPriv(t *testing.T) {
	pub := "BbfY4Dc5s8dTP1Z1yixnetezRKYREHqwbP8GQGh3WyVS"
	pri := "6hXsHQ4fdWQ9UY1XkBYCYRouAagRW8rXxYSLgpveQNYY"
	pub2 := GetPubByPriv(pri)
	if pub != pub2 {
		t.Error("Test_GetPubByPriv error")
	}
}

func Test_Sign(t *testing.T) {
	buf, _ := ioutil.ReadFile("./produce_contract.json")
	t.Log(string(buf))
	var contract interface{}
	json.Unmarshal(buf, &contract)
	t.Log(contract)

	//msg := "hello unichain 2017"
	msg := "china"
	t.Log(msg)
	//pri := "6hXsHQ4fdWQ9UY1XkBYCYRouAagRW8rXxYSLgpveQNYY"
	pri := "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	sig := "48cpAsUuNf6qKCMFFKitSNjaA8nfPM4o7MacVp8U3QVMbVUr34SSRTTpahi3WEv3GaF2bVWG7J4SLTojgDoacLxT"
	sig2 := Sign(pri, msg)
	t.Log(sig2)
	if sig != sig2 {
		t.Error("Test_Sign error")
	}
}

func Test_Verify(t *testing.T) {
	pub := "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	pub2 := "AZfjdKxEr9G3NwdAkco22nN8PfgQvCr5TDPK1tqsGZrk"
	msg := "hello unichain 2017"
	sig := "48cpAsUuNf6qKCMFFKitSNjaA8nfPM4o7MacVp8U3QVMbVUr34SSRTTpahi3WEv3GaF2bVWG7J4SLTojgDoacLxT"
	if !Verify(pub, msg, sig) {
		t.Error("Test_Verify error")
	}
	if Verify(pub2, msg, sig) {
		t.Error("Test_Verify error")
	}
}
