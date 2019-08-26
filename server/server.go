package server

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

// Server serves as the handler for the graphql and also the handler for all REST API
type Server struct {
	*handler.Handler
}

// NewServer returns new Server Object
func NewServer(h *handler.Handler) *Server {
	return &Server{
		h,
	}
}

// GraphQL pass the gin context to go-handler to execute the graphql query
func (s *Server) GraphQL(c *gin.Context) {
	s.ServeHTTP(c.Writer, c.Request)
}

// TODO populate with REST Controller
