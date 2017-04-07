## unicontract

### 目录结构描述
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

### Golang env
#### 1. Download
```
sudo wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz -c -P /opt

sudo tar -C /usr/local/ -xzf /opt/go1.8.linux-amd64.tar.gz

```
#### 2. Configure
创建`go workspace`

```
mkdir -p $HOME/work/golang
```

在 `/etc/profile` 或者 `$HOME/.profile` 中添加以下几行:
```
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

export GOPATH=$HOME/work/golang

```

如果想在任意目录执行生成的文件，则可以添加`$GOPATH/bin`到`$PATH`
```
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```
然后, 执行 `source /etc/profile` 或者 `source ~/.profile`

检查是否生效
```
go version
```

### contract
#### download
```
git clone https://git.oschina.net/uni-ledger/unicontract.git $GOPATH/src/
```

#### start
```
bash build.sh
```

#### model generate(through the `*.proto`)
Now `*.proto` files are in `src/core/protos/`, the generate `*.go` should be in
`src/core/model/`.

生成 `*.proto` 对应的 `.go` 文件
```
bash proto-gen-go.sh init
```
如果生成的`.go` 文件跟`.proto` 文件不在同一目录下，则需要修改生成的 .go 文件的 package信息