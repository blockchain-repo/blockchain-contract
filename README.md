
## 目录结构描述
```
bin:
pkg:
src:
  |__unicontract
       |__ api:    合约执行层对合约应用层提供的API
       |__ chain:  合约执行层与区块链层交互接口
       |__ commands:命令行支持
       |__ common: 公用工具集包
       |__ core:   合约执行层核心处理逻辑******
            |__
            |__
       |__ demo:   调用使用样例
       |__ tools:  语法检查、逻辑检查等工具
```

## 产出目录结构
```
unicontract:
  |__bin:  产出bin文件
  |__conf: conf配置文件
  |__data: 依赖数据文件
  |__log:  过程log文件,按日期切割
      |__ unicontract.log
      |__ unicontract.log.wf
      |__ unicontract.log.debug
```
