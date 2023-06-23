package usecase

import (
	"github.com/po1nt-1/todo-app-backend/internal/entity"
)

type Task interface {
	Get(uint) (entity.Task, error)
}
