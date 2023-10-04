package utils

import (
	"math/rand"
	"strconv"
)

func GenerateOtp() string {
	var otpLength = 5
	var otp string
	for i := 0; i < otpLength; i++ {
		num := rand.Intn(10)
		otp += strconv.Itoa(num)
	}
	return otp
}
