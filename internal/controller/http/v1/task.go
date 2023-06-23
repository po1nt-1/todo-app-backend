package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/po1nt-1/todo-app-backend/internal/entity"
	"github.com/po1nt-1/todo-app-backend/internal/usecase"
)

type taskRoutes struct {
	t usecase.Task
}

func newTaskRoutes(handler *gin.RouterGroup, t usecase.Task) {
	r := &taskRoutes{t}

	fmt.Println(handler, t)

	h := handler.Group("/task")
	{
		h.GET("/", r.get)
	}
}

type taskResponse struct {
	Task entity.Task `json:"task"`
}

func (r *taskRoutes) get(c *gin.Context) {
	task, err := r.t.Get(0)
	if err != nil {
		log.Fatal(err, "http v1 get")
	}
	c.JSON(http.StatusOK, taskResponse{task})
}
