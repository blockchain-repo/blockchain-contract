package data

import (
	"fmt"
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
)

func CreateIntDataObject() *IntData {
	t_int := new(IntData)
	t_int.InitIntData()
	t_int.SetCname("TestMoney")
	t_int.SetDefaultValue(0)
	t_int.SetCaption("money")
	t_int.SetDescription("money of china")
	t_int.SetUnit("元")

	return t_int
}

func TestIntInit(t *testing.T) {
	t_int := new(IntData)
	t_int.InitIntData()
	t_int.SetCname("TestMoney")
	t_int.SetDefaultValue(0)
	t_int.SetCaption("money")
	t_int.SetDescription("money of china")
	t_int.SetUnit("元")
	if t_int == nil {
		t.Error("IntData init Error!")
	}
	if t_int.GetCtype() != constdef.ComponentType[constdef.Component_Data]+"."+constdef.DataType[constdef.Data_Numeric_Int] {
		t.Error("t_ctype value Error!")
	}
	if t_int.GetCname() != "TestMoney" {
		t.Error("t_name value Error!")
	}
	if t_int.GetDefaultValue() != 0 {
		t.Error("t_default value Error!")
	}
	if t_int.GetUnit() != "元" {
		t.Error("t_unit value Error!")
	}
	if t_int.GetCaption() != "money" {
		t.Error("t_caption value Error!")
	}
	if t_int.GetDescription() != "money of china" {
		t.Error("t_description value Error!")
	}
	if t_int.GetHardConvType() != "int" {
		t.Error("hardConvType value Error!")
	}
	fmt.Println(t_int.GetDataRange())
	if t_int.GetDataRange()[0] != -2147483647 {
		t.Error("dataRange left value Error!")
	}
	if t_int.GetDataRange()[1] != 2147483647 {
		t.Error("dataRange right value Error!")
	}
	fmt.Println(t_int.GetCname(), " ", t_int.GetDefaultValue(), " ", t_int.GetUnit(), " ", t_int.GetCaption(), " ", t_int.GetDescription())
}

func TestIntDataRange(t *testing.T) {
	t_int := CreateIntDataObject()
	var t_range_1 [2]int = [2]int{-1, -1}
	var t_range_2 [2]int = [2]int{-1, 0}
	var t_range_3 [2]int = [2]int{0, 1}
	var t_range_4 [2]int = [2]int{1, 1}
	var t_range_5 [2]int = [2]int{1, -1}
	//default [0, 0]
	var t_range_6 [2]int
	if err := t_int.SetDataRange(t_range_1); err != nil {
		t.Error("[-1,-1] process error")
	}
	if err := t_int.SetDataRange(t_range_2); err != nil {
		t.Error("[-1,0] process error")
	}
	if err := t_int.SetDataRange(t_range_3); err != nil {
		t.Error("[0,1] process error")
	}
	if err := t_int.SetDataRange(t_range_4); err != nil {
		t.Error("[1,1] process error")
	}
	if err := t_int.SetDataRange(t_range_5); err == nil {
		t.Error("[1,-1] process error")
	}
	if err := t_int.SetDataRange(t_range_6); err != nil {
		t.Error("[] process error")
	} else {
		if t_int.GetDataRange()[0] == 0 {
			t.Error("dataRange left value Default Error!")
		}
	}
}

func TestCheckRange(t *testing.T) {
	t_int := CreateIntDataObject()
	var t_range [2]int = [2]int{0, 100}
	t_int.SetDataRange(t_range)
	if t_int.CheckRange(10) != true {
		t.Error("check error, 10 must in [0, 100]")
	}
	if t_int.CheckRange(101) != false {
		t.Error("check error, 101 not in [0, 100]")
	}
	if t_int.CheckRange(0) != true {
		t.Error("check error, 0 must in [0, 100]")
	}
	if t_int.CheckRange(100) != true {
		t.Error("check error, 100 must in [0, 100]")
	}
}

func TestAdd(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(100)
	if v_int, err := t_int.Add(200); err == nil {
		if v_int != 300 {
			t.Error("IntData add error,result not equal 300!")
		} else {
			fmt.Println("100 + 200 = ", v_int)
		}
	} else {
		t.Error("IntData add error!")
	}
}

func TestRAdd(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(100)
	if v_int, err := t_int.RAdd(200); err == nil {
		if v_int != 300 {
			t.Error("IntData add error,result not equal 300!")
		} else {
			fmt.Println("200 + 100 = ", v_int)
		}
	} else {
		t.Error("IntData add error!")
	}
}

func TestSub(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(400)
	if v_int, err := t_int.Sub(200); err == nil {
		if v_int != 200 {
			t.Error("IntData Sub error,result not equal 200!")
		} else {
			fmt.Println("400 - 200 = ", v_int)
		}
	} else {
		t.Error("IntData Sub error!")
	}
}

func TestRSub(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(100)
	if v_int, err := t_int.RSub(300); err == nil {
		if v_int != 200 {
			t.Error("IntData RSub error,result not equal 200!")
		} else {
			fmt.Println("300 - 100 = ", v_int)
		}
	} else {
		t.Error("IntData RSub error!")
	}
}

