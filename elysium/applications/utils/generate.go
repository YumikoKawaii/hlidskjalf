package utils

import "math/rand"

// FakeString generates a random string of given length
func FakeString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// FakeInt generates a random integer between min and max (inclusive)
func FakeInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// FakeDouble generates a random double between min and max
func FakeDouble(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// FakeBoolean generates a random boolean value
func FakeBoolean() bool {
	return rand.Intn(2) == 1
}
