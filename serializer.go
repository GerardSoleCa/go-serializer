package serializer

import (
	"errors"
	"encoding/hex"
	"github.com/GerardSoleCa/go-aes"
	"math/rand"
	"time"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1 << letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func SecureStringify(str, encryptKey, validateKey string) (string, error) {
	nonceCheck := randomString(8)
	nonceCrypt := randomString(8)

	encrypted, err := go_aes.Encrypt([]byte(nonceCheck + str), encryptKey + nonceCrypt)
	if err != nil {
		return "", errors.New("ERROR ON ENCRYPTION")
	}
	digest := signStr(str, validateKey + nonceCheck)
	return digest + nonceCrypt + hex.EncodeToString(encrypted), nil
}

func SecureParse(str string, encryptKey string, validateKey string) (string, error) {
	// Split the encrypted string (str)
	expectedDigest := str[0:28]
	nonceCrypt := str[28:36]
	encryptedString := str[36:len(str)]

	// Get the binary representation of the encrypted string
	binaryData, _ := hex.DecodeString(encryptedString)

	decryptionKey := encryptKey + nonceCrypt;

	decrypted, err := go_aes.Decrypt(binaryData, decryptionKey)

	if err != nil {
		return "", errors.New("ERROR ON DECRYPTION")
	}

	data := string(decrypted)

	// Split the decrypted data to get nonce and plain
	nonceCheck := data[0:8]
	plain := data[8:len(data)]

	var digest = signStr(plain, validateKey + nonceCheck)

	if expectedDigest != digest {
		return "", errors.New("BAD SIGNATURE")
	} else {
		return plain, nil
	}
}

func randomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n - 1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}



