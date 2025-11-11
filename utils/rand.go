package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"

	"github.com/oklog/ulid/v2"
)

// Common character sets for random string generation.
const (
	AlphaNumeric   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	AlphaLowerCase = "abcdefghijklmnopqrstuvwxyz"
	AlphaUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric        = "0123456789"
	AlphaOnly      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	HexLowerCase   = "0123456789abcdef"
	HexUpperCase   = "0123456789ABCDEF"
	Base64Safe     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	SpecialChars   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// OTPLength is the default length for generated OTPs.
var OTPLength int = 6

// GenerateRandomString generates a random string of the specified length using the given charset.
// It uses crypto/rand for cryptographically secure randomness.
func GenerateRandomString(charset string, length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	if len(charset) == 0 {
		return "", errors.New("charset cannot be empty")
	}

	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range length {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result), nil
}

// GenerateAlphaNumeric generates a random alphanumeric string.
func GenerateAlphaNumeric(length int) string {
	s, err := GenerateRandomString(AlphaNumeric, length)
	if err != nil {
		panic(err)
	}
	return s
}

// GenerateNumeric generates a random numeric string.
func GenerateNumeric(length int) string {
	s, err := GenerateRandomString(Numeric, length)
	if err != nil {
		panic(err)
	}
	return s
}

// GenerateAlphaOnly generates a random alphabetic string.
func GenerateAlphaOnly(length int) string {
	s, err := GenerateRandomString(AlphaOnly, length)
	if err != nil {
		panic(err)
	}
	return s
}

// GenerateHex generates a random hexadecimal string (lowercase).
func GenerateHex(length int) string {
	s, err := GenerateRandomString(HexLowerCase, length)
	if err != nil {
		panic(err)
	}
	return s
}

// GeneratePassword generates a strong password with alphanumeric and special characters.
func GeneratePassword(length int) string {
	charset := AlphaNumeric + SpecialChars
	s, err := GenerateRandomString(charset, length)
	if err != nil {
		panic(err)
	}
	return s
}

// GenerateOTP generates a one-time password (numeric string).
// If length is provided, uses that; otherwise uses the default OTPLength.
func GenerateOTP(length ...int) string {
	otpLen := OTPLength
	if len(length) > 0 {
		otpLen = length[0]
	}
	return GenerateNumeric(otpLen)
}

// GenerateULID generates a ULID (Universally Unique Lexicographically Sortable Identifier).
func GenerateULID() string {
	return ulid.Make().String()
}

// GenerateNumber generates a random number with the specified length (as a numeric string converted to int64).
func GenerateNumber(length int) int64 {
	numStr, err := GenerateRandomString(Numeric, length)
	if err != nil {
		panic(err)
	}
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}

// RandomBytes generates random bytes of the specified length.
// Returns nil if generation fails.
func RandomBytes(length int) []byte {
	result := make([]byte, length)
	if _, err := rand.Read(result); err != nil {
		return nil
	}
	return result
}
