package jsontest

type BB struct {
	AA              `json:"a"`
	Description string `json:"description"`
}

type B struct {
	A
	Description string
}
