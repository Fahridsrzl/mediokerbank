package middleware

import (
	"net/http"
	"strings"

	"medioker-bank/utils/common"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService common.JwtToken
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var aH authHeader
		if err := ctx.ShouldBindHeader(&aH); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": err.Error()})
			return
		}

		tokenString := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		claims, err := a.jwtService.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": err.Error()})
			return
		}

		if claims["tokenType"] != "access token" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden Token Type"})
			return
		}

		var validRole bool
		for _, role := range roles {
			if role == claims["role"] {
				validRole = true
				break
			}
		}

		if !validRole {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden Resource"})
			return
		}

		ctx.Set("role", claims["role"])
		ctx.Set("id", claims["userId"])

		ctx.Next()
	}
}

func NewAuthMiddleware(jwtService common.JwtToken) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
