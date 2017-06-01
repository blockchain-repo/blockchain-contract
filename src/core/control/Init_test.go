// Init_test
package control

import (
	"testing"
)

func Test_Init(t *testing.T) {
	t.Logf("%+v\n", Conf)
	Init()
	t.Logf("%+v\n", Conf)
}
