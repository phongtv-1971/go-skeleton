package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
