package key

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
)

func AESEncrypt(key []byte, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText = PKCS7Padding(plainText, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	crypted := make([]byte, len(plainText))
	blockMode.CryptBlocks(crypted, plainText)
	return crypted, nil
}

func AESDecrypt(key []byte, crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	plainText := make([]byte, len(crypted))
	blockMode.CryptBlocks(plainText, crypted)
	plainText = PKCS7UnPadding(plainText)
	return plainText, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - (len(ciphertext) % blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func EcdhKeyAgreement(receiverPubKey *ecdsa.PublicKey, senderPrivKey *ecdsa.PrivateKey) ([]byte, error) {
	x, _ := receiverPubKey.Curve.ScalarMult(receiverPubKey.X, receiverPubKey.Y, senderPrivKey.D.Bytes())
	return x.Bytes(), nil
}
