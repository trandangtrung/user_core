package utils

import "math/rand"

func GenerateOTP() string {

	// Generate a random 6-digit OTP
	otp := make([]byte, 6)
	for i := range otp {
		otp[i] = '0' + byte(rand.Intn(10))
	}
	return string(otp)
}
