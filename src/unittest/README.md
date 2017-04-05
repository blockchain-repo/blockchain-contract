## 单元测试规范
```
1. 目录同src下的源码目录保持一致
2. 文件名以被测试的源文件名为前缀, 格式为: xxxxx_test.go
3. 函数名以被测试的原函数名为后缀，格式为: func Testxxxx(t *testing.T)
4. 单测运行方法： go test
```

### 单元测试样例
```
源方法:
func Max(x int, y int) int{
    if x >= y {
        return x
    }
    return y
}

单测Demo:
package test

import "testing"

func TestMax(t *testing.T){
    const x,y = 2,3
    max := Max(x,y)
    if max == x {
        t.Error("Max Error!")
    }
}

运行结果：
PASS
ok      uniledger.com/unicontract/unittest  0.003s
```
