## unicontract
### Pre

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

## start
```
bash build.sh
```