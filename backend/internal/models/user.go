package models

import (
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Role represents user roles in the system
type Role string

const (
	// AdminRole represents an administrator user
	AdminRole Role = "admin"
	// GeneralRole represents a regular user
	GeneralRole Role = "general"
)

// User represents a user in the system
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	FullName  string             `bson:"fullName" json:"fullName"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"` // Password is not included in JSON responses
	Role      Role               `bson:"role" json:"role"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// CreateUserInput represents data needed to create a new user
type CreateUserInput struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     Role   `json:"role" validate:"required,oneof=admin general"`
}

// UpdateUserInput represents data needed to update an existing user
type UpdateUserInput struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Role     Role   `json:"role" validate:"omitempty,oneof=admin general"`
}

// LoginInput represents data needed for user login
type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RegisterInput represents data needed for simplified user registration
type RegisterInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

// HashPassword creates a bcrypt hash of the password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares the provided password with the stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// PrepareCreate sets fields needed for creating a new user
func (u *User) PrepareCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}

// PrepareUpdate sets fields needed for updating a user
func (u *User) PrepareUpdate() {
	u.UpdatedAt = time.Now()
}

// ValidatePassword checks if password meets complexity requirements
// Password must contain at least one uppercase letter, one lowercase letter, and one digit
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}

	return nil
}