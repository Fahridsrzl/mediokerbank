package delivery

import (
	"fmt"
	"log"

	"medioker-bank/config"
	cMaster "medioker-bank/delivery/controller/master"
	cOther "medioker-bank/delivery/controller/other"
	"medioker-bank/delivery/middleware"
	"medioker-bank/manager"
	oUsecase "medioker-bank/usecase/other"
	"medioker-bank/utils/common"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uc         manager.UseCaseManager
	engine     *gin.Engine
	host       string
	logService common.MyLogger
	auth       oUsecase.AuthUseCase
	jwt        common.JwtToken
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
	authMiddleware := middleware.NewAuthMiddleware(s.jwt)
	rg := s.engine.Group("/api/v1")
	cMaster.NewLoanProductController(s.uc.LoanProductUseCase(), rg).Router()
	cMaster.NewUserController(s.uc.UserUseCase(), rg).Router()
	cOther.NewAuthController(s.auth, rg, s.jwt).Router()
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

	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUseCaseManager(repoManager)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.LogFileConfig)
	jwt := common.NewJwtToken(cfg.TokenConfig)
	mailer := common.NewMailer(cfg.MailerConfig)
	return &Server{
		uc:         usecaseManager,
		engine:     engine,
		host:       host,
		logService: logService,
		auth:       oUsecase.NewAuthUseCase(repoManager.AuthRepo(), jwt, mailer),
		jwt:        jwt,
	}
}
