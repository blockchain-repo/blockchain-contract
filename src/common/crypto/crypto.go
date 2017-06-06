/*************************************************************************
  > File Name: scantaskschedule.go
  > Module:
  > Function: AES 加解密操作
  > Author: wangyp
  > Company:
  > Department:
  > Mail: wangyepeng87@163.com
  > Created Time: 2017.06.06
 ************************************************************************/
package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

//---------------------------------------------------------------------------
// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
var gstrKey = []byte("s*e0)2|f_9fd&{}9*1qa^3dj^&u#~!fl")
var gslPaddingData = []byte("e�5�D�")
var gFlag = "?"

//---------------------------------------------------------------------------
func AesEncryptToFile(strFilePath, strSaveFilePath string) error {
	if len(strFilePath) == 0 ||
		len(strSaveFilePath) == 0 {
		return fmt.Errorf("param is null")
	}

	slOrigData, err := ReadFile(strFilePath)
	if err != nil {
		return err
	}

	// 第一次加密
	result, err := AesEncrypt(slOrigData, gstrKey)
	if err != nil {
		return err
	}

	// 处理一次加密后的数据
	var slHandleDataOver []byte
	slHandleDataOver = append(slHandleDataOver, result[:10]...)
	slHandleDataOver = append(slHandleDataOver, gslPaddingData...)
	slHandleDataOver = append(slHandleDataOver, result[10:]...)

	// 第二次加密
	result, err = AesEncrypt(slHandleDataOver, gstrKey)
	if err != nil {
		return err
	}
	secondEncryptData := base64.StdEncoding.EncodeToString(result)

	// 加标志位
	slSaveData := []byte(gFlag)
	slSaveData = append(slSaveData, []byte(secondEncryptData)...)

	// 存储文件
	count, err := WriteFile(strSaveFilePath, string(slSaveData))
	if err != nil {
		return err
	}
	if count != len(slSaveData) {
		return fmt.Errorf("write error")
	}
	return nil
}

//---------------------------------------------------------------------------
func AesDecryptFromFile(strFilePath string) ([]byte, error) {
	if len(strFilePath) == 0 {
		return nil, fmt.Errorf("param is null")
	}

	slData, err := ReadFile(strFilePath)
	if err != nil {
		return nil, err
	}
	if slData[0] == []byte(gFlag)[0] {
		sldata, _ := base64.StdEncoding.DecodeString(string(slData[1:]))
		origData, err := AesDecrypt(sldata, gstrKey)
		if err != nil {
			return nil, err
		}

		// 处理第一次解密数据
		var handleDataOver_ []byte
		handleDataOver_ = append(handleDataOver_, origData[:10]...)
		handleDataOver_ = append(handleDataOver_, origData[len(gslPaddingData)+10:]...)

		// 第二次解密
		origData, err = AesDecrypt(handleDataOver_, gstrKey)
		if err != nil {
			return nil, err
		}
		return origData, nil
	} else {
		return slData, nil
	}
	return nil, nil
}

//---------------------------------------------------------------------------
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//---------------------------------------------------------------------------
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

//---------------------------------------------------------------------------
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

//---------------------------------------------------------------------------
func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//---------------------------------------------------------------------------
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//---------------------------------------------------------------------------
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//---------------------------------------------------------------------------
func ReadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []byte(""), err
	}
	defer file.Close()
	detailByte, err := ioutil.ReadAll(file)
	return detailByte, err
}

//---------------------------------------------------------------------------
func WriteFile(fileName string, content string) (int, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.WriteString(content)
}

//---------------------------------------------------------------------------
