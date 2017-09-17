package api

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"strings"
	"time"
)

var ALL_TOKENS_MAP = map[string]string{
	//appId: accessKey
	"15b12db830bb402258d616e7e8c80830": "4akKiBhiFNgwQkqG33jK6U6WJWkAT5MNSxyWkJcgWS5s", //66B227778E75D76FDB02DEABB474551C
	//"e220674cf7205de0434f136c6f3e0c63": "7YTbEB5iM2P16rYo9YiSDEsa5FnF4ncMvz7nQ2PY5CtG",//50FDAC52A32B35879A2C6519F745AB97
	//"ccc6f40da3875728db92186249e9a81c": "4Nwno7L31Dp3GStp7ncH1sK9AmBwd7Jq45VyDKaFXhzi",//C692C2C536EFD3C11730364D93CF8B0E
}

// server store the map[app_id]=accesskey

//var blurChar string = "!@#qwe!@#"
var blurChar string = "_"

// length is not sure
func generateAccessKey(appId string) string {
	current := time.Now().String()
	hashStr := hashData(appId + blurChar + current)
	return base58.Encode([]byte(md5Encode(hashStr)))
}

// 32 bit
func generateToken(appId string, accessKey string) string {
	md5Str := md5Encode(appId + blurChar + accessKey)
	token := strings.ToUpper(md5Str)
	out := fmt.Sprintf("appId=%s, accessKey=%s, token=%s", appId, accessKey, token)
	fmt.Println(out)
	return token
}

// todo token must upper letters
func VerifyToken(token string, appId string) bool {
	accessKey := ALL_TOKENS_MAP[appId]
	md5Str := md5Encode(appId + blurChar + accessKey)
	tokenReal := strings.ToUpper(md5Str)
	return token == tokenReal
}
