package main

import (
	"medioker-bank/delivery"
	_ "medioker-bank/docs"
)

// @title Tag Service API
// @version 1.0
// @description A Tag service API in Go using Gin framework

// @host 16.78.3.230:8081
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @scope.write Grants write acces
// @description Bearer <token>
func main() {
	delivery.NewServer().Run()
}
