package usecase

import(
	_"github.com/sokungz01/cpe241-project-backend/repository"
	"github.com/sokungz01/cpe241-project-backend/domain"
	"net/mail"
	"errors"
)

type userUseCase struct{
	userRepository domain.UserRepository
}

func NewUserUseCase (userRepository domain.UserRepository) (domain.UserUseCase){
	return &userUseCase{userRepository : userRepository}
}

func (u *userUseCase)Create(newUser *domain.User) error{
	//Validate the email format 
	if err := MailValidator(newUser.Email); err != nil{
		return err
	}
	//Validate that the E-mail not already in the database
	if _,err := u.GetByEmail(newUser.Email);err == nil {
		return errors.New("user : already used email")
	}
	u.userRepository.Create(newUser)
	return nil
}

func (u *userUseCase)GetById(id int) (*domain.User,error){
	resp,err := u.userRepository.GetById(id)
	if err != nil {
		return nil,err
	}
	return resp,nil
}

func (u *userUseCase)GetByEmail(email string)(*domain.User, error){
	resp,err := u.userRepository.GetByEmail(email)
	if err != nil {
		return nil,err
	}
	return resp,nil
}

func MailValidator (address string) error {
	_,err := mail.ParseAddress(address)
	return err
}

