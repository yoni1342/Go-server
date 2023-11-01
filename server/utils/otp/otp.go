package otp

import (
	"log"
	"math/rand"
)

func GenerateOTP() string {
	var digits = "0123456789"
	var otp string
	for i := 0; i < 4; i++ {
		otp += string(digits[rand.Intn(10)])
	}
	log.Println(otp)
	return otp
}