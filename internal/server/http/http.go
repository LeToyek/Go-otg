package http

import (
	// golang package
	"go-otg/internal/app/http/server"
	"net/http"

	// external package
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	handlers *server.Handlers
}

// NewHTTPServer new http server by given Handlers.
//
// It returns pointer of httpServer when successful.
// Otherwise, nil pointer of httpServer will be returned.
func NewHTTPServer(handlers *server.Handlers) *httpServer {
	return &httpServer{
		handlers: handlers,
	}
}

// RegisterHandler register handler by given r pointer of gin.Engine.
func (server *httpServer) RegisterHandler(r *gin.Engine) {
	r.GET("/health-check", server.handlers.Common.HandleGetCommonMessage)
	r.GET("/test", func(ctx *gin.Context) { ctx.JSON(http.StatusAccepted, gin.H{"message": "aowkwoak"}) })
}
