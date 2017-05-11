
-----------------------------------
proto及model 修改记录 2017-05-11
-----------------------------------
### proto 新增字段流程
#### 1. 字段修改
```
contract.proto

ContractHead:
    增加Timestamp       //合约执行时间戳    string
ContractBody:
    增加MetaAttribute  //合约属性MetaData  map[string]interface{} ....bytes
ContractBody.ContractComponents：
    增加TaskExecuteIdx //任务执行索引次数 int


Relation:
transaction.relation:
    增加：
      “ContractHashId”:”xxxxxxxxxxxxxxxxxxx”,   #合约HashID(原ContractId中的内容)   string
      “TaskExecuteIdx”:0,                                          #合约任务执行索引   int

    修改：
    “ContractId”: “xxxxxxxxx”,                             #合约ID string（存储内容改为第一次创建合约，合约描述态的ID）
```
#### 2. 根据 proto 生成 go文件
```
# unicontract/ 根目录下执行
bash protoc-tools.sh go java js python
```
--------------------------------------