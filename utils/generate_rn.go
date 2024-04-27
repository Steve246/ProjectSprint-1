package utils

import (
	"math/rand"
	"strconv"
)

func GenerateSN() string {
	result := strconv.Itoa(rand.Intn(9999)) + " " + strconv.Itoa(rand.Intn(9999)) + " " + strconv.Itoa(rand.Intn(9999)) + " " + strconv.Itoa(rand.Intn(9999))

	return result
}
