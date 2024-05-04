package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"math/big"
	mathrand "math/rand"
	"time"
)

func Md5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func guidPart() string {
	b := make([]byte, 2)

	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func GetGuid() string {
	return guidPart() + guidPart() + guidPart() + guidPart()
}

func aesEncrypt(text, secKey string, iv []byte) (string, error) {
	block, err := aes.NewCipher([]byte(secKey))
	if err != nil {
		return "", err
	}
	paddedText := pkcs7Padding([]byte(text), block.BlockSize())
	ciphertext := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedText)

	return hex.EncodeToString(ciphertext), nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func rsaEncrypt(data string) (string, error) {
	eStr := "10001"
	e := new(big.Int)
	e.SetString(eStr, 16)

	nStr := "00C1E3934D1614465B33053E7F48EE4EC87B14B95EF88947713D25EECBFF7E74C7977D02DC1D9451F79DD5D1C10C29ACB6A9B4D6FB7D0A0279B6719E1772565F09AF627715919221AEF91899CAE08C0D686D748B20A3603BE2318CA6BC2B59706592A9219D0BF05C9F65023A21D2330807252AE0066D59CEEFA5F2748EA80BAB81"
	n := new(big.Int)
	n.SetString(nStr, 16)

	pubKey := rsa.PublicKey{
		N: n,
		E: int(e.Int64()),
	}

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, &pubKey, []byte(data))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(encrypted), nil
}

func EncryptPayload(data string) (string, error) {
	newg := GetGuid()

	cipherText, err := rsaEncrypt(newg)

	if err != nil {
		return "", err
	}

	encr, err := aesEncrypt(data, newg, []byte("0000000000000000"))

	if err != nil {
		return "", err
	}

	return (encr + cipherText), nil
}

func GetRandom() int64 {
	return int64(mathrand.Float64()*10000) + time.Now().UnixMilli()
}
