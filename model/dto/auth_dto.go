package dto

type AuthRegisterDto struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
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
	Token string `json:"token"`
}
