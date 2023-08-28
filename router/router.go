package router

import (
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()

	// ユーザー関係のルーティング
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	// echoのインスタンスを返す
	return e
}
