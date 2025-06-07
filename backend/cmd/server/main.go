package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kynmh69/futo-marching-dashboad/backend/internal/config"
	"github.com/kynmh69/futo-marching-dashboad/backend/internal/handlers"
	"github.com/kynmh69/futo-marching-dashboad/backend/internal/middleware"
	"github.com/kynmh69/futo-marching-dashboad/backend/internal/models"
	"github.com/kynmh69/futo-marching-dashboad/backend/internal/repositories"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	defer cfg.Close()

	// Create repositories
	userRepo := repositories.NewUserMongoRepository(cfg.DBClient, cfg.DBName)

	// Create handlers
	userHandler := handlers.NewUserHandler(userRepo, cfg.JWTSecret)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"https://trusted-domain.com", "https://another-trusted-domain.com"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType},
	}))

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Welcome to FUTO Marching Dashboard API"})
	})

	// Auth routes
	e.POST("/api/auth/register", userHandler.Register)
	e.POST("/api/auth/login", userHandler.Login)

	// API routes
	api := e.Group("/api")
	api.Use(middleware.JWTMiddleware(cfg.JWTSecret))

	// User routes
	api.GET("/users/me", userHandler.GetMe)
	
	// Admin routes
	admin := api.Group("/admin")
	admin.Use(middleware.RoleMiddleware(models.AdminRole))
	
	admin.GET("/users", userHandler.GetAllUsers)
	admin.GET("/users/:id", userHandler.GetUser)
	admin.PUT("/users/:id", userHandler.UpdateUser)
	admin.DELETE("/users/:id", userHandler.DeleteUser)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Printf("Server running on port %s\n", port)
	e.Logger.Fatal(e.Start(":" + port))
}