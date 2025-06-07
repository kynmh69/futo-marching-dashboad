package models

import (
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func TestUserHashPassword(t *testing.T) {
	user := &User{
		Username: "testuser",
		Password: "password123",
	}

	err := user.HashPassword()
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	// Check that password was actually hashed
	if user.Password == "password123" {
		t.Error("Password was not hashed")
	}

	// Verify that the hash is valid
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password123"))
	if err != nil {
		t.Errorf("Hashed password does not match original: %v", err)
	}
}

func TestUserCheckPassword(t *testing.T) {
	user := &User{
		Username: "testuser",
		Password: "password123",
	}

	user.HashPassword()

	// Test with correct password
	if !user.CheckPassword("password123") {
		t.Error("CheckPassword failed for correct password")
	}

	// Test with incorrect password
	if user.CheckPassword("wrongpassword") {
		t.Error("CheckPassword passed for incorrect password")
	}
}

func TestUserPrepareCreate(t *testing.T) {
	user := &User{
		Username: "testuser",
		FullName: "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Role:     GeneralRole,
	}

	// Fields should be empty initially
	if !user.CreatedAt.IsZero() {
		t.Error("CreatedAt should be zero before PrepareCreate")
	}
	if !user.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should be zero before PrepareCreate")
	}

	user.PrepareCreate()

	// Fields should be set
	if user.CreatedAt.IsZero() {
		t.Error("CreatedAt not set after PrepareCreate")
	}
	if user.UpdatedAt.IsZero() {
		t.Error("UpdatedAt not set after PrepareCreate")
	}

	// Fields should be set to approximately now
	now := time.Now()
	if now.Sub(user.CreatedAt) > 2*time.Second {
		t.Error("CreatedAt not close to current time")
	}
	if now.Sub(user.UpdatedAt) > 2*time.Second {
		t.Error("UpdatedAt not close to current time")
	}
}

func TestUserPrepareUpdate(t *testing.T) {
	user := &User{
		Username: "testuser",
		FullName: "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Role:     GeneralRole,
	}

	// Set an old updated time
	oldTime := time.Now().Add(-24 * time.Hour)
	user.UpdatedAt = oldTime

	user.PrepareUpdate()

	// UpdatedAt should be newer
	if user.UpdatedAt.Equal(oldTime) || user.UpdatedAt.Before(oldTime) {
		t.Error("UpdatedAt not updated after PrepareUpdate")
	}

	// UpdatedAt should be approximately now
	now := time.Now()
	if now.Sub(user.UpdatedAt) > 2*time.Second {
		t.Error("UpdatedAt not close to current time")
	}
}