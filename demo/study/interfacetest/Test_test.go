package interfacetest

import "testing"

func TestInterface(t *testing.T){
	var aa ITest = &Test{name:"testname", ttype:100}
	aa.GetName()

	var bb *Test = aa.(*Test)
	bb.GetTtype()
}
