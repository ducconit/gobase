package hash

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestNew(t *testing.T) {
	password := "testpassword123"

	// Test with default rounds
	hash, err := New(password)
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	if hash == "" {
		t.Error("New() returned empty hash")
	}

	// Verify hash is valid bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		t.Errorf("Generated hash is invalid: %v", err)
	}
}

func TestNewWithCustomRounds(t *testing.T) {
	password := "testpassword123"
	customRounds := 10

	hash, err := New(password, customRounds)
	if err != nil {
		t.Fatalf("New() with custom rounds failed: %v", err)
	}

	if hash == "" {
		t.Error("New() returned empty hash")
	}

	// Verify hash is valid
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		t.Errorf("Generated hash is invalid: %v", err)
	}
}

func TestMake(t *testing.T) {
	password := "testpassword123"

	hash := Make(password)
	if hash == "" {
		t.Error("Make() returned empty hash")
	}

	// Verify hash is valid
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		t.Errorf("Generated hash is invalid: %v", err)
	}
}

func TestDifferentHashesForSamePassword(t *testing.T) {
	password := "testpassword123"

	hash1 := Make(password)
	hash2 := Make(password)

	if hash1 == hash2 {
		t.Error("Two hashes of the same password should be different")
	}

	// But both should verify correctly
	if err := bcrypt.CompareHashAndPassword([]byte(hash1), []byte(password)); err != nil {
		t.Errorf("hash1 verification failed: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash2), []byte(password)); err != nil {
		t.Errorf("hash2 verification failed: %v", err)
	}
}

func TestMakePanicLongPassword(t *testing.T) {
	// Create password longer than 72 bytes
	longPassword := strings.Repeat("a", 73)

	defer func() {
		if r := recover(); r == nil {
			t.Error("Make() should panic with password > 72 bytes")
		}
	}()

	Make(longPassword)
}

func TestNewErrorLongPassword(t *testing.T) {
	// Create password longer than 72 bytes
	longPassword := strings.Repeat("a", 73)

	_, err := New(longPassword)
	if err == nil {
		t.Error("New() should return error with password > 72 bytes")
	}
}

func TestCheck(t *testing.T) {
	password := "testpassword123"
	hash := Make(password)

	// Test correct password
	if !Check(password, hash) {
		t.Error("Check() should return true for correct password")
	}

	// Test incorrect password
	if Check("wrongpassword", hash) {
		t.Error("Check() should return false for incorrect password")
	}
}

func TestCheckWithEmptyPassword(t *testing.T) {
	password := ""
	hash := Make(password)

	if !Check(password, hash) {
		t.Error("Check() should return true for correct empty password")
	}

	if Check("nonempty", hash) {
		t.Error("Check() should return false for incorrect password")
	}
}

func TestIsHash(t *testing.T) {
	password := "testpassword123"
	validHash := Make(password)

	// Test valid hash
	if !IsHash(validHash) {
		t.Error("IsHash() should return true for valid bcrypt hash")
	}

	// Test invalid hash - too short
	if IsHash("tooshort") {
		t.Error("IsHash() should return false for short string")
	}

	// Test invalid hash - wrong prefix
	if IsHash(strings.Repeat("a", 60)) {
		t.Error("IsHash() should return false for string without bcrypt prefix")
	}

	// Test invalid hash - empty string
	if IsHash("") {
		t.Error("IsHash() should return false for empty string")
	}

	// Test invalid hash - wrong length
	if IsHash(strings.Repeat("a", 59)) {
		t.Error("IsHash() should return false for wrong length")
	}

	// Test invalid hash - correct prefix but invalid cost
	invalidHash := "$2a$99" + strings.Repeat("a", 53)
	if IsHash(invalidHash) {
		t.Error("IsHash() should return false for invalid bcrypt cost")
	}
}
