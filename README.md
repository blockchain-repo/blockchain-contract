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
            |__ pipelines:  流水线管道
        |__tools   语法检查、逻辑检查等工具
        |__vendor  第三方包管理工具 govendor
```

### Links
- http://git.oschina.net/uni-ledger/unicontract/issues/6

### build.sh
使用build脚本中init参数时，一定要用如下方法：   
./build.sh init && . ~/.bashrc
