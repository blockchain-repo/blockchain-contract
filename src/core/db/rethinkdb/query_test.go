package rethinkdb

import (
	"fmt"
	"testing"

	"unicontract/src/common"
)

func Test_Get(t *testing.T) {
	res :=Get("Unicontract","Contract","123151f1ddassd")
	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}
	str := common.Serialize(blo)
	fmt.Printf("blo:%s\n",str)

}

func Test_Insert(t *testing.T) {
	res :=Insert("bigchain","votes","{\"back\":\"jihhh\"}")
	fmt.Printf("%d row inserted", res.Inserted)
}

func Test_Update(t *testing.T) {
	res :=Update("bigchain","votes","37adc1b6-e22a-4d39-bc99-f1f44608a15b","{\"1111back\":\"j111111111111ihhh\"}")
	fmt.Printf("%d row replaced", res.Replaced)
}

func Test_Delete(t *testing.T) {
        res :=Delete("bigchain","votes","37adc1b6-e22a-4d39-bc99-f1f44608a15b")
        fmt.Printf("%d row deleted", res.Deleted)
}
