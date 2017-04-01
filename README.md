
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
            |__ conf    合约配置文件
            |__ control 合约进程控制等
            |__ db      合约db处理
            |__ model   合约模型处理
            |__ protos  合约接口描述proto
       |__ demo:     调用使用样例
       |__ unittest: 单元测试、集成测试等
       |__ tools:    语法检查、逻辑检查等工具
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