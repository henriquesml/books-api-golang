package server

import (
	"log"

	"github.com/gin-gonic/gin"
	router "github.com/henriquesml/go-api/server/routers"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := router.ConfigRoutes(s.server)
	log.Print("Server is running on port", s.port)
	log.Fatal(router.Run(":" + s.port))
}
