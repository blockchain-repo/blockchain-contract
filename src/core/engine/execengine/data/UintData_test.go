package data

import (
	"fmt"
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
)

func CreateUintDataObject() *UintData {
	t_uint := new(UintData)
	t_uint.InitUintData()
	t_uint.SetCname("TestMoney")
	t_uint.SetDefaultValue(0)
	t_uint.SetCaption("money")
	t_uint.SetDescription("money of china")
	t_uint.SetUnit("元")

	return t_uint
}

func TestUintInit(t *testing.T) {
	t_uint := new(UintData)
	err := t_uint.InitUintData()
	if err != nil {
		t.Error("InitUintData Error!")
	}
	t_uint.SetCname("TestMoney")
	t_uint.SetDefaultValue(0)
	t_uint.SetCaption("money")
	t_uint.SetDescription("money of china")
	t_uint.SetUnit("元")
	if t_uint == nil {
		t.Error("UintData init Error!")
	}
	if t_uint.GetCtype() != constdef.ComponentType[constdef.Component_Data]+"."+constdef.DataType[constdef.Data_Numeric_Uint] {
		t.Error("t_ctype value Error!")
	}
	if t_uint.GetCname() != "TestMoney" {
		t.Error("t_name value Error!")
	}
	if t_uint.GetDefaultValue() != 0 {
		t.Error("t_default value Error!")
	}
	if t_uint.GetUnit() != "元" {
		t.Error("t_unit value Error!")
	}
	if t_uint.GetCaption() != "money" {
		t.Error("t_caption value Error!")
	}
	if t_uint.GetDescription() != "money of china" {
		t.Error("t_description value Error!")
	}
	if t_uint.GetHardConvType() != "uint" {
		t.Error("hardConvType value Error!")
	}
	if t_uint.GetDataRange()[0] != 0 {
		t.Error("dataRange left value Error!")
	}
	if t_uint.GetDataRange()[1] != 2147483647 {
		t.Error("dataRange right value Error!")
	}
	fmt.Println(t_uint.GetCname(), " ", t_uint.GetDefaultValue(), " ", t_uint.GetUnit(), " ", t_uint.GetCaption(), " ", t_uint.GetDescription())
}

func TestUintDataRange(t *testing.T) {
	t_uint := CreateUintDataObject()
	var t_range_1 [2]uint = [2]uint{0, 0}
	var t_range_2 [2]uint = [2]uint{0, 1}
	var t_range_3 [2]uint = [2]uint{1, 1}
	var t_range_4 [2]uint = [2]uint{1, 0}
	//default [0, 0]
	var t_range_5 [2]uint
	if err := t_uint.SetDataRange(t_range_1); err != nil {
		t.Error("[0,0] process error")
	}
	if err := t_uint.SetDataRange(t_range_2); err != nil {
		t.Error("[0,1] process error")
	}
	if err := t_uint.SetDataRange(t_range_3); err != nil {
		t.Error("[1,1] process error")
	}
	if err := t_uint.SetDataRange(t_range_4); err == nil {
		t.Error("[1,0] process error")
	}
	if err := t_uint.SetDataRange(t_range_5); err != nil {
		t.Error("[] process error")
	} else {
		if t_uint.GetDataRange()[0] != 0 {
			t.Error("dataRange left value Default Error!")
		}
	}
}

func TestUintCheckRange(t *testing.T) {
	t_uint := CreateUintDataObject()
	var t_range [2]uint = [2]uint{0, 100}
	t_uint.SetDataRange(t_range)
	if t_uint.CheckRange(10) != true {
		t.Error("check error, 10 must in [0, 100]")
	}
	if t_uint.CheckRange(101) != false {
		t.Error("check error, 101 not in [0, 100]")
	}
	if t_uint.CheckRange(0) != true {
		t.Error("check error, 0 must in [0, 100]")
	}
	if t_uint.CheckRange(100) != true {
		t.Error("check error, 100 must in [0, 100]")
	}
}

func TestUintAdd(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(100))
	if v_int, err := t_uint.Add(uint(200)); err == nil {
		if v_int != uint(300) {
			t.Error("UintData add error,result not equal 300!")
		} else {
			fmt.Println("100 + 200 = ", v_int)
		}
	} else {
		t.Error("UintData add error!")
	}
}

func TestUintRAdd(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(100))
	if v_int, err := t_uint.RAdd(uint(200)); err == nil {
		if v_int != uint(300) {
			t.Error("UintData add error,result not equal 300!")
		} else {
			fmt.Println("200 + 100 = ", v_int)
		}
	} else {
		t.Error("UintData add error!")
	}
}

func TestUintSub(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(400))
	if v_int, err := t_uint.Sub(uint(200)); err == nil {
		if v_int != uint(200) {
			t.Error("UintData Sub error,result not equal 200!")
		} else {
			fmt.Println("400 - 200 = ", v_int)
		}
	} else {
		t.Error("UintData Sub error!")
	}
}

func TestUintRSub(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(100))
	if v_int, err := t_uint.RSub(uint(300)); err == nil {
		if v_int != uint(200) {
			t.Error("UintData RSub error,result not equal 200!")
		} else {
			fmt.Println("300 - 100 = ", v_int)
		}
	} else {
		t.Error("UintData RSub error!")
	}
}

