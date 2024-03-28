package delivery

import (
	"fmt"
	"log"

	"medioker-bank/config"
	"medioker-bank/delivery/controller"
	"medioker-bank/delivery/middleware"
	"medioker-bank/manager"
	"medioker-bank/usecase"
	"medioker-bank/utils/common"
	"github.com/gin-gonic/gin"
)

type Server struct {
	uc         manager.UseCaseManager
	auth       usecase.AuthUseCase
	engine     *gin.Engine
	host       string
	logService common.MyLogger
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.LogFileConfig)
	return &Server{
		engine:     engine,
		host:       host,
		logService: logService,
	}
}
