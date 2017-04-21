# unicontract

### Quick Start
1. sudo apt-get install git
2. mkdir -p go/src && cd go/src
2. git clone https://git.oschina.net/uni-ledger/unicontract.git
4. cd unicontract &&./build.sh init && . ~/.bashrc
5. go run main.go

### Project Contents
```
bin:
pkg:
src:
  |__*unicontract
       |__ conf    配置文件
       |__ demo:   调用使用样例
       |__ src
            |__ api:    合约执行层对合约应用层提供的API
            |__ chain:  合约执行层与区块链层交互接口
            |__ commands:命令行支持
            |__ common: 公用工具集包
            |__ config: 配置文件工具包
            |__ core:   合约执行层核心处理逻辑******
                 |__ conf    合约配置文件
                 |__ control 合约进程控制等
                 |__ db      合约db处理
                 |__ model   合约模型处理
                 |__ protos  合约接口描述proto
            |__ pipelines: 流水线管道
       |__ tools:  语法检查、逻辑检查等工具
       |__ verdor: 第三方包管理工具 govendor
```

### Output Package
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

### Links for Developers
- [Wiki](http://git.oschina.net/uni-ledger/unicontract/wikis/home)
- [Open issues](http://git.oschina.net/uni-ledger/unicontract/issues)
- [Useful links](http://git.oschina.net/uni-ledger/unicontract/issues/6)
- [Gitter chatroom](https://gitter.im/uni-ledger/unicontract)