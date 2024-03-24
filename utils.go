package apitoken

import (
	"math/rand"
	"time"
)

var chars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var charsLen = len(chars)

func randomString(size int) string {
	random := make([]byte, size)

	seed := time.Now().UnixMilli()
	rnd := rand.New(rand.NewSource(seed))

	for i := 0; i < size; i++ {
		random[i] = chars[rnd.Intn(charsLen)]
	}

	return string(random)
}
