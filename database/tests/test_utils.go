package dbq

import (
	"fmt"
	"math/rand"
)

// GenerateRandomEmail generates a random email address
func GenerateRandomEmail() string {
	domains := []string{"gmail.com", "yahoo.com", "outlook.com", "icloud.com"}
	randomDomain := domains[rand.Intn(len(domains))]

	username := fmt.Sprintf("user%d", rand.Intn(1000000)) // Random username

	return fmt.Sprintf("%s@%s", username, randomDomain)
}

// GenerateRandomPassword generates a random password of given length
func GenerateRandomPassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
