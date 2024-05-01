package domain

type AuthenUsecase interface {
	SignIn (email string, password string) 
}