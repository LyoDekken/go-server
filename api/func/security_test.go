package security

import (
	"testing"
)

func TestHashAndVerifyPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := Hash(password)
	if err != nil {
		t.Fatalf("Hash function returned an error: %v", err)
	}

	// Test if the hashed password is not equal to the original password
	if string(hashedPassword) == password {
		t.Errorf("Hashed password is equal to the original password")
	}

	// Test if VerifyPassword function returns nil error for correct password
	err = VerifyPassword(string(hashedPassword), password)
	if err != nil {
		t.Errorf("VerifyPassword function returned an error for correct password: %v", err)
	}

	// Test if VerifyPassword function returns an error for incorrect password
	err = VerifyPassword(string(hashedPassword), "wrongPassword")
	if err == nil {
		t.Errorf("VerifyPassword function should have returned an error for incorrect password")
	}
}
