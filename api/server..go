package api

import (
	"github.com/Himneesh-Kalra/makerble-coding-assessment/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiServer struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func NewApiServer(db *gorm.DB) *ApiServer {
	router := gin.Default()

	server := &ApiServer{
		Router: router,
		DB:     db,
	}

	// Initialize routes
	server.setupRoutes()

	return server
}

func (s *ApiServer) Start(addr string) error {
	return s.Router.Run(addr)
}

func (s *ApiServer) setupRoutes() {
	api := s.Router.Group("/api")

	// Pass router group and DB to route registration functions
	routes.RegisterAuthRoutes(api.Group("/auth"), s.DB)
	routes.RegisterPatientRoutes(api.Group("/patients"), s.DB)
}
