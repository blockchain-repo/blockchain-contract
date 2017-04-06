
package protoc3

import (
	"github.com/golang/protobuf/proto"
	logger "common/seelog"
	uniprotoc "common/protoc3"
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"reflect"
)

func CreatePerson() proto.Message{


	p := &uniprotoc.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*uniprotoc.Person_PhoneNumber{
			{Number: "555-4321", Type: uniprotoc.Person_HOME},
		},
	}

	p1 := &uniprotoc.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*uniprotoc.Person_PhoneNumber{
			{Number: "555-4321", Type: uniprotoc.Person_HOME},
		},
	}
	fmt.Println(proto.Equal(p,p1))

	fmt.Println("==========",p.Email)
	p.Reset()
	fmt.Println()
	fmt.Println("==========",p.Email)


	a := &uniprotoc.AddressBook{
		[]*uniprotoc.Person{p,p},
	}
	fmt.Println(a)
	return p
}

func EqueMessage(){
	p := &uniprotoc.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*uniprotoc.Person_PhoneNumber{
			{Number: "555-4321", Type: uniprotoc.Person_HOME},
		},
	}

	p1 := &uniprotoc.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*uniprotoc.Person_PhoneNumber{
			{Number: "555-4321", Type: uniprotoc.Person_HOME},
		},
	}
	proto.MarshalText(os.Stdout,p)
	fmt.Println(proto.MarshalTextString(p))
	fmt.Println(proto.Equal(p,p1))
}

func FileDescriptor(){
	b := proto.FileDescriptor("addressbook.proto")
	proto.RegisterFile("aa.txt",b)
	fmt.Println(string(b))
	fmt.Println(b)
}

func MessageName(){
	p1 := &uniprotoc.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*uniprotoc.Person_PhoneNumber{
			{Number: "555-4321", Type: uniprotoc.Person_HOME},
		},
	}
	p1.Email = "aaaaaaaaaaaaaaaaaa"
	fmt.Println(p1)
	fmt.Println(proto.MessageName(p1))
}

func Float32(){
	f := proto.Float32(0.32)
	fmt.Println(f)
	fmt.Println(*f)

}

func Size(){
	fmt.Println(proto.Size(CreatePerson()))
	p1 := &uniprotoc.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*uniprotoc.Person_PhoneNumber{
			{Number: "555-4321", Type: uniprotoc.Person_HOME},
		},
	}
	fmt.Println(proto.Size(p1))
}

func Marshal(pb proto.Message) []byte{

	out,err := proto.Marshal(pb)

	if err != nil{
		logger.Error("Failed to encode address book:",err)
	}
	if err := ioutil.WriteFile("aa.txt", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	return out
}

func UnMarshal() proto.Message{
	in,err := ioutil.ReadFile("aa.txt")
	if err != nil{
		log.Fatalln("Error reading file:", err)
	}
	p := &uniprotoc.Person{}
	if err := proto.Unmarshal(in,p);err != nil{
		log.Fatalln("Failed to parse address book:", err)
	}
	return p
}

func MarshalText(){
	p1 := &uniprotoc.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*uniprotoc.Person_PhoneNumber{
			{Number: "555-4321", Type: uniprotoc.Person_HOME},
		},
	}
	a:=proto.MarshalTextString(p1)
	fmt.Println(p1)
	fmt.Println(reflect.TypeOf(p1))
	fmt.Println(reflect.TypeOf(a))
	fmt.Println("=====",a)
	proto.UnmarshalText(a,p1)
	fmt.Println(p1)
	fmt.Println(reflect.TypeOf(p1))
}




















