package data

import (
	"fmt"
	"math"
	"testing"
	"unicontract/src/core/engine/execengine/constdef"
)

func CreateFloatDataObject() *FloatData {
	t_int := new(FloatData)
	t_int.InitFloatData()
	t_int.SetCname("TestMoney")
	t_int.SetDefaultValue(0.0)
	t_int.SetCaption("money")
	t_int.SetDescription("money of china")
	t_int.SetUnit("元")

	return t_int
}

func Testfloat64Init(t *testing.T) {
	t_int := new(FloatData)
	t_int.InitFloatData()
	t_int.SetCname("TestMoney")
	t_int.SetDefaultValue(0.0)
	t_int.SetCaption("money")
	t_int.SetDescription("money of china")
	t_int.SetUnit("元")
	if t_int == nil {
		t.Error("FloatData init Error!")
	}
	if t_int.GetCtype() != constdef.ComponentType[constdef.Component_Data]+"."+constdef.DataType[constdef.Data_Numeric_Float] {
		t.Error("t_ctype value Error!")
	}
	if t_int.GetCname() != "TestMoney" {
		t.Error("t_name value Error!")
	}
	if t_int.GetDefaultValue() != 0.0 {
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
	if t_int.GetHardConvType() != "float64" {
		t.Error("hardConvType value Error!")
	}
	if t_int.GetDataRange()[0] != -math.MaxFloat64 {
		t.Error("dataRange left value Error!")
	}
	if t_int.GetDataRange()[1] != math.MaxFloat64 {
		t.Error("dataRange right value Error!")
	}
	fmt.Println(t_int.GetCname(), " ", t_int.GetDefaultValue(), " ", t_int.GetUnit(), " ", t_int.GetCaption(), " ", t_int.GetDescription())
}

func TestFloatDataRange(t *testing.T) {
	t_int := CreateFloatDataObject()
	var t_range_1 [2]float64 = [2]float64{0.0, 0.0}
	var t_range_2 [2]float64 = [2]float64{0.0, 1.0}
	var t_range_3 [2]float64 = [2]float64{1.0, 1.0}
	var t_range_4 [2]float64 = [2]float64{1.0, 0.0}
	var t_range_5 [2]float64 = [2]float64{1.0, -1.0}
	var t_range_6 [2]float64 = [2]float64{-1.0, -1.0}
	//default [0, 0]
	var t_range_7 [2]float64
	if err := t_int.SetDataRange(t_range_1); err != nil {
		t.Error("[0,0] process error")
	}
	if err := t_int.SetDataRange(t_range_2); err != nil {
		t.Error("[0,1] process error")
	}
	if err := t_int.SetDataRange(t_range_3); err != nil {
		t.Error("[1,1] process error")
	}
	if err := t_int.SetDataRange(t_range_4); err == nil {
		t.Error("[1,0] process error")
	}
	if err := t_int.SetDataRange(t_range_5); err == nil {
		t.Error("[1,-1] process error")
	}
	if err := t_int.SetDataRange(t_range_6); err != nil {
		t.Error("[-1,-1] process error")
	}
	if err := t_int.SetDataRange(t_range_7); err != nil {
		t.Error("[] process error")
	} else {
		if t_int.GetDataRange()[0] == 0 {
			t.Error("dataRange left value Default Error!")
		}
	}
}

func Testfloat64CheckRange(t *testing.T) {
	t_int := CreateFloatDataObject()
	var t_range [2]float64 = [2]float64{0, 100}
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

func Testfloat64Add(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(100))
	if v_int, err := t_int.Add(float64(200)); err == nil {
		if v_int != float64(300) {
			t.Error("FloatData add error,result not equal 300!")
		} else {
			fmt.Println("100 + 200 = ", v_int)
		}
	} else {
		t.Error("FloatData add error!")
	}
}

func Testfloat64RAdd(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(100))
	if v_int, err := t_int.RAdd(float64(200)); err == nil {
		if v_int != float64(300) {
			t.Error("FloatData add error,result not equal 300!")
		} else {
			fmt.Println("200 + 100 = ", v_int)
		}
	} else {
		t.Error("FloatData add error!")
	}
}

func Testfloat64Sub(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(400))
	if v_int, err := t_int.Sub(float64(200)); err == nil {
		if v_int != float64(200) {
			t.Error("FloatData Sub error,result not equal 200!")
		} else {
			fmt.Println("400 - 200 = ", v_int)
		}
	} else {
		t.Error("FloatData Sub error!")
	}
}

func Testfloat64RSub(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(100))
	if v_int, err := t_int.RSub(float64(300)); err == nil {
		if v_int != float64(200) {
			t.Error("FloatData RSub error,result not equal 200!")
		} else {
			fmt.Println("300 - 100 = ", v_int)
		}
	} else {
		t.Error("FloatData RSub error!")
	}
}

