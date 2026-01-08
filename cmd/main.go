package main

import (
	"log/slog"
	"os"

	"pizza-tracker-go/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := loadConfig()

	// Logger
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, nil),
	)
	slog.SetDefault(logger)

	// Database
	dbModel, err := models.InitDB(cfg.DBPath)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	slog.Info("Database initialized successfully")

	// Validators
	RegisterCustomValidators()

	// Handlers
	h := NewHandler(dbModel)

	// Router
	router := gin.Default()

	// Templates (CRÍTICO)
	if err := loadTemplates(router); err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	// Routes
	setupRoutes(router, h)

	slog.Info("Server starting", "url", "http://localhost:"+cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
