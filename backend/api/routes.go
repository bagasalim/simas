package api

import (
	"github.com/bagasalim/simas/socket_system"
	"github.com/bagasalim/simas/todos"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
	}))

	rtRepo := socket_system.NewRepository(s.DB)
	rtService := socket_system.NewService(rtRepo)
	rtHandler := socket_system.NewHandler(rtService)

	s.Router.GET("/ws/:room", rtHandler.ConnectRoom)

	todosRepo := todos.NewRepository(s.DB)
	todosService := todos.NewService(todosRepo)
	todosHandler := todos.NewHandler(todosService)

	s.Router.GET("/", todosHandler.GetTodos)
	s.Router.POST("/send", todosHandler.CreateTodo)
	// s.Router.GET("/ws", )
}
