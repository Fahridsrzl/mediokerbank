package main

import (
	_ "medioker-bank/docs"
	"medioker-bank/delivery"
)

// @title Tag Service API
// @version 1.0
// @description A Tag service API in Go using Gin framework

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @scope.write Grants write acces
// @description Bearer <token>
func main() {
	delivery.NewServer().Run()
}
