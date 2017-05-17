package study

import (
	"testing"
	"fmt"
)

func AddString(p_map map[string]interface{}, p_strname string, p_strvalue string){
	if p_map == nil {
		p_map = make(map[string]interface{}, 0)
	}
	p_map[p_strname] = p_strvalue
}

func TestMap(t *testing.T){
	var v_map map[string]interface{} = make(map[string]interface{}, 0)
	AddString(v_map, "a", "aaaaaaaaa")
	AddString(v_map, "b", "bbbbbbbbb")
	AddString(v_map, "c", "ccccccccc")
	for v_key,v_value := range v_map {
		fmt.Println(v_key, ":", v_value)
	}

	//传递为空时，无法进行赋值，没有对应指针
	var t_map map[string]interface{}
	AddString(t_map, "1", "1111111111")
	for v_key,v_value := range t_map {
		fmt.Println(v_key, ":", v_value)
	}
	//验证map len
	var m_map map[string]interface{}
	fmt.Println(m_map == nil , len(m_map))
	var m2_map map[string]interface{} = make(map[string]interface{}, 0)
	fmt.Println(m2_map == nil , len(m2_map))
}



