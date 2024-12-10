// internal/interfaces/http/server.go
package http

import (
	"context"
	"fmt"
	"net/http"

	"ddd/internal/interfaces/http/handlers/asset"
	"ddd/pkg/config"
	"ddd/shared/httputil"
	"ddd/shared/logger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config       *config.Config
	logger       logger.Logger
	router       *gin.Engine
	httpServer   *http.Server
	assetHandler *assethandler.AssetHandler
	// measureHandler   *handlers.MeasureHandler
	// dashboardHandler *handlers.DashboardHandler
	// widgetHandler    *handlers.WidgetHandler
}

type ServerDependencies struct {
	AssetHandler *assethandler.AssetHandler
	// MeasureHandler   *handlers.MeasureHandler
	// DashboardHandler *handlers.DashboardHandler
	// WidgetHandler    *handlers.WidgetHandler
}

func NewServer(cfg *config.Config, logger logger.Logger, deps ServerDependencies) *Server {
	router := gin.New()

	// Setup middleware
	router.Use(
		gin.Recovery(),
		//httputil.LoggerMiddleware(logger),
		httputil.RecoveryMiddleware(logger),
		httputil.OrgIDMiddleware(),
	)

	return &Server{
		config:       cfg,
		logger:       logger,
		router:       router,
		assetHandler: deps.AssetHandler,
		// measureHandler:   deps.MeasureHandler,
		// dashboardHandler: deps.DashboardHandler,
		// widgetHandler:    deps.WidgetHandler,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.setupRoutes()

	addr := fmt.Sprintf("%s:%d", s.config.HTTP.Host, s.config.HTTP.Port)
	s.httpServer = &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  s.config.HTTP.ReadTimeout,
		WriteTimeout: s.config.HTTP.WriteTimeout,
	}

	s.logger.Info(ctx, "Starting HTTP server on %s"+addr, map[string]interface{}{})
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