func TestMul(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(4)
	if v_int, err := t_int.Mul(2); err == nil {
		if v_int != 8 {
			t.Error("IntData Mul error,result not equal 8!")
		} else {
			fmt.Println("4 * 2 = ", v_int)
		}
	} else {
		t.Error("IntData Mul error!")
	}
}

func TestRMul(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(2)
	if v_int, err := t_int.RMul(4); err == nil {
		if v_int != 8 {
			t.Error("IntData RMul error,result not equal 8!")
		} else {
			fmt.Println("2 * 4 = ", v_int)
		}
	} else {
		t.Error("IntData RMul error!")
	}
}

func TestDiv(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(4)
	if v_int, err := t_int.Div(2); err == nil {
		if v_int != 2 {
			t.Error("IntData Div error,result not equal 8!")
		} else {
			fmt.Println("4 / 2 = ", v_int)
		}
	} else {
		t.Error("IntData Div error!")
	}
}

func TestDivError(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(4)
	if _, err := t_int.Div(0); err == nil {
		fmt.Println("4 / 0 = error")
		t.Error("IntData Div error, zero exist!")
	}
}

func TestRDiv(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(2)
	if v_int, err := t_int.RDiv(4); err == nil {
		if v_int != 2 {
			t.Error("IntData RDiv error,result not equal 8!")
		} else {
			fmt.Println("4 / 2 = ", v_int)
		}
	} else {
		t.Error("IntData RDiv error!")
	}
}

func TestMod(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(5)
	if v_int, err := t_int.Mod(2); err == nil {
		if v_int != 1 {
			t.Error("IntData Mod error,result not equal 1!")
		} else {
			fmt.Println("5 % 2 = ", v_int)
		}
	} else {
		t.Error("IntData Mod error!")
	}
}

func TestNeg(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(100)
	var v_int int = t_int.Neg()
	if v_int != -100 {
		t.Error("IntData Neg error!")
	}
}

func TestLt(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(100)
	if v_bool, err := t_int.Lt(200); err == nil {
		if !v_bool {
			t.Error("IntData Lt error,result not equal true!")
		} else {
			fmt.Println("100 < 200: ", v_bool)
		}
	} else {
		t.Error("IntData Lt error!")
	}
}

func TestLtFail(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(200)
	if v_bool, err := t_int.Lt(100); err == nil {
		if v_bool {
			t.Error("IntData Lt error,result not equal true!")
		} else {
			fmt.Println("200 < 100: ", v_bool)
		}
	} else {
		t.Error("IntData Lt error!")
	}
}

func TestLtError(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(200)
	if _, err := t_int.Lt("a"); err == nil {
		t.Error("IntData Lt error!")
	}
}

func TestLe(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(100)
	if v_bool, err := t_int.Le(100); err == nil {
		if !v_bool {
			t.Error("IntData Le error,result not equal true!")
		} else {
			fmt.Println("100 <= 100: ", v_bool)
		}
	} else {
		t.Error("IntData Le error!")
	}
}

func TestLeFail(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(200)
	if v_bool, err := t_int.Le(100); err == nil {
		if v_bool {
			t.Error("IntData Le error,result not equal true!")
		} else {
			fmt.Println("200 < 100: ", v_bool)
		}
	} else {
		t.Error("IntData Le error!")
	}
}

func TestLeError(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(200)
	if _, err := t_int.Le("a"); err == nil {
		t.Error("IntData Lt error!")
	}
}

func TestGt(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(200)
	if v_bool, err := t_int.Gt(100); err == nil {
		if !v_bool {
			t.Error("IntData Gt error,result not equal true!")
		} else {
			fmt.Println("200 > 100: ", v_bool)
		}
	} else {
		t.Error("IntData Gt error!")
	}
}

func TestGtFail(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(100)
	if v_bool, err := t_int.Gt(200); err == nil {
		if v_bool {
			t.Error("IntData Gt error,result not equal true!")
		} else {
			fmt.Println("100 > 200: ", v_bool)
		}
	} else {
		t.Error("IntData Gt error!")
	}
}

func TestGtError(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(200)
	if _, err := t_int.Gt("a"); err == nil {
		t.Error("IntData Gt error!")
	}
}

func TestGe(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(200)
	if v_bool, err := t_int.Ge(200); err == nil {
		if !v_bool {
			t.Error("IntData Ge error,result not equal true!")
		} else {
			fmt.Println("200 > 200: ", v_bool)
		}
	} else {
		t.Error("IntData Ge error!")
	}
}

func TestGeFail(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(100)
	if v_bool, err := t_int.Ge(200); err == nil {
		if v_bool {
			t.Error("IntData Ge error,result not equal true!")
		} else {
			fmt.Println("100 > 200: ", v_bool)
		}
	} else {
		t.Error("IntData Ge error!")
	}
}

func TestGeError(t *testing.T) {
	t_int := CreateIntDataObject()
	t_int.SetValue(200)
	if _, err := t_int.Ge("a"); err == nil {
		t.Error("IntData Ge error!")
	}
}
