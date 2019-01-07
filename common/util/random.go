package util

import "math/rand"

const baseStr = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`

func RandomStr(length int) string {
	randomStr := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		limit := int64(len(baseStr))
		c := byte(baseStr[rand.Int63()%limit])
		randomStr = append(randomStr, c)
	}
	return string(randomStr)
}