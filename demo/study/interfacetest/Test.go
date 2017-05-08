package interfacetest

import "fmt"

type Test struct {
	name string
	ttype int
}

func (t Test) GetName()string {
	fmt.Println("Name: ", t.name)
	return t.name
}

func (t *Test) GetTtype() int{
	fmt.Println("Type: ", t.ttype)
	return t.ttype
}
