# unicontract


### Quick Start
- http://git.oschina.net/uni-ledger/unicontract/issues/7

### Contents
```
bin:
pkg:
src:
    |__unicontract
        |__conf/app.conf  api configure file
        |__src
            |__ api:    合约执行层对合约应用层提供的API
                |__...
                ......
            |__ chain:  合约执行层与区块链层交互接口
            |__ commands:命令行支持
            |__ common: 公用工具集包
            |__ core:   合约执行层核心处理逻辑******
                |__ conf     合约配置文件
                |__ control  合约进程控制等
                |__ db       合约db处理
                |__ model    合约模型处理
                |__ protos   合约接口描述proto
            |__ demo:       调用使用样例
            |__ unittest:   单元测试、集成测试等
            |__ tools:      语法检查、逻辑检查等工具
```

### Links
- http://git.oschina.net/uni-ledger/unicontract/issues/6