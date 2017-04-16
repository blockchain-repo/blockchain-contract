// request
package request

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type RequestParam struct {
	Method       string // POST/GET
	URL          string
	ParamJSON    string
	ResponseJSON string
}

/*
 * request 方法
 * \param [in, out] requestParam : request的相关参数，以及response返回的信息
 * \return nil 正常
 * \       not nil 有错误
 */
func Request(requestParam *RequestParam) error {
	method := strings.ToUpper(requestParam.Method)
	requestParam.Method = method
	if method == "GET" {
		return _Get(requestParam)
	} else if method == "POST" {
		return _Post(requestParam)
	} else {
		return errors.New(fmt.Sprintf("%s is unsupported method", requestParam.Method))
	}
}

func _Get(requestParam *RequestParam) error {
	resp, err := http.Get(requestParam.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	requestParam.ResponseJSON = string(body)
	return err
}

func _Post(requestParam *RequestParam) error {
	req, err := http.NewRequest(requestParam.Method,
		requestParam.URL,
		strings.NewReader(requestParam.ParamJSON))
	if err != nil {
		return err
	}

	// Body Type
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	// 完成后断开连接
	//req.Header.Set("Connection", "close")

	DefaultClient := http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*5)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(5 * time.Second))
				return c, nil
			},
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	resp, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	requestParam.ResponseJSON = string(body)
	return err
}
