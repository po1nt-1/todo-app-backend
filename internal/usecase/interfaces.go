package usecase

import (
	"context"

	"github.com/po1nt-1/todo-app-backend/internal/entity"
)

type (
	// Task -.
	Task interface {
		GetAllTasks(context.Context) ([]entity.Task, error)
		GetTask(context.Context, uint) (entity.Task, error)
	}
)
