package usecase

import (
	"errors"

	"github.com/sokungz01/cpe241-project-backend/domain"
	_ "github.com/sokungz01/cpe241-project-backend/repository"
	"golang.org/x/crypto/bcrypt"
)

type authenUsecase struct {
	authenRepository domain.AuthenRepository
}

func NewAuthUseCase(authenRepository domain.AuthenRepository) domain.AuthenUsecase {
	return &authenUsecase{authenRepository: authenRepository}
}

func (u *authenUsecase) SignIn(user *domain.AuthenPayload) (*domain.AuthenPayload, error) {
	if err := MailValidator(user.Email); err != nil {
		return nil, errors.New("signin : Email is not valid")
	}

	AuthResponse, err := u.authenRepository.SignIn(user)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(AuthResponse.Password), []byte(user.Password)) != nil {
		return nil, errors.New("signin : Password not match")
	}
	return AuthResponse, nil
}
