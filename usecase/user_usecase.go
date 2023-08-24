package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

// ユーザー関係のユースケースを行うためのインターフェース
type IUserUsecase interface {
	Signup(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

// ユーザー関係のユースケースを行うための構造体
type userUsecase struct {
	ur repository.IUserRepository
}
