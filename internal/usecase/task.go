package usecase

import (
	"github.com/po1nt-1/todo-app-backend/internal/entity"
)

type TaskUseCase struct {
}

func New() *TaskUseCase {
	return &TaskUseCase{}
}

func (uc *TaskUseCase) Get(ID uint) (entity.Task, error) {
	task := entity.Task{
		ID:        0,
		Details:   "bla-bla-bla task",
		Completed: false,
	}

	return task, nil
}
