package dto

import "time"

type AuthRegisterDto struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type AuthVcodeDto struct {
	VCode int `json:"vCode" binding:"required"`
}

type AuthVerifyDto struct {
	Username string
	Email    string
	Password string
	VCode    int
}

type AuthLoginDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponseDto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Admin struct {
	Id        string
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
