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
	fmt.Println(Config.Keypair.Private)
	fmt.Println(Config.Keypair.Public)

}

func TestReadUnicontractConfig(t *testing.T) {
	fmt.Println(ReadUnicontractConfig())
}