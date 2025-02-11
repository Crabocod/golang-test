package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashGenerate(text string, algorithm string) (string, error) {
	switch algorithm {
	case "md5":
		hash := md5.Sum([]byte(text))
		return hex.EncodeToString(hash[:]), nil
	case "sha256":
		hash := sha256.Sum256([]byte(text))
		return hex.EncodeToString(hash[:]), nil
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
	}
}
