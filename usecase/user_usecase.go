package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// ユーザー関係のユースケースを行うためのインターフェース
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	LogIn(user model.User) (string, error)
}

// ユーザー関係のユースケースを行うための構造体
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// リポジトリをユースケースにDIするためのコンストラクタ
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

// サインイン処理を行うメソッド
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// クライアントから送られてきた情報のバリデーションを行う
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
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

// ログイン処理を行うメソッド
func (uu *userUsecase) LogIn(user model.User) (string, error) {
	//
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	// クライアントから送られてきた情報がDBに存在するか確認
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// クライアントから送られてきたemailが存在する場合はパスワードの検証を行う
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// パスワードが一致する場合はJWTトークンの発行を行う
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
