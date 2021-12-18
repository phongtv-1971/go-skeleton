package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) swaggerDoc(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Go base Api document",
	})
}
