package utils

import (
	nanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

const (
	alphabetNum = "0123456789abcdefghijklmnopqrstuvwxyz"
	length      = 12
)

// New generates a unique public ID.
func NanoIDS() (string, error) { return nanoid.Generate(alphabetNum, length) }

//PASSWORDS

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 	hashedPasswordBase64 := base64.StdEncoding.EncodeToString(bytes)
// 	return hashedPasswordBase64, err
// }
