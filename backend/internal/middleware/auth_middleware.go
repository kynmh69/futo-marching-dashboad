package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kynmh69/futo-marching-dashboad/backend/internal/models"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware creates a middleware for JWT authentication
func JWTMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header required"})
			}

			// Check if the Authorization header has the correct format
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Authorization header format"})
			}

			tokenString := parts[1]

			// Parse JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token signing method")
				}
				return []byte(jwtSecret), nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
			}

			// Check if token is valid
			if !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			// Set user claims in context
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token claims"})
			}

			c.Set("user", claims)

			return next(c)
		}
	}
}

// RoleMiddleware creates a middleware for role-based authorization
func RoleMiddleware(roles ...models.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get user claims from context
			claims, ok := c.Get("user").(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}

			// Check if user has required role
			userRole, ok := claims["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user role"})
			}

			// Check if the user's role is in the allowed roles
			hasRole := false
			for _, role := range roles {
				if string(role) == userRole {
					hasRole = true
					break
				}
			}

			if !hasRole {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "Insufficient permissions"})
			}

			return next(c)
		}
	}
}