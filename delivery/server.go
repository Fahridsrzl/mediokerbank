package delivery

import (
	"fmt"
	"log"

	"medioker-bank/config"
	cMaster "medioker-bank/delivery/controller/master"
	cOther "medioker-bank/delivery/controller/other"
	cTransaction "medioker-bank/delivery/controller/transaction"
	"medioker-bank/delivery/middleware"
	"medioker-bank/manager"
	uOther "medioker-bank/usecase/other"
	uTransaction "medioker-bank/usecase/transaction"
	"medioker-bank/utils/common"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

type Server struct {
	uc             manager.UseCaseManager
	engine         *gin.Engine
	host           string
	logService     common.MyLogger
	jwt            common.JwtToken
	auth           uOther.AuthUseCase
	installmentTrx uTransaction.InstallmentTransactionUseCase
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
	authMiddleware := middleware.NewAuthMiddleware(s.jwt)
	rg := s.engine.Group("/api/v1")

	cMaster.NewLoanProductController(s.uc.LoanProductUseCase(), rg, authMiddleware).Router()
	cMaster.NewUserController(s.uc.UserUseCase(), rg, authMiddleware).Router()
	cOther.NewAuthController(s.auth, rg, s.jwt).Router()
	cTransaction.NewInstallmentTransactionController(s.installmentTrx, rg, authMiddleware).Router()
	cTransaction.NewTopupController(s.uc.TopupUseCase(), rg, authMiddleware).Router()
	cTransaction.NewTransferController(s.uc.TransferUseCase(), rg, authMiddleware).Router()
	cTransaction.NewLoanTransactionController(s.uc.LoanTransactionUseCase(), rg, authMiddleware).Router()
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
	// add swagger
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.LogFileConfig)
	jwt := common.NewJwtToken(cfg.TokenConfig)
	mailer := common.NewMailer(cfg.MailerConfig)
	bcrypt := common.NewBcryptService()
	midtransService := common.NewMidtransService(cfg.MidtransConfig)
	return &Server{
		uc:             usecaseManager,
		engine:         engine,
		host:           host,
		logService:     logService,
		jwt:            jwt,
		auth:           uOther.NewAuthUseCase(repoManager.AuthRepo(), jwt, mailer, bcrypt),
		installmentTrx: uTransaction.NewInstallmentTransactionUseCase(repoManager.InstallmentTransactionRepo(), repoManager.LoanRepo(), usecaseManager.UserUseCase(), usecaseManager.LoanProductUseCase(), midtransService),
	}
}
