// request_test
package request

import (
	"testing"
)

var requestParam RequestParam

func Test_Request_post(t *testing.T) {
	requestParam.Method = "POST"
	if err := Request(&requestParam); err != nil {
		t.Errorf("Test_Request_post is failed, err is %v\n", err)
	} else {
		t.Log("Test_Request_post is pass.")
	}
}

func Test_Request_get(t *testing.T) {
	requestParam.Method = "GET"
	if err := Request(&requestParam); err != nil {
		t.Errorf("Test_Request_get is failed, err is %v\n", err)
	} else {
		t.Log("Test_Request_get is pass.")
	}
}

func Test_Request_other(t *testing.T) {
	requestParam.Method = "PUT"
	if err := Request(&requestParam); err == nil {
		t.Errorf("Test_Request_other is failed, err is %v\n", err)
	} else {
		t.Log("Test_Request_other is pass.")
	}
}