func Testfloat64Mul(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(4))
	if v_int, err := t_int.Mul(float64(2)); err == nil {
		if v_int != float64(8) {
			t.Error("FloatData Mul error,result not equal 8!")
		} else {
			fmt.Println("4 * 2 = ", v_int)
		}
	} else {
		t.Error("FloatData Mul error!")
	}
}

func Testfloat64RMul(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(2))
	if v_int, err := t_int.RMul(float64(4)); err == nil {
		if v_int != float64(8) {
			t.Error("FloatData RMul error,result not equal 8!")
		} else {
			fmt.Println("2 * 4 = ", v_int)
		}
	} else {
		t.Error("FloatData RMul error!")
	}
}

func Testfloat64Div(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(4))
	if v_int, err := t_int.Div(float64(2)); err == nil {
		if v_int != float64(2) {
			t.Error("FloatData Div error,result not equal 8!")
		} else {
			fmt.Println("4 / 2 = ", v_int)
		}
	} else {
		t.Error("FloatData Div error!")
	}
}

func Testfloat64DivError(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(4))
	if _, err := t_int.Div(float64(0)); err == nil {
		fmt.Println("4 / 0 = error")
		t.Error("FloatData Div error, zero exist!")
	}
}

func Testfloat64RDiv(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(2))
	if v_int, err := t_int.RDiv(float64(4)); err == nil {
		if v_int != float64(2) {
			t.Error("FloatData RDiv error,result not equal 8!")
		} else {
			fmt.Println("4 / 2 = ", v_int)
		}
	} else {
		t.Error("FloatData RDiv error!")
	}
}

func Testfloat64Lt(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(100))
	if v_bool, err := t_int.Lt(float64(200)); err == nil {
		if !v_bool {
			t.Error("FloatData Lt error,result not equal true!")
		} else {
			fmt.Println("100 < 200: ", v_bool)
		}
	} else {
		t.Error("FloatData Lt error!")
	}
}

func Testfloat64LtFail(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(200))
	if v_bool, err := t_int.Lt(float64(100)); err == nil {
		if v_bool {
			t.Error("FloatData Lt error,result not equal true!")
		} else {
			fmt.Println("200 < 100: ", v_bool)
		}
	} else {
		t.Error("FloatData Lt error!")
	}
}

func Testfloat64LtError(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(200))
	if _, err := t_int.Lt("a"); err == nil {
		t.Error("FloatData Lt error!")
	}
}

func Testfloat64Le(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(100))
	if v_bool, err := t_int.Le(float64(100)); err == nil {
		if !v_bool {
			t.Error("FloatData Le error,result not equal true!")
		} else {
			fmt.Println("100 <= 100: ", v_bool)
		}
	} else {
		t.Error("FloatData Le error!")
	}
}

func Testfloat64LeFail(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(200))
	if v_bool, err := t_int.Le(float64(100)); err == nil {
		if v_bool {
			t.Error("FloatData Le error,result not equal true!")
		} else {
			fmt.Println("200 < 100: ", v_bool)
		}
	} else {
		t.Error("FloatData Le error!")
	}
}

func Testfloat64LeError(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(200))
	if _, err := t_int.Le("a"); err == nil {
		t.Error("FloatData Lt error!")
	}
}

func Testfloat64Gt(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(200))
	if v_bool, err := t_int.Gt(float64(100)); err == nil {
		if !v_bool {
			t.Error("FloatData Gt error,result not equal true!")
		} else {
			fmt.Println("200 > 100: ", v_bool)
		}
	} else {
		t.Error("FloatData Gt error!")
	}
}

func Testfloat64GtFail(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(100))
	if v_bool, err := t_int.Gt(float64(200)); err == nil {
		if v_bool {
			t.Error("FloatData Gt error,result not equal true!")
		} else {
			fmt.Println("100 > 200: ", v_bool)
		}
	} else {
		t.Error("FloatData Gt error!")
	}
}

func Testfloat64GtError(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(200))
	if _, err := t_int.Gt("a"); err == nil {
		t.Error("FloatData Gt error!")
	}
}

func Testfloat64Ge(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(200))
	if v_bool, err := t_int.Ge(float64(200)); err == nil {
		if !v_bool {
			t.Error("FloatData Ge error,result not equal true!")
		} else {
			fmt.Println("200 > 200: ", v_bool)
		}
	} else {
		t.Error("FloatData Ge error!")
	}
}

func Testfloat64GeFail(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(100))
	if v_bool, err := t_int.Ge(float64(200)); err == nil {
		if v_bool {
			t.Error("FloatData Ge error,result not equal true!")
		} else {
			fmt.Println("100 > 200: ", v_bool)
		}
	} else {
		t.Error("FloatData Ge error!")
	}
}

func Testfloat64GeError(t *testing.T) {
	t_int := CreateFloatDataObject()
	t_int.SetValue(float64(200))
	if _, err := t_int.Ge("a"); err == nil {
		t.Error("FloatData Ge error!")
	}
}
