package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Address string
	Router  *gin.Engine
}

type HandlerDetails struct {
	HttpMethod string
	Handler    gin.HandlerFunc
}

func NewServer(address string, router *gin.Engine) *Server {
	return &Server{
		Address: address,
		Router:  router,
	}
}

func (s *Server) Start() {
	srv := &http.Server{
		Addr:    s.Address,
		Handler: s.Router,
	}
	// TODO:Lógica para iniciar o servidor
	srv.ListenAndServe()
}
