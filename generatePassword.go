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

	randomString = randomString[:16]
	return randomString, err
}
