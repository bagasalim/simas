package api

import (
	"github.com/bagasalim/simas/auth"
	"github.com/bagasalim/simas/custom"
	"github.com/bagasalim/simas/managelink"
	"github.com/bagasalim/simas/zoomhistory"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders: []string{"*"},
	}))
	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	s.Router.POST("/create-account", authHandler.CreateUser)
	s.Router.POST("/login", authHandler.Login)
	//example validation auth route

	middleware := custom.MiddleWare{}
	authRoute := s.Router.Group("")
	authRoute.Use(middleware.Auth)

	manageLinkRepo := managelink.NewRepository(s.DB)
	manageLinkService := managelink.NewService(manageLinkRepo)
	manageLinkHandler := managelink.NewHandler(manageLinkService)

	csRoute := authRoute.Group("")
	csRoute.Use(middleware.IsCS)
	csRoute.PUT("/updatelink", manageLinkHandler.UpdateLink)
	csRoute.GET("/getlink", manageLinkHandler.GetLink)

	s.Router.GET("/get-link/:type", manageLinkHandler.GetLinkRequest)
	zoomHistoryRepo := zoomhistory.NewRepository(s.DB)
	zoomHistoryService := zoomhistory.NewService(zoomHistoryRepo)
	zoomHistoryHandler := zoomhistory.NewHandler(zoomHistoryService)
	s.Router.POST("/createzoomhistory", zoomHistoryHandler.CreateZoom)
	// s.Router.Use(custom.MiddlewareAuth)

	// s.Router.POST("/test", authHandler.Test)

	//manage link

	// s.Router.PUT("/updatelink", manageLinkHandler.UpdateLink)
	// s.Router.GET("/getlink", manageLinkHandler.GetLink)
}
