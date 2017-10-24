package config

import (
	"fmt"
	"testing"
)

/**
 * function :
 * param   :
 * return :
 */

func TestWriteConToFile(t *testing.T) {
	WriteConToFile()
}

func TestTest(t *testing.T) {
	fmt.Println(Config.Keyring)
	fmt.Println(Config.Keypair.PrivateKey)
	fmt.Println(Config.Keypair.PublicKey)
}

func TestReadUnicontractConfig(t *testing.T) {
	fmt.Println(ReadUnicontractConfig())
}

func TestGetAllPublicKey(t *testing.T) {
	a := GetAllPublicKey()
	fmt.Println(a)
}
