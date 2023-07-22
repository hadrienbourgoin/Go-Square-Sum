package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"time"
)

// TOTP settings
const (
	TimeStep     = 30 // TOTP's Time Step X in seconds
	T0           = 0  // T0 is 0
	PasswordSize = 10 // Length of the TOTP password (10 digits)
)

func generateTOTP(secret string) (string, error) {
	// Calculate the number of time steps since T0 (0).
	timeSteps := time.Now().Unix() / TimeStep

	// Convert the time steps to an 8-byte big-endian byte slice.
	timeStepsBytes := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		timeStepsBytes[i] = byte(timeSteps)
		timeSteps >>= 8
	}

	// Concatenate the secret and "HENNGECHALLENGE003".
	sharedSecret := []byte(secret + "HENNGECHALLENGE003")

	// Calculate the HMAC-SHA-512 of the concatenated value.
	h := hmac.New(sha512.New, sharedSecret)
	h.Write(timeStepsBytes)
	hashed := h.Sum(nil)

	// Get the offset value from the last 4 bits of the HMAC-SHA-512 digest.
	offset := hashed[len(hashed)-1] & 0x0F

	// Get the 4 bytes at the offset in the HMAC-SHA-512 digest.
	truncatedHash := hashed[offset : offset+4]

	// Convert the truncated hash to an integer value.
	otp := int(truncatedHash[0])<<24 | int(truncatedHash[1])<<16 | int(truncatedHash[2])<<8 | int(truncatedHash[3])

	// Apply a mask to limit the OTP to 10 digits.
	otp = otp & 0x7FFFFFFF

	// Convert the OTP to a 10-digit string.
	otpString := fmt.Sprintf("%010d", otp%int(1e10))

	return otpString, nil
}

func main() {
	// Set the shared secret based on the specified setup.
	sharedSecret := "hadrienbourgoin@gmail.com"

	totp, err := generateTOTP(sharedSecret)
	if err != nil {
		fmt.Println("Error generating TOTP:", err)
		return
	}

	fmt.Println("TOTP Password (10 digits):", totp)
}
