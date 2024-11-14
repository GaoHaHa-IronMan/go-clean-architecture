package app

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/api"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *api.Router
}

func (s *Server) Start() {
	s.apiRouter.With(s.engine)
	err := s.engine.Run()
	if err != nil {
		panic(err)
	}
}

func NewServer(engine *gin.Engine, apiRouter *api.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}
