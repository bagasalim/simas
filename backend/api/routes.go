package api

import (
	"github.com/bagasalim/simas/auth"
	"github.com/bagasalim/simas/custom"
	"github.com/bagasalim/simas/managelink"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
	}))
	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	s.Router.POST("/create-account", authHandler.CreateUser)
	s.Router.POST("/login", authHandler.Login)
	//example validation auth route
	s.Router.Use(custom.MiddlewareAuth)
	s.Router.POST("/test", authHandler.Test)

	//manage link
	manageLinkRepo := managelink.NewRepository(s.DB)
	manageLinkService := managelink.NewService(manageLinkRepo)
	manageLinkHandler := managelink.NewHandler(manageLinkService)
	s.Router.GET("/getlink", manageLinkHandler.GetLink)
}
