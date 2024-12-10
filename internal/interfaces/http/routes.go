package http

import "ddd/shared/httputil"

func (s *Server) setupRoutes() {
	v1 := s.router.Group("api/v1/:orgid")
	{
		v1.Use(httputil.OrgIDMiddleware())
		v1.Use(httputil.RecoveryMiddleware(s.logger))
		// Asset routes
		assets := v1.Group("assets")
		{
			assets.POST("", s.assetHandler.CreateAsset)
			assets.GET("", s.assetHandler.ListAssets)
			assets.GET(":id", s.assetHandler.GetAsset)
			assets.PUT(":id", s.assetHandler.UpdateAsset)
			assets.DELETE(":id", s.assetHandler.DeleteAsset)
		}

		// 	// Measure routes
		// 	measures := v1.Group("/measures")
		// 	{
		// 		measures.POST("", s.measureHandler.Create)
		// 		measures.GET("", s.measureHandler.List)
		// 		measures.GET("/:id", s.measureHandler.Get)
		// 	}

		// 	// Dashboard routes
		// 	dashboards := v1.Group("/dashboards")
		// 	{
		// 		dashboards.POST("", s.dashboardHandler.Create)
		// 		dashboards.GET("", s.dashboardHandler.List)
		// 		dashboards.GET("/:id", s.dashboardHandler.Get)
		// 		dashboards.PUT("/:id", s.dashboardHandler.Update)
		// 		dashboards.DELETE("/:id", s.dashboardHandler.Delete)
		// 	}

		// 	// Widget routes
		// 	widgets := v1.Group("/widgets")
		// 	{
		// 		widgets.POST("", s.widgetHandler.Create)
		// 		widgets.GET("", s.widgetHandler.List)
		// 		widgets.GET("/:id", s.widgetHandler.Get)
		// 		widgets.PUT("/:id", s.widgetHandler.Update)
		// 		widgets.DELETE("/:id", s.widgetHandler.Delete)
		// 	}
	}
}
