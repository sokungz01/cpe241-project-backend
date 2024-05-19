package domain

type User struct {
	Id        int     `json:"id" db:"employeeID"`
	FirstName string  `json:"name" db:"name"`
	LastName  string  `json:"surname" db:"surname"`
	ImageURL  string  `json:"imageURL" db:"imageURL"`
	Position  int     `json:"positionID" db:"positionID"`
	Bonus     float64 `json:"bonus" db:"bonus"`
	Email     string  `json:"email" db:"email"`
	Password  string  `json:"password,omitempty" db:"password"`
}

type UserRepository interface {
	Create(newUser *User) error
	GetById(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	Getall() (*[]User, error)
	DeleteUser(id int) error
	UpdateUser(id int, newUser *User) error
	//GetBymail(email string) *User
}

type UserUseCase interface {
	Create(newUser *User) error
	GetById(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() (*[]User, error)
	UpdateUser(id int, newUser *User) (*User, error)
	DeleteUser(id int) error
}
