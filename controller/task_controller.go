package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// タスク関係のコントローラーを行うためのインターフェース
type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskByID(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

// タスク関係のコントローラーを行うための構造体
type taskController struct {
	tu usecase.ITaskUsecase
}

// usecaseをcontrollerにDIするためのコンストラクタ
func NewTaskContlorller(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

// 全てのタスクを一覧で取得するメソッド
func (tc *taskController) GetAllTasks(c echo.Context) error {
	// ユーザー情報の取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// タスクの取得エラーハンドリング
	tasksRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON((http.StatusInternalServerError), err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}

// タスクをIDで取得するメソッド
func (tc *taskController) GetTaskByID(c echo.Context) error {
	// ユーザー情報の取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// タスクIDの取得と型変換
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// タスクの取得とエラーハンドリング
	taskRes, err := tc.tu.GetTaskByID(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

// タスクを作成するメソッド
func (tc *taskController) CreateTask(c echo.Context) error {
	// ユーザー情報の取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// リクエストデータのバインドとエラーハンドリング
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ユーザーIDの設定
	task.UserID = uint(userId.(float64))

	// タスクの作成とエラーハンドリング
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

// タスクを更新するメソッド
func (tc *taskController) UpdateTask(c echo.Context) error {
	// ユーザー情報の取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// タスクIDの取得と型変換
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// リクエストデータのバインドとエラーハンドリング
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// タスクの更新とエラーハンドリング
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

// タスクを削除するメソッド
func (tc *taskController) DeleteTask(c echo.Context) error {
	// ユーザー情報の取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// タスクIDの取得と型変換
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// タスクの削除とエラーハンドリング
	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
