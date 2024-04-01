package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	repository "medioker-bank/repository/other"
	"medioker-bank/utils/common"
	"regexp"
	"strconv"
	"time"
)

type AuthUseCase interface {
	RegisterUser(payload dto.AuthRegisterDto) (string, error)
	VerifyUser(code int) (model.User, error)
	LoginUser(payload dto.AuthLoginDto) (dto.AuthResponseDto, error)
	LoginAdmin(payload dto.AuthLoginDto) (dto.AuthResponseDto, error)
}

type authUseCase struct {
	repo   repository.AuthRepository
	jwt    common.JwtToken
	mailer common.Mailer
}

func (a *authUseCase) RegisterUser(payload dto.AuthRegisterDto) (string, error) {
	emailPattern := `[a-z0-9]+@gmail.com`
	validateEmail, _ := regexp.MatchString(emailPattern, payload.Email)
	var err error
	if !validateEmail {
		return "", errors.New("allowed email: 'xxxx@gmail.com'")
	}
	if len(payload.Password) < 6 {
		return "", errors.New("password min 6 characters")
	}
	if payload.Password != payload.ConfirmPassword {
		return "", errors.New("password and confirmPassword do not match")
	}
	payload.Password, err = common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return "", err
	}
	var vCode int
	for {
		codePhase1 := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
		codePhase2 := rand.New(rand.NewSource(int64(codePhase1))).Int() / 1e13
		codeLength := strconv.Itoa(codePhase2)
		if len(codeLength) == 6 {
			vCode = codePhase2
			break
		}
	}
	newPayload := dto.AuthVerifyDto{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		VCode:    vCode,
	}
	message, err := a.repo.CreateQueue(newPayload)
	if err != nil {
		return "", err
	}
	newMail := dto.Mail{
		Receiver: newPayload.Email,
		Subject:  `Medioker Bank Verification Code`,
		Body: fmt.Sprintf(`<h1>Welcome to Medioker Bank</h1> <p>This is your verification code:</p>
		<p style="padding: 10px; background-color: #EEEEEE">%d<p>
		<p>This is very secret, don't tell anyone even your mom</p><p>We always waiting your money :)</p>`, newPayload.VCode),
	}
	err = a.mailer.SendEmail(newMail)
	if err != nil {
		return "", err
	}
	return message, nil
}

func (a *authUseCase) VerifyUser(code int) (model.User, error) {
	queue, err := a.repo.FindByVCode(code)
	if err != nil {
		return model.User{}, err
	}
	user, err := a.repo.CreateUser(queue)
	if err != nil {
		return model.User{}, err
	}
	err = a.repo.DeleteQueue(code)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *authUseCase) LoginUser(payload dto.AuthLoginDto) (dto.AuthResponseDto, error) {
	if payload.Email == "" && payload.Username == "" {
		return dto.AuthResponseDto{}, errors.New("username or email required")
	}
	if payload.Email != "" && payload.Username != "" {
		return dto.AuthResponseDto{}, errors.New("choose between username or email")
	}
	user, err := a.repo.FindUniqueUser(payload)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	err = common.ComparePasswordHash(user.Password, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, errors.New("wrong password")
	}
	accessToken, err := a.jwt.GenerateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	refreshToken, err := a.jwt.GenerateRefreshToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	tokens := dto.AuthResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokens, nil
}

func (a *authUseCase) LoginAdmin(payload dto.AuthLoginDto) (dto.AuthResponseDto, error) {
	if payload.Email == "" && payload.Username == "" {
		return dto.AuthResponseDto{}, errors.New("username or email required")
	}
	if payload.Email != "" && payload.Username != "" {
		return dto.AuthResponseDto{}, errors.New("choose between username or email")
	}
	admin, err := a.repo.FindUniqueAdmin(payload)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	mockUser := model.User{
		ID:   admin.Id,
		Role: admin.Role,
	}
	accessToken, err := a.jwt.GenerateToken(mockUser)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	refreshToken, err := a.jwt.GenerateRefreshToken(mockUser)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	tokens := dto.AuthResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokens, nil
}

func NewAuthUseCase(repo repository.AuthRepository, jwt common.JwtToken, mailer common.Mailer) AuthUseCase {
	return &authUseCase{repo: repo, jwt: jwt, mailer: mailer}
}
