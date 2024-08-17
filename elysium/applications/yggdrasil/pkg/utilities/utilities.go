package utilities

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const sexes = "MFU"

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandomNumber(lower, upper int) int {
	return rand.Intn(upper-lower+1) + lower
}

func RandomSex() string {
	return string(sexes[RandomNumber(0, 2)])
}
