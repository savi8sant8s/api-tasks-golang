package service

import (
  "crypto/rand"
  "fmt"
)

func GenerateToken(size uint) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
