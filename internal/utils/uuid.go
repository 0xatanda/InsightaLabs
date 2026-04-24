package utils

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

func UUIDv7() string {
	b := make([]byte, 16)
	rand.Read(b)

	b[6] = (b[6] & 0x0f) | 0x70
	b[8] = (b[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func ParseInt(val string, fallback int) int {
	n, err := strconv.Atoi(val)
	if err != nil || n <= 0 {
		return fallback
	}
	return n
}
