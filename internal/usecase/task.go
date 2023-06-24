package usecase

import (
	"context"
	"fmt"

	"github.com/po1nt-1/todo-app-backend/internal/entity"
)

// TaskUseCase -.
type TaskUseCase struct {
}

// New -.
func New() *TaskUseCase {
	return &TaskUseCase{}
}

// GetAllTasks - get a list of tasks.
func (uc *TaskUseCase) GetAllTasks(ctx context.Context) ([]entity.Task, error) {
	testTasks := make([]entity.Task, 4)
	for i := 0; i < len(testTasks); i++ {
		testTasks[i] = entity.Task{
			ID:        uint(i),
			Details:   fmt.Sprintf("my homework %v", i),
			Completed: false,
		}
	}

	return testTasks, nil
}

// GetTask - get a tasks.
func (uc *TaskUseCase) GetTask(ctx context.Context, ID uint) (entity.Task, error) {
	testTask := entity.Task{
		ID:        ID,
		Details:   fmt.Sprintf("my homework %v", ID),
		Completed: false,
	}

	return testTask, nil
}
