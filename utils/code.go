package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateCode(data string) string {

	timeStmp := time.Now().UnixNano()

	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}

	combined := fmt.Sprintf("%x%d%s", randomBytes, timeStmp, data)

	encoded := base64.StdEncoding.EncodeToString([]byte(combined))

	if len(encoded) < 7 {
		return ""
	}

	return encoded[:7]
}
