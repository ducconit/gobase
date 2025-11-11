// Package hash provides bcrypt hashing utilities.
package hash

import "golang.org/x/crypto/bcrypt"

const (
	// DefaultRounds is the default bcrypt cost parameter.
	DefaultRounds = 12
)

// New generates a bcrypt hash from the input string with optional custom rounds.
// If no rounds are specified, DefaultRounds is used.
// Returns the hashed string or an error if hashing fails.
func New(password string, rounds ...int) (string, error) {
	cost := DefaultRounds
	if len(rounds) > 0 {
		cost = rounds[0]
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// Make generates a bcrypt hash from the input string using default rounds.
// It panics if hashing fails. Use New() for error handling.
func Make(password string) string {
	hashedPassword, err := New(password)
	if err != nil {
		panic(err)
	}

	return hashedPassword
}

// Check verifies if the password matches the bcrypt hash.
// Returns true if password is correct, false otherwise.
func Check(plainText, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText)) == nil
}

// IsHash checks if the input string is a valid bcrypt hash.
func IsHash(input string) bool {
	_, err := bcrypt.Cost([]byte(input))
	return err == nil
}
