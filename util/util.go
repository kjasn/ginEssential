package util

import (
	"math/rand"
	"time"
)

func GetRandomString(sz int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*")
	var ret = make([]byte, sz)
	rand.Seed(time.Now().Unix())
	for i := range ret {
		ret[i] = letters[rand.Intn(len(letters))]
	}

	return string(ret)
}
