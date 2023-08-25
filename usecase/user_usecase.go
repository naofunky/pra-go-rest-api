package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"

	"golang.org/x/crypto/bcrypt"
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

// リポジトリをユースケースにDIするためのコンストラクタ
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

// サインイン処理を行うメソッド
func (uu *userUsecase) Signup(user model.User) (model.UserResponse, error) {
	// パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	// 引数で受け取ったemail情報を元にユーザー情報を作成
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	// レスポンス用のユーザー情報を作成
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}
