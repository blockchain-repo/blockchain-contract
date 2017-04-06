package protoc3

import (
	"testing"
	"fmt"
)

func TestCreatePerson(t *testing.T) {

	p := CreatePerson()
	fmt.Println(p)
}

func TestMarshal(t *testing.T) {
	Marshal(CreatePerson())
}

func TestUnMarshal(t *testing.T) {
	UnMarshal()
}

func TestEqueMessage(t *testing.T) {
	EqueMessage()
}

func TestFileDescriptor(t *testing.T) {
	FileDescriptor()
}

func TestFloat32(t *testing.T) {
	Float32()
}

func TestMessageName(t *testing.T) {
	MessageName()
}

func TestSize(t *testing.T) {
	Size()
}

func TestMarshalText(t *testing.T) {
	MarshalText()
}

