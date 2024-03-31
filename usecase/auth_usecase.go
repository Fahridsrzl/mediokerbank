package usecase

import (
	"errors"
	"math/rand"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	"medioker-bank/repository"
	"medioker-bank/utils/common"
	"regexp"
	"time"
)

type AuthUseCase interface {
	RegisterUser(payload dto.AuthRegisterDto) (int, error)
	VerifyUser(code int) (model.User, error)
	LoginUser(payload dto.AuthLoginDto) (dto.AuthResponseDto, error)
	LoginAdmin(payload dto.AuthLoginDto) (dto.AuthResponseDto, error)
}

type authUseCase struct {
	repo repository.AuthRepository
	uc   UserUseCase
	jwt  common.JwtToken
}

func (a *authUseCase) RegisterUser(payload dto.AuthRegisterDto) (int, error) {
	emailPattern := `[a-z0-9]+@gmail.com`
	validateEmail, _ := regexp.MatchString(emailPattern, payload.Email)
	var err error
	if !validateEmail {
		return 0, errors.New("allowed email: 'xxxx@gmail.com'")
	}
	if len(payload.Password) < 6 {
		return 0, errors.New("password min 6 characters")
	}
	if payload.Password != payload.ConfirmPassword {
		return 0, errors.New("password and confirmPassword do not match")
	}
	payload.Password, err = common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return 0, err
	}
	codePhase1 := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	codePhase2 := rand.New(rand.NewSource(int64(codePhase1))).Int() / 1e13
	newPayload := dto.AuthVerifyDto{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		VCode:    codePhase2,
	}
	vCode, err := a.repo.Create(newPayload)
	if err != nil {
		return 0, err
	}
	return vCode, nil
}

func (a *authUseCase) VerifyUser(code int) (model.User, error) {
	queue, err := a.repo.FindByVCode(code)
	if err != nil {
		return model.User{}, err
	}
	user, err := a.uc.Create(queue)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *authUseCase) LoginUser(payload dto.AuthLoginDto) (dto.AuthResponseDto, error) {
	user, err := a.uc.FindByUniqueAndPassword(payload)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	token, err := a.jwt.GenerateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	return token, nil
}

func (a *authUseCase) LoginAdmin(payload dto.AuthLoginDto) (dto.AuthResponseDto, error) {
	user, err := a.repo.FindByUniqueAndPassword(payload)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	token, err := a.jwt.GenerateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	return token, nil
}

func NewAuthUseCase(uc UserUseCase, jwt common.JwtToken) AuthUseCase {
	return &authUseCase{uc: uc, jwt: jwt}
}
