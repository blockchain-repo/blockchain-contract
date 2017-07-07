package common

import (
	"bytes"
	"encoding/hex"
	"hash"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/sha3"
	"unicontract/src/common/uniledgerlog"
)

func HashData(val string) string {
	var hash hash.Hash
	var x string = ""
	hash = sha3.New256()
	if hash != nil {
		hash.Write([]byte(val))
		x = hex.EncodeToString(hash.Sum(nil))
	}
	return x
}

func GenerateKeyPair() (string, string) {
	publicKeyBytes, privateKeyBytes, err := ed25519.GenerateKey(nil)
	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	publicKeyBase58 := base58.Encode(publicKeyBytes)
	privateKeyBase58 := base58.Encode(privateKeyBytes[0:32])
	return publicKeyBase58, privateKeyBase58
}

func GetPubByPriv(priv string) string {
	privByte := base58.Decode(priv)
	publicKeyBytes, _, err := ed25519.GenerateKey(bytes.NewReader(privByte))
	if err != nil {
		uniledgerlog.Error(err.Error())
	}
	publicKeyBase58 := base58.Encode(publicKeyBytes)
	return publicKeyBase58
}

func Sign(priv string, msg string) string {
	pub := GetPubByPriv(priv)
	privByte := base58.Decode(priv)
	pubByte := base58.Decode(pub)
	privateKey := make([]byte, 64)
	copy(privateKey[:32], privByte)
	copy(privateKey[32:], pubByte)
	sigByte := ed25519.Sign(privateKey, []byte(msg))
	return base58.Encode(sigByte)
}

func Verify(pub string, msg string, sig string) bool {
	pubByte := base58.Decode(pub)
	publicKey := make([]byte, 32)
	copy(publicKey, pubByte)
	return ed25519.Verify(publicKey, []byte(msg), base58.Decode(sig))
}
