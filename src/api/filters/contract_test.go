package filters


import (
	"fmt"
	"testing"
	"bytes"
	"net/http"
	"io/ioutil"
)

func Test_ContractFilter(t *testing.T) {
	client := &http.Client{}
	url := "http://192.168.1.14:8088/v1/contract/signature"
	body_byte :=[]byte{10 ,1, 50, 48, 209, 249, 171, 199, 5}
	req_body := bytes.NewReader(body_byte)

	req, err := http.NewRequest("POST", url, req_body)
	if err != nil {
		// handle error
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Content-Type", "application/octet-stream")
	//req.Header.Set("RequestDataType", "proto")
	//req.Header.Set("RequestDataType", "json")
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("RequestDataType", "proto")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}