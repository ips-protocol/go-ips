package keystore

import (
	"encoding/hex"
	"io/ioutil"

	ethKs "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ipfs/go-ipfs/chain/base64"
)

func VerifyKeystore(ksFile string, password string) (valid bool, err error) {
	ksData, err := ioutil.ReadFile(ksFile)
	if err != nil {
		return false, err
	}
	_, err = ethKs.DecryptKey(ksData, password)
	if err != nil {
		return false, err
	}
	return true, nil
}

func PrivateKeyFromKeystore(ksFile string, password string) (privateKey string, err error) {
	ksData, err := ioutil.ReadFile(ksFile)
	if err != nil {
		return "", err
	}
	key, err := ethKs.DecryptKey(ksData, password)
	if err != nil {
		return "", err
	}
	privateKey = hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	return privateKey, nil
}

func GeneratePassword(id string) string {
	base64 := base64.RawStdEncoding.EncodeToString([]byte(id[2:]))
	return base64[10:26]
}

func GenerateKeystore(ksPath string, hexkey string, passphrase string) (ksFile string, err error) {
	ks := ethKs.NewKeyStore(ksPath, ethKs.StandardScryptN, ethKs.StandardScryptP)

	privKey, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return "", err
	}

	_, err = ks.ImportECDSA(privKey, passphrase)
	if err != nil {
		return "", err
	}

	files, err := ioutil.ReadDir(ksPath)
	if err != nil {
		return "", err
	}
	
	ksFile = ksPath + "/" + files[0].Name()

	return ksFile, nil
}
