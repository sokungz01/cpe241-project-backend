package domain

type AuthenPayload struct {
	ID       int    `json:"id" db:"employeeID"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
	Position int    `json:"positionID,omitempty" db:"positionID"`
}

type AuthenDetail struct {
	Id        int    `json:"id" db:"employeeID"`
	FirstName string `json:"fname" db:"name"`
	LastName  string `json:"lname" db:"surname"`
	Position  int    `json:"position" db:"positionID"`
}

type AuthenResponse struct {
	Token string `json:"token"`
}

type AuthenUsecase interface {
	SignIn(user *AuthenPayload) (*AuthenPayload, error)
	Me(userID int) (*AuthenDetail, error)
}

type AuthenRepository interface {
	SignIn(user *AuthenPayload) (*AuthenPayload, error)
	Me(userID int) (*AuthenDetail, error)
}
