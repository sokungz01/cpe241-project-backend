package repository

import (
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
	_, err := s.db.NamedExec("INSERT INTO `employee` (name,surname,imageURL,bonus,positionID,email,password)"+
		"VALUE (:name,:surname,:imageURL,:bonus,:positionID,:email,:password)",
		newUser)
	if err != nil {
		return err
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

func (s *userRepository) GetByEmail(email string) (*domain.User, error) {
	var response domain.User
	err := s.db.Get(&response, "SELECT * FROM `employee` WHERE `email` = ?", email)
	if err != nil {
		return nil, err
	}
	return &response, nil

}

func (s *userRepository) Getall() (*[]domain.User, error) {
	response := make([]domain.User, 0)
	err := s.db.Select(&response, "SELECT `employeeID`,`name`,`surname`,`imageURL`,`positionID`,`bonus`,`email`"+
		"FROM `employee`")

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *userRepository) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM `employee`"+
		"WHERE `employeeID` = ?", id)
	return err
}

func (s *userRepository) UpdateUser(id int, newUser *domain.User) error {
	_, err := s.db.Exec("UPDATE `employee`"+
		"SET `name` = ?,`surname` = ?, `imageURL` = ?,`positionID` = ?,`bonus`=?,`email` = ? WHERE `employeeID` = ?",
		newUser.FirstName, newUser.LastName, newUser.ImageURL, newUser.Position, newUser.Bonus, newUser.Email, id)
	return err
}
