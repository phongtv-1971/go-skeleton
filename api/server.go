package api

import (
	"github.com/gin-gonic/gin"
	"github.com/phongtv-1971/go-skeleton/constants"
	db "github.com/phongtv-1971/go-skeleton/db/sqlc"
	"net/http"
)

type Server struct {
	store db.Store
	router *gin.Engine
}

func NewServer(store db.Store, environment string) *Server {
	server := &Server{store: store}

	/*Config default Gin*/
	router := gin.Default()

	/* Load static file for api doc */
	if environment != constants.Test {
		router.Static("/assets", "./doc/assets")
		router.StaticFS("/api-docs", http.Dir("./doc/api-docs"))
		router.LoadHTMLGlob("./doc/swagger/*")
		router.GET("/api_doc", server.swaggerDoc)
	}

	/* Health check route */
	router.GET("/health_check", server.healthCheck)

	/* Mapping Api V1 route*/
	v1 := router.Group("/v1")
	{
		v1.POST("/users", server.createUser)
		v1.GET("/users/:id", server.getUser)
		v1.GET("/users", server.listUser)
	}

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
