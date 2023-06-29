package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	rand2 "math/rand"
	"os"
	"time"
)

const (
	HashAlgorithm = "pbkdf2_sha256"
	Iterations    = 216000
	KeyLength     = 32
	SaltLength    = 12

	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitBytes  = "0123456789"
)

func main() {

	// 从命令行参数中获取参数值
	args := os.Args[1:]

	// 检查是否传入了参数
	if len(args) == 0 {
		fmt.Println("请传入一个参数")
		return
	}

	// 获取第一个参数值
	password := args[0]
	println("原始密码:", password)
	hashedPassword := GeneratePBKDF2SHA256Hash(password)
	println("最终生成加密串:", hashedPassword)
	//fmt.Println(hashedPassword)
}

func GeneratePBKDF2SHA256Hash(password string) string {
	rand2.Seed(time.Now().UnixNano())

	saltByte := generateRandomString(SaltLength)
	//salt_byte := generateRandomSalt(SaltLength)

	salt := []byte(saltByte)

	println("加密盐:", string(salt))
	hashedPassword := pbkdf2.Key([]byte(password), salt, Iterations, KeyLength, sha256.New)
	hashedPasswordBase64 := base64.StdEncoding.EncodeToString(hashedPassword)

	return fmt.Sprintf("%s$%d$%s$%s", HashAlgorithm, Iterations, string(salt), hashedPasswordBase64)
}

func generateRandomSalt(length int) []byte {
	salt := make([]byte, length)
	_, _ = rand.Read(salt)

	return salt
}
func generateRandomString(length int) string {
	charset := letterBytes + digitBytes
	chars := []byte(charset)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = chars[rand2.Intn(len(chars))]
	}

	return string(result)
}
