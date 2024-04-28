package repository

import (
	"fmt"

	"github.com/sokungz01/cpe241-project-backend/domain"
	"github.com/sokungz01/cpe241-project-backend/platform"
)

type userRepository struct {
	db *platform.Mysql
}

func NewUSerRepository(db *platform.Mysql) domain.UserRepository {
	return &userRepository{db: db}
}

func (s *userRepository) Create(newUser *domain.User) error {
	_,err := s.db.NamedExec("INSERT INTO `employee` (firstname,lastname,email,password)"+
							"VALUE (:firstname,:lastname,:email,:password)",
							newUser)
	if err != nil{
		return nil
	}
	return nil
}

func (s *userRepository) GetById(id int) (*domain.User, error) {
	var response domain.User
	if err := s.db.Get(&response, "SELECT * FROM `employee` WHERE `employeeID` = ?", id); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *userRepository) GetByEmail(email string) (*domain.User, error){
	var response domain.User
	err := s.db.Get(&response,	"SELECT *" + 
								"FROM `employee`" +
								"WHERE `email` = ?",email)
	if err != nil {
		return nil,err
	}
	return &response,nil

}

func (s *userRepository) Getall()(*domain.User, error) {
	var response domain.User
	err := s.db.Get(&response,	"SELECT *" + 
								"FROM `employee`")
	if err != nil {
		return nil,err
	}
	fmt.Println(response)
	return nil,nil
}