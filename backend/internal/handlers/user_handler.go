package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kynmh69/futo-marching-dashboad/backend/internal/models"
	"github.com/kynmh69/futo-marching-dashboad/backend/internal/repositories"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	userRepo   repositories.UserRepository
	jwtSecret  string
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userRepo repositories.UserRepository, jwtSecret string) *UserHandler {
	return &UserHandler{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Register registers a new user
func (h *UserHandler) Register(c echo.Context) error {
	var input models.CreateUserInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Check if username already exists
	existingUser, err := h.userRepo.FindByUsername(c.Request().Context(), input.Username)
	if err == nil && existingUser != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username already exists"})
	}

	// Check if email already exists
	existingUser, err = h.userRepo.FindByEmail(c.Request().Context(), input.Email)
	if err == nil && existingUser != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email already exists"})
	}

	// Create new user
	user := &models.User{
		Username: input.Username,
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}

	user.PrepareCreate()
	if err := user.HashPassword(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	id, err := h.userRepo.Create(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	user.ID, _ = primitive.ObjectIDFromHex(id)
	user.Password = "" // Remove password from response

	return c.JSON(http.StatusCreated, user)
}

// Login logs in a user
func (h *UserHandler) Login(c echo.Context) error {
	var input models.LoginInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Find user by username
	user, err := h.userRepo.FindByUsername(c.Request().Context(), input.Username)
	if err != nil || user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	// Check password
	if !user.CheckPassword(input.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	// Create JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID.Hex()
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

// GetMe gets the current user
func (h *UserHandler) GetMe(c echo.Context) error {
	userID := c.Get("user").(jwt.MapClaims)["id"].(string)
	
	user, err := h.userRepo.FindByID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}
	
	user.Password = "" // Remove password from response
	
	return c.JSON(http.StatusOK, user)
}

// GetAllUsers gets all users
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.userRepo.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get users"})
	}
	
	// Remove passwords from response
	for _, user := range users {
		user.Password = ""
	}
	
	return c.JSON(http.StatusOK, users)
}

// GetUser gets a user by ID
func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	
	user, err := h.userRepo.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}
	
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	
	user.Password = "" // Remove password from response
	
	return c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	
	var input models.UpdateUserInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	
	user, err := h.userRepo.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}
	
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	
	// Update fields
	if input.Username != "" {
		user.Username = input.Username
	}
	
	if input.FullName != "" {
		user.FullName = input.FullName
	}
	
	if input.Email != "" {
		user.Email = input.Email
	}
	
	if input.Password != "" {
		user.Password = input.Password
		if err := user.HashPassword(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
		}
	}
	
	if input.Role != "" {
		user.Role = input.Role
	}
	
	user.PrepareUpdate()
	
	if err := h.userRepo.Update(c.Request().Context(), id, user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}
	
	user.Password = "" // Remove password from response
	
	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	
	if err := h.userRepo.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	
	return c.NoContent(http.StatusNoContent)
}