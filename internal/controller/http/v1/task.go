package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/po1nt-1/todo-app-backend/internal/entity"
	"github.com/po1nt-1/todo-app-backend/internal/usecase"
)

type taskRoutes struct {
	t usecase.Task
}

func newTaskRoutes(handler *gin.RouterGroup, t usecase.Task) {
	r := &taskRoutes{t}

	h := handler.Group("/tasks")
	{
		h.GET("/", r.getAllTasks)
		h.GET("/:id", r.getTask)
	}
}

type getAllTasksResponse struct {
	Task []entity.Task `json:"tasks"`
}

// @Summary		Get all tasks
// @Description	get a list of tasks
// @ID				getAllTasks
// @Tags			task
// @Accept			json
// @Produce		json
// @Success		200	{object}	getAllTasksResponse
// @Failure		500	{object}	response
// @Router			/tasks [get]
func (r *taskRoutes) getAllTasks(ctx *gin.Context) {
	tasks, err := r.t.GetAllTasks(ctx)
	if err != nil {
		log.Printf("http v1 getAllTasks: %v", err)
		errorResponse(ctx, http.StatusInternalServerError, "tasks service problems")
		return
	}

	ctx.JSON(http.StatusOK, getAllTasksResponse{tasks})
}

type getTasksResponse struct {
	Task entity.Task `json:"task"`
}

// @Summary		Get a task
// @Description	get a task by id
// @ID				getTask
// @Tags			task
// @Accept			json
// @Produce		json
// @Success		200	{object}	getAllTasksResponse
// @Failure		500	{object}	response
// @Router			/tasks/:id [get]
func (r *taskRoutes) getTask(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Printf("http v1 getTasks: %v", err)
		errorResponse(ctx, http.StatusInternalServerError, "bad id")
		return
	}

	tasks, err := r.t.GetTask(ctx, uint(ID))
	if err != nil {
		log.Printf("http v1 getTasks: %v", err)
		errorResponse(ctx, http.StatusInternalServerError, "getTasks service problems")
		return
	}

	ctx.JSON(http.StatusOK, getTasksResponse{tasks})
}
