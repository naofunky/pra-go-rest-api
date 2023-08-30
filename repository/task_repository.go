package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// 　タスク関係のDB操作を行うためのインターフェース
type ITasksRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskByID(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpadateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

// タスク関係のDB操作を行うための構造体
type tasksRepository struct {
	db *gorm.DB
}

// repositoryにDBのインスタンスをDIするためのコンストラクタ
func NewTasksRepository(db *gorm.DB) ITasksRepository {
	return &tasksRepository{db}
}

// 全てのタスクを一覧で取得するメソッド
func (tr *tasksRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("create_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}
