package pipelines

import (
	"log"
	"testing"
)

func TestP(t *testing.T) {
	isInsert := false
	isDelete := false
	isUpdate := !isInsert && !isDelete
	log.Println(isUpdate)

	re := (4 | 1)

	log.Println(re)
}

//三元运算符计算
func TernaryOperator(p_cond bool, p_true_cond interface{}, p_false_cond interface{}) interface{} {
	if p_cond {
		return p_true_cond
	}
	return p_false_cond
}
