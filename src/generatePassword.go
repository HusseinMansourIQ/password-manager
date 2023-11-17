package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func genereatePassword() (string, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	fmt.Println(randomBytes)
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	fmt.Println(randomString)
	// base64.URLEncoding may include '+' and '/' characters, which are not URL-safe.
	// Replace them with '-' and '_' respectively.
	randomString = randomString[:16]
	return randomString, nil
}
