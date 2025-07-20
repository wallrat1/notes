package server

import (
	"auth/internal/config"
	"auth/internal/handler"
	"auth/internal/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
}

func NewServer(cfg *config.Config) (*Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("Пустая конфигурация")
	}
	handler := handler.NewHandler(cfg)
	if handler == nil {
		return nil, fmt.Errorf("no handler")
	}
	router := routes.SetupRoutes(handler)
	if router == nil {
		return nil, fmt.Errorf("no router")
	}
	return &Server{
		cfg:    cfg,
		router: router,
	}, nil

}
func (s *Server) ServerStop() error {
	fmt.Println("server stopped")
	return nil
}
func (s *Server) Serve() error {
	addres := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	fmt.Sprintf("server is ready %s \n", addres)
	return s.router.Run(addres)
}
