package transaction

import (
	"github.com/astaxie/beego/logs"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	//data := strings
	//a := StringSlice(data[0:])
	//Sort(a)
	//var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

	var silce = []string{"a", "c", "b", "A", "C", "B", "1", "3", "2"}
	sort.Strings(silce)
	//sliceSort := sort.StringSlice(silce)
	logs.Info(silce)
	//logs.Info(sliceSort)
}

func TestTransferAssetComplete(t *testing.T) {
	//TransferAssetComplete("", ",")
}
