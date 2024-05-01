package domain

type User struct {
	Id        int     `json:"id" db:"employeeID"`
	FirstName string  `json:"fname" db:"firstname"`
	LastName  string  `json:"lname" db:"lastname"`
	Position  int     `json:"position"`
	Bonus     float64 `json:"bonus"`
	Email     string  `json:"email" db:"email"`
	Password  string  `json:"password" db:"password"`
}

type UserRepository interface {
	Create(newUser *User) error
	GetById(id int) (*User, error)
	GetByEmail(email string)(*User, error)
	Getall()(*User, error)
	DeleteUser(user *User) error
	//GetBymail(email string) *User
}

type UserUseCase interface {
	Create(newUser *User) error
	GetById(id int) (*User,error)
	GetByEmail(email string)(*User, error)
}
