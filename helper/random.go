package helper

import (
	"math/rand"
	"strings"
)

func RandString(length int) string {
	baseString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	baseStringLen := len(baseString)
	var randString strings.Builder
	for i := 0; i < length; i++ {

		randIndex := rand.Intn(baseStringLen)
		randString.WriteString(string(baseString[randIndex]))
	}

	return randString.String()
}
