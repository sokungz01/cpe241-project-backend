package domain

type User struct {
	Id        int     `json:"id" db:"employeeID"`
	FirstName string  `json:"fname" db:"name"`
	LastName  string  `json:"lname" db:"surname"`
	Position  int     `json:"position" db:"positionID"`
	Bonus     float64 `json:"bonus" db:"bonus"`
	Email     string  `json:"email" db:"email"`
	Password  string  `json:"password,omitempty" db:"password"`
}

type UserRepository interface {
	Create(newUser *User) error
	GetById(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	Getall() (*[]User, error)
	DeleteUser(user *User) error
	//GetBymail(email string) *User
}

type UserUseCase interface {
	Create(newUser *User) error
	GetById(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() (*[]User, error)
}
