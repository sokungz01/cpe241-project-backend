package domain

type AuthenPayload struct {
	Id       int    `json:"id" db:"employeeID"`
	Email    string `json:"email" db:"email"`
	Position int    `json:"positionID" db:"positionID"`
	Password string `json:"password,omitempty" db:"password"`
}

type AuthenDetail struct {
	Id        int    `json:"id" db:"employeeID"`
	FirstName string `json:"name" db:"name"`
	LastName  string `json:"surname" db:"surname"`
	ImageURL  string `json:"imageURL" db:"imageURL"`
	Position  int    `json:"positionID" db:"positionID"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password,omitempty" db:"password"`
}

type AuthenResponse struct {
	Data  AuthenDetail `json:"data"`
	Token string       `json:"token,omitempty"`
}

type AuthenUsecase interface {
	SignIn(user *AuthenPayload) (*AuthenPayload, error)
	Me(userID int) (*AuthenDetail, error)
}

type AuthenRepository interface {
	SignIn(user *AuthenPayload) (*AuthenPayload, error)
	Me(userID int) (*AuthenDetail, error)
}
