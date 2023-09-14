package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// ユーザー関係のコントローラーを行うためのインターフェース
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

// ユーザー関係のコントローラーを行うための構造体
type userController struct {
	uu usecase.IUserUsecase
}

// usecaseをcontrollerにDIするためのコンストラクタ
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// サインイン処理を行うメソッド
func (uc *userController) SignUp(c echo.Context) error {
	// クライアントから受け取ったリクエストbodyを構造体に変換する処理
	user := model.User{}

	// バインディングに失敗した時の処理
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// バインディングに成功した時の処理
	userRes, err := uc.uu.SignUp(user)

	// サインアップ失敗した時の処理
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// サインアップに成功した時の処理
	return c.JSON(http.StatusOK, userRes)
}

// ログイン処理を行うメソッド
func (uc *userController) LogIn(c echo.Context) error {
	// クライアントから受け取ったリクエストbodyを構造体に変換する処理
	user := model.User{}

	// バインディングに失敗した時の処理
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// バインディングに成功した時の処理
	tokenString, err := uc.uu.LogIn(user)

	// ログイン失敗した時の処理
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// ログインに成功した時情報をクッキーに保存する処理
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// ログアウト処理を行うメソッド
func (uc *userController) LogOut(c echo.Context) error {
	// クッキーを削除する処理
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// CSRFトークンを発行するメソッド
func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
