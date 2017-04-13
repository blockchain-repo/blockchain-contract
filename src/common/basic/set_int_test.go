package basic

import (
	"fmt"
	"testing"
	//"sort"
	//"unicontract/src/common"
)

func Test_SetInt(t *testing.T) {
	//初始化
	s := NewSetInt()

	s.Add(1)
	s.Add(1)
	s.Add(0)
	s.Add(2)
	s.Add(4)
	s.Add(3)

	fmt.Println(s.List())
	//fmt.Println(s.SortList())
	s.Clear()
	if s.IsEmpty() {
		fmt.Println("0 item")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 exist in item")
	}

	s.Remove(2)
	//s.Remove(3)
	fmt.Println("无序的切片", s.List())
	//fmt.Println("有序的切片", s.SortList())
	fmt.Println("无序的切片", s.List())
}
