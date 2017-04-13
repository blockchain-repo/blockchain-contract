package basic

// base Set
type Set interface {
	Add(e interface{}) bool      // 添加元素值
	Remove(e interface{})        // 删除元素值
	Clear()                      // 清除所有元素值
	Contains(e interface{}) bool // 判断是否包含某个元素值
	Has(e interface{}) bool      // 判断是否包含某个元素值
	Len() int                    // 获取元素值的数量
	Same(other Set) bool         // 判断与其他Set类型值是否相同
	Elements() []interface{}     // 获取所有元素值，即生成可迭代的快照
	String() string              // 获取自身的字符串表示形式
}


// 判断集合 one 是否是集合 other 的超集
// 读者应重点关注IsSuperset与附属于HashSet类型的IsSuperset方法的区别
func IsSuperSet(one Set, other Set) bool {
	if one == nil || other == nil {
		return false
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if !one.Contains(v) {
			return false
		}
	}
	return true
}