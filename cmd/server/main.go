package main

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	assetApp "ddd/internal/features/asset/app"
	infraAsset "ddd/internal/features/asset/infra"
	"ddd/internal/interfaces/http"
	assethandler "ddd/internal/interfaces/http/handlers/asset"
	"ddd/pkg/config"
	"ddd/shared/logger"
	"ddd/shared/logger/openob"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Asset Service API
// @version 1.0
// @description IoT Asset Management Service API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.your-company.com/support
// @contact.email support@your-company.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @tag.name assets
// @tag.description Asset management operations

// @tag.name measures
// @tag.description Measurement operations

// @tag.name dashboards
// @tag.description Dashboard operations

// @tag.name widgets
// @tag.description Widget operations
func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	var logger logger.Logger
	// Initialize logger
	logger, err = openob.NewLogger(openob.Config{
		ServiceName:   cfg.App.Name,
		Environment:   cfg.App.Environment,
		Endpoint:      cfg.Telemetry.Endpoint,
		Headers:       cfg.Telemetry.Headers,
		OrgID:         cfg.Telemetry.OrgID,
		StreamName:    cfg.Telemetry.StreamName,
		ConsoleOutput: true,
		EnableTracing: cfg.Telemetry.TracingEnabled,
	})
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	ctx := context.Background()
	// Initialize tracer if enabled
	cfg.Telemetry.TracingEnabled = false
	if cfg.Telemetry.TracingEnabled {

		// Info logging
		logger.Info(ctx, "Server starting", map[string]interface{}{
			"port": cfg.HTTP.Port,
			"env":  cfg.App.Environment,
		})

		logger.Info(ctx, "Attempting to connect to database", map[string]interface{}{}) // Log before database connection
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logger.Error(ctx, err, "Failed to connect to database", nil)
		return
	}

	assetRepo := infraAsset.NewRepository(db)
	assetService := assetApp.NewService(assetRepo, logger)

	deps := http.ServerDependencies{
		AssetHandler: assethandler.New(assethandler.Dependencies{
			Logger:       logger,
			AssetService: assetService,
		}),
	}

	server := http.NewServer(cfg, logger, deps)
	err = server.Start(context.Background())
	if err != nil {
		logger.Error(ctx, err, "error init http server", nil)
	}

	u := struct {
		ID       int
		Name     string
		Email    string
		Age      int
		IsActive bool
	}{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Age:      30,
		IsActive: true,
	}
	d := struct {
		ID          int
		Car         string
		Manufacture bool
	}{
		ID:          1,
		Car:         "Ford",
		Manufacture: true,
	}
	for {
		u.Age++
		d.Car = "Ford" + strconv.Itoa(u.Age)
		logger.Info(ctx, "hello", u)
		logger.Info(ctx, "ani", d)
		time.Sleep(5 * time.Second)
		if u.Age > 35 {
			logger.Error(ctx, errors.New("too old"), "this guy is too old", nil)
		}
	}
}
