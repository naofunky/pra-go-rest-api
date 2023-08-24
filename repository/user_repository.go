package repository

import "go-rest-api/model"

// ユーザー関係のDB操作を行うためのインターフェース
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
