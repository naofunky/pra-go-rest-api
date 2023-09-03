package usecase

import "go-rest-api/model"

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
}
