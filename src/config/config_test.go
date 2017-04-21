package config

import (
	"testing"
	"fmt"
)

/**
 * function : 
 * param   :
 * return : 
 */

func TestWriteConToFile(t *testing.T) {
	WriteConToFile()
}

func TestTest(t *testing.T){
	fmt.Println(Config.Keyrings)
	fmt.Println(Config.Keypair.PrivateKey)
	fmt.Println(Config.Keypair.PublicKey)

}

func TestReadUnicontractConfig(t *testing.T) {
	fmt.Println(ReadUnicontractConfig())
}

