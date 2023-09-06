package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(65 + rand.Intn(25))  // A=65 and Z=90
	}
	return string(bytes)
}

func CapitalizeFirstLetter(input string) string {
	return strings.Title(input)
}

func IsStringInSlice(item string, list []string) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
  
}

func TruncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[0:length] + "..."
}
