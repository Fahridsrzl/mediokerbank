package common

import (
	"errors"
	"time"

	"medioker-bank/config"
	"medioker-bank/model"

	modelutil "medioker-bank/utils/model_util"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken interface {
	GenerateToken(payload model.User) (string, error)
	GenerateRefreshToken(payload model.User) (string, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
	RefreshToken(oldTokenString string) (string, error)
}

type jwtToken struct {
	cfg config.TokenConfig
}

func (j *jwtToken) RefreshToken(oldRefreshToken string) (string, error) {
	token, err := jwt.Parse(oldRefreshToken, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["iss"] != j.cfg.IssuerName {
		return "", errors.New("invalid claim token")
	}

	if claims["tokenType"] != "refresh token" {
		return "", errors.New("invalid token type")
	}

	claims["tokenType"] = "access token"
	claims["exp"] = float64(time.Now().Add(1*time.Hour).UTC().UnixNano() / 1e9)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return "", err
	}

	return newTokenString, nil
}

func (j *jwtToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return nil, errors.New("invalid claim token")
	}

	return claims, nil
}

func (j *jwtToken) GenerateToken(payload model.User) (string, error) {
	claims := modelutil.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.AccessTokenLifeTime)),
		},
		UserId:    payload.ID,
		Role:      payload.Role,
		TokenType: "access token",
	}

	jwtNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtNewClaims.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return "", errors.New("failed to generate token: " + err.Error())
	}
	return token, nil
}

func (j *jwtToken) GenerateRefreshToken(payload model.User) (string, error) {
	claims := modelutil.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.RefreshTokenLifeTime)),
		},
		UserId:    payload.ID,
		Role:      payload.Role,
		TokenType: "refresh token",
	}

	jwtNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtNewClaims.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return "", errors.New("failed to generate token: " + err.Error())
	}
	return token, nil
}

func NewJwtToken(cfg config.TokenConfig) JwtToken {
	return &jwtToken{cfg: cfg}
}
