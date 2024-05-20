// ย่ากานต์ชื่อศรีสุนทร
package usecase

import (
	"errors"
	"net/mail"

	"github.com/sokungz01/cpe241-project-backend/domain"
	_ "github.com/sokungz01/cpe241-project-backend/repository"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) domain.UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (u *userUseCase) Create(newUser *domain.User) error {
	//Validate the email format
	if err := MailValidator(newUser.Email); err != nil {
		return err
	}
	//Validate that the E-mail not already in the database
	if _, err := u.GetByEmail(newUser.Email); err == nil {
		return errors.New("user : already used email")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)
	if err := u.userRepository.Create(newUser); err != nil {
		return err
	}

	//return nil for no error
	return nil
}

func (u *userUseCase) GetById(id int) (*domain.User, error) {
	resp, err := u.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *userUseCase) GetByEmail(email string) (*domain.User, error) {
	resp, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *userUseCase) GetAll() (*[]domain.User, error) {
	resp, err := u.userRepository.Getall()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *userUseCase) UpdateUser(id int, newUser *domain.User) (*domain.User, error) {
	if err := u.userRepository.UpdateUser(id, newUser); err != nil {
		return nil, err
	}
	response, err := u.GetById(id)
	response.Password = ""
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *userUseCase) DeleteUser(id int) error {
	return u.userRepository.DeleteUser(id)
}

func MailValidator(address string) error {
	_, err := mail.ParseAddress(address)
	return err
}
