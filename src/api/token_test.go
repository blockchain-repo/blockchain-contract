package api

import (
	"github.com/google/uuid"
	"testing"
)

func generateAppId() string {
	uuidVal := uuid.New().String()
	appId := "uniapp_" + uuidVal
	appId = md5Encode(appId)
	return appId
}
func Test_GenerateAppId(t *testing.T) {
	appId := generateAppId
	t.Log("appId", appId)
}

// length is not sure
func Test_GenerateAccessKey(t *testing.T) {
	appId := "4B3103422D6180E2E4E03D22DC74EB2B"
	accessKey := generateAccessKey(appId)
	t.Log("accessKey", accessKey)
}

func Test_GenerateTokenData(t *testing.T) {
	appId := "4B3103422D6180E2E4E03D22DC74EB2B"
	accessKey := generateAccessKey(appId)
	token := generateToken(appId, accessKey)
	t.Log("token", token)
}

func Test_VerifyToken(t *testing.T) {
	appId := generateAppId()
	//accessKey := "4ecxjG9fJtXir1FvG59dnXoNeLfNjTwGND35dTR5bFmw"
	// 4B3103422D6180E2E4E03D22DC74EB2B
	accessKey := generateAccessKey(appId)
	token := generateToken(appId, accessKey)
	t.Log("token", token)
	t.Log("verify ", VerifyToken(token, appId))
}
