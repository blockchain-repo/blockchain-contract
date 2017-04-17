## protoc3使用规则
1. 下载安装protobuf
```
wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip -c -P /opt

```
2. 配置环境变量
```
export PATH=$PATH:/opt/protoc/bin
source /etc/profile

```
3. 安装Go protocl buffers plugin
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
4. 创建.proto文件
```
syntax = "proto3";
package tutorial;
message Person {
  string name = 1;
  int32 id = 2;  // Unique ID number for this person.
  string email = 3;

  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }

  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }

  repeated PhoneNumber phones = 4;
}

// Our address book file is just one of these.
message AddressBook {
  repeated Person people = 1;
}
```
5. 运行编译器生成go文件
```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```
6. 使用protocol buffer api处理message
```
func CreatePerson() proto.Message{

	p := &uniprotoc.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*uniprotoc.Person_PhoneNumber{
			{Number: "555-4321", Type: uniprotoc.Person_HOME},
		},
	}

	return p
}
```











