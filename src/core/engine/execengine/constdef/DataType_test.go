package constdef

import (
	"testing"
)

func TestDataType(t *testing.T){
	var ctype0 int = Data_Unknown
	if ctype0 != 0{
		t.Error("Data_Unknown value Error!")
	}
	if DataType[Data_Unknown] != "Data_Unknown" {
		t.Error("Data_Unknown string value Error!")
	}
	var ctype1 int = Data_Numeric_Uint
	if ctype1 != 1{
		t.Error("Data_Numeric_Uint value Error!")
	}
	if DataType[Data_Numeric_Uint] != "Data_Numeric_Uint" {
		t.Error("Data_Numeric_Uint string value Error!")
	}
	var ctype2 int = Data_Numeric_Int
	if ctype2 != 2{
		t.Error("Data_Numeric_Int value Error!")
	}
	if DataType[Data_Numeric_Int] != "Data_Numeric_Int" {
		t.Error("Data_Numeric_Int string value Error!")
	}
	var ctype3 int = Data_Numeric_Float
	if ctype3 != 3{
		t.Error("Data_Numeric_Float value Error!")
	}
	if DataType[Data_Numeric_Float] != "Data_Numeric_Float" {
		t.Error("Data_Numeric_Float string value Error!")
	}
	var ctype4 int = Data_Text
	if ctype4 != 4{
		t.Error("Data_Text value Error!")
	}
	if DataType[Data_Text] != "Data_Text" {
		t.Error("Data_Text string value Error!")
	}
	var ctype5 int = Data_Date
	if ctype5 != 5{
		t.Error("Data_Date value Error!")
	}
	if DataType[Data_Date] != "Data_Date" {
		t.Error("Data_Date string value Error!")
	}
	var ctype6 int = Data_Array
	if ctype6 != 6{
		t.Error("Data_Array value Error!")
	}
	if DataType[Data_Array] != "Data_Array" {
		t.Error("Data_Array string value Error!")
	}
	var ctype7 int = Data_Compound
	if ctype7 != 7{
		t.Error("Data_Compound value Error!")
	}
	if DataType[Data_Compound] != "Data_Compound" {
		t.Error("Data_Compound string value Error!")
	}
	var ctype8 int = Data_Matrix
	if ctype8 != 8{
		t.Error("Data_Matrix value Error!")
	}
	if DataType[Data_Matrix] != "Data_Matrix" {
		t.Error("Data_Matrix string value Error!")
	}
	var ctype9 int = Data_DecisionCandidate
	if ctype9 != 9{
		t.Error("Data_DecisionCandidate value Error!")
	}
	if DataType[Data_DecisionCandidate] != "Data_DecisionCandidate" {
		t.Error("Data_DecisionCandidate string value Error!")
	}
}