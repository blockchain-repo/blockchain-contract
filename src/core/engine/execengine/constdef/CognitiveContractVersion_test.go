package constdef

import "testing"

func TestContratVersion(t *testing.T) {
	var version string = UCVM_Version
	var copyright = UCVM_CopyRight
	var date = UCVM_Date
	if version != "v1.0" {
		t.Error("CUVM_Version value Error!")
	}
	if copyright != "uni-ledger" {
		t.Error("CUVM_CopyRight value Error!")
	}
	if date != "2017-06-01 12:00:00" {
		t.Error("CUVM_Date value Error!")
	}
}
