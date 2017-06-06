// cryptotool
package main

import (
	"fmt"
	"os"
)

import (
	"unicontract/src/common/crypto"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(fmt.Errorf("error : input filename is null.\neg. ./cryptotool ./unicontract"))
		os.Exit(-1)
	}

	strOriginalFile := os.Args[1]
	strCryptoFile := strOriginalFile + ".crypto"

	err := crypto.AesEncryptToFile(strOriginalFile, strCryptoFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("ok")
}
