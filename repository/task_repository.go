package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 　タスク関係のDB操作を行うためのインターフェース
type ITasksRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskByID(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
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
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// タスクをIDで取得するメソッド
func (tr *tasksRepository) GetTaskByID(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

// タスクを作成するメソッド
func (tr *tasksRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// タスクを更新するメソッド
func (tr *tasksRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("user_id=? AND id=?", userId, taskId).Update("title", task.Title)
	if result != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

// タスクを削除するメソッド
func (tr *tasksRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("user_id=? AND id=?", userId, taskId).Delete(&model.Task{})
	if result != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
