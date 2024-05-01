package domain

type AuthenPayload struct {
	ID       int    `json:"id" db:"employeeID"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
	Position int    `json:"positionID,omitempty" db:"positionID"`
}

type AuthenResponse struct {
	Token string `json:"token"`
}

type AuthenUsecase interface {
	SignIn(user *AuthenPayload) (*AuthenPayload, error)
}

type AuthenRepository interface {
	SignIn(user *AuthenPayload) (*AuthenPayload, error)
}
