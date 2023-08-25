package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// ユーザー関係のDB操作を行うためのインターフェース
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// ユーザー関係のDB操作を行うための構造体
type userRepository struct {
	db *gorm.DB
}

// repositoryにDBのインスタンスをDIするためのコンストラクタ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// ユーザー情報を取得するメソッド
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザー情報を作成するメソッド
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
