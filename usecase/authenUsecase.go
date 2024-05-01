package usecase

import (
	_"github.com/sokungz01/cpe241-project-backend/repository"
	"github.com/sokungz01/cpe241-project-backend/domain"
	//"golang.org/x/crypto/bcrypt"
	//"net/mail"
	//"errors"	
)

type authenUsecase struct {
	userRepository domain.UserRepository
}