func TestUintMul(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(4))
	if v_int, err := t_uint.Mul(uint(2)); err == nil {
		if v_int != uint(8) {
			t.Error("UintData Mul error,result not equal 8!")
		} else {
			fmt.Println("4 * 2 = ", v_int)
		}
	} else {
		t.Error("UintData Mul error!")
	}
}

func TestUintRMul(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(2))
	if v_int, err := t_uint.RMul(uint(4)); err == nil {
		if v_int != uint(8) {
			t.Error("UintData RMul error,result not equal 8!")
		} else {
			fmt.Println("2 * 4 = ", v_int)
		}
	} else {
		t.Error("UintData RMul error!")
	}
}

func TestUintDiv(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(4))
	if v_int, err := t_uint.Div(uint(2)); err == nil {
		if v_int != uint(2) {
			t.Error("UintData Div error,result not equal 8!")
		} else {
			fmt.Println("4 / 2 = ", v_int)
		}
	} else {
		t.Error("UintData Div error!")
	}
}

func TestUintDivError(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(4))
	if _, err := t_uint.Div(uint(0)); err == nil {
		fmt.Println("4 / 0 = error")
		t.Error("UintData Div error, zero exist!")
	}
}

func TestUintRDiv(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(2))
	if v_int, err := t_uint.RDiv(uint(4)); err == nil {
		if v_int != uint(2) {
			t.Error("UintData RDiv error,result not equal 8!")
		} else {
			fmt.Println("4 / 2 = ", v_int)
		}
	} else {
		t.Error("UintData RDiv error!")
	}
}

func TestUintMod(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(5))
	if v_int, err := t_uint.Mod(uint(2)); err == nil {
		if v_int != uint(1) {
			t.Error("UintData Mod error,result not equal 1!")
		} else {
			fmt.Println("5 % 2 = ", v_int)
		}
	} else {
		t.Error("UintData Mod error!")
	}
}

func TestUintLt(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(100))
	if v_bool, err := t_uint.Lt(uint(200)); err == nil {
		if !v_bool {
			t.Error("UintData Lt error,result not equal true!")
		} else {
			fmt.Println("100 < 200: ", v_bool)
		}
	} else {
		t.Error("UintData Lt error!")
	}
}

func TestUintLtFail(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(200))
	if v_bool, err := t_uint.Lt(uint(100)); err == nil {
		if v_bool {
			t.Error("UintData Lt error,result not equal true!")
		} else {
			fmt.Println("200 < 100: ", v_bool)
		}
	} else {
		t.Error("UintData Lt error!")
	}
}

func TestUintLtError(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(200))
	if _, err := t_uint.Lt("a"); err == nil {
		t.Error("UintData Lt error!")
	}
}

func TestUintLe(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(100))
	if v_bool, err := t_uint.Le(uint(100)); err == nil {
		if !v_bool {
			t.Error("UintData Le error,result not equal true!")
		} else {
			fmt.Println("100 <= 100: ", v_bool)
		}
	} else {
		t.Error("UintData Le error!")
	}
}

func TestUintLeFail(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(200))
	if v_bool, err := t_uint.Le(uint(100)); err == nil {
		if v_bool {
			t.Error("UintData Le error,result not equal true!")
		} else {
			fmt.Println("200 < 100: ", v_bool)
		}
	} else {
		t.Error("UintData Le error!")
	}
}

func TestUintLeError(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(200))
	if _, err := t_uint.Le("a"); err == nil {
		t.Error("UintData Lt error!")
	}
}

func TestUintGt(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(200))
	if v_bool, err := t_uint.Gt(uint(100)); err == nil {
		if !v_bool {
			t.Error("UintData Gt error,result not equal true!")
		} else {
			fmt.Println("200 > 100: ", v_bool)
		}
	} else {
		t.Error("UintData Gt error!")
	}
}

func TestUintGtFail(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(100))
	if v_bool, err := t_uint.Gt(uint(200)); err == nil {
		if v_bool {
			t.Error("UintData Gt error,result not equal true!")
		} else {
			fmt.Println("100 > 200: ", v_bool)
		}
	} else {
		t.Error("UintData Gt error!")
	}
}

func TestUintGtError(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(200))
	if _, err := t_uint.Gt("a"); err == nil {
		t.Error("UintData Gt error!")
	}
}

func TestUintGe(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(200))
	if v_bool, err := t_uint.Ge(uint(200)); err == nil {
		if !v_bool {
			t.Error("UintData Ge error,result not equal true!")
		} else {
			fmt.Println("200 > 200: ", v_bool)
		}
	} else {
		t.Error("UintData Ge error!")
	}
}

func TestUintGeFail(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(100))
	if v_bool, err := t_uint.Ge(uint(200)); err == nil {
		if v_bool {
			t.Error("UintData Ge error,result not equal true!")
		} else {
			fmt.Println("100 > 200: ", v_bool)
		}
	} else {
		t.Error("UintData Ge error!")
	}
}

func TestUintGeError(t *testing.T) {
	t_uint := CreateUintDataObject()
	t_uint.SetValue(uint(200))
	if _, err := t_uint.Ge("a"); err == nil {
		t.Error("UintData Ge error!")
	}
}
