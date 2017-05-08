package data

import (
	"testing"
	"fmt"
	"time"
)

func CreateDateDataObject() *DateData{
	t_int := new(DateData)
	t_int.InitDateData()
	t_int.SetCname("TestDate")
	t_int.SetCaption("date")
	t_int.SetDescription("Date Test")
	t_int.SetUnit("")

	return t_int
}

func TestDateInit(t *testing.T){
	t_int := new(DateData)
	t_int.InitDateData()
	t_int.SetCname("TestDate")
	t_int.SetCaption("date")
	t_int.SetDescription("Date Test")
	t_int.SetUnit("")
	if t_int == nil {
		t.Error("DateData init Error!")
	}
	if t_int.GetCname() != "TestDate" {
		t.Error("t_name value Error!")
	}
	if t_int.GetUnit() != "" {
		t.Error("t_unit value Error!")
	}
	if t_int.GetCaption() != "date" {
		t.Error("t_caption value Error!")
	}
	if t_int.GetDescription() != "Date Test" {
		t.Error("t_description value Error!")
	}
	if t_int.GetHardConvType() != "strToDate" {
		t.Error("hardConvType value Error!")
	}
	if t_int.GetFormat() != "2006-01-02 15:04:05" {
		t.Error("format value error!")
	}

	fmt.Println(t_int.GetCname(), " ", t_int.GetDefaultValue(), " ", t_int.GetFormat(), " ", t_int.GetCaption(), " ", t_int.GetDescription())
}

func TestFormatOP(t *testing.T){
	t_date := CreateDateDataObject()
	var v_format string = "2017-12-01 01:01:30"
	t_date.SetFormat(v_format)
	if t_date.GetFormat() != "2017-12-01 01:01:30" {
		t.Error("DateData SetFormat Error!")
	}
}

func TestStrToDateNull(t *testing.T){
	t_date := CreateDateDataObject()
	var now_date_miao int64 = time.Now().Unix()
	var not_date_formaat string = time.Unix(now_date_miao, 0).Format(t_date.GetFormat())
	if _,err := t_date.strToDate(not_date_formaat);err != nil {
		t.Error("StrToDate(param is null) Error!")
	}
	if _,err := t_date.strToDate("");err == nil {
		t.Error("StrToDate(param is null) Error!")
	}
}

func TestStrToDateOK(t *testing.T){
	t_date := CreateDateDataObject()
	var v_date_int int64 = 1491927053
	var v_date_format string = "2017-04-11 16:10:53"
	if v_time,err := t_date.strToDate(v_date_format);err != nil{
		t.Error("StrToDate Error!")
	} else {
		if v_time.Unix() != v_date_int {
			t.Error("StrToDate Error!")
		}
		if v_time.Format(t_date.GetFormat()) != v_date_format {
			t.Error("StrToDate Error!")
		}
	}
}

func TestGetValueInt(t *testing.T){
	t_date := CreateDateDataObject()
	var v_date_int int64 = 1491927053
	var v_date_format string = "2017-04-11 16:10:53"
	t_date.SetValue(v_date_format)
	if t_date.GetValueFormat() != v_date_format {
		t.Error("GetValueFormat Error!")
	}
	var r_date int64
	var r_err error
	r_date,r_err = t_date.GetValueInt()
	if r_date != v_date_int {
		t.Error("GetValueInt Error(" +  r_err.Error() + ")!")
	}
}

func TestDateAdd(t *testing.T){
	t_date := CreateDateDataObject()
	var v_date_format string = "2017-04-11 16:10:53"
	t_date.SetValue(v_date_format)
    var v_date_add time.Time
	var v_err error
	v_date_add,v_err = t_date.Add(1)
	if v_err != nil {
		t.Error("Date Add Error!")
	} else if v_date_add.Format(t_date.GetFormat()) != "2017-04-12 16:10:53" {
		t.Error("Date Add fail!")
	}

}

func TestDateLt(t *testing.T){
	t_date := CreateDateDataObject()
	var v_date_int int64 = 1491927053
	var v_date_format string = "2017-04-11 16:10:53"
	t_date.SetValue(v_date_format)
	if t_date.GetValueFormat() != v_date_format {
		t.Error("GetValueFormat Error!")
	}
	var r_date int64
	var r_err error
	r_date,r_err = t_date.GetValueInt()
	if r_date != v_date_int {
		t.Error("GetValueInt Error(" +  r_err.Error() + ")!")
	}
}

func TestDateGt(t *testing.T){
	t_date := CreateDateDataObject()
	var v_date_int int64 = 1491927053
	var v_date_format string = "2017-04-11 16:10:53"
	t_date.SetValue(v_date_format)
	if t_date.GetValueFormat() != v_date_format {
		t.Error("GetValueFormat Error!")
	}
	var r_date int64
	var r_err error
	r_date,r_err = t_date.GetValueInt()
	if r_date != v_date_int {
		t.Error("GetValueInt Error(" +  r_err.Error() + ")!")
	}
}

