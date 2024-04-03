package middlewaremock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type AuthMiddlewareMock struct {
	mock.Mock
}

func (a *AuthMiddlewareMock) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
