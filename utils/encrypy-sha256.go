package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

const (
	saltSize   = 16
	iterations = 1e4
)

// 密码验证
func ComparePassword(hashStr, pw string) bool {

	hash, _ := hex.DecodeString(hashStr)

	return bytes.Equal(hash, hashPassword([]byte(pw), hash[:saltSize]))
}

// 生成加密 密码
func GenerateHashPassword(password string) (string, error) {
	salt, err := generateRandomBytes(saltSize)
	if err != nil {
		return "", err
	}

	hash := hashPassword([]byte(password), salt)

	return hex.EncodeToString(hash), nil
}

// hash the password with the provided salt using the pbkdf2 algorithm
// return byte slice containing salt (first 64 bytes) and hash (last 32 bytes) => total of 96 bytes
func hashPassword(pw, salt []byte) []byte {
	ret := make([]byte, len(salt))
	copy(ret, salt)
	return append(ret, key(pw, salt, iterations, sha256.Size, sha256.New)...)
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

// copied from "golang.org/x/crypto/pbkdf2" because it's not available in playground
func key(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	prf := hmac.New(h, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen

	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		// N.B.: || means concatenation, ^ means XOR
		// for each block T_i = U_1 ^ U_2 ^ ... ^ U_iter
		// U_1 = PRF(password, salt || uint(i))
		prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)

		// U_n = PRF(password, U_(n-1))
		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:keyLen]
}
