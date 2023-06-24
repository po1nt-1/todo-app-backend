package v1

import (
	"github.com/gin-gonic/gin"
	_ "github.com/po1nt-1/todo-app-backend/docs"
	"github.com/po1nt-1/todo-app-backend/internal/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter -.
// Swagger spec:
//
//	@title			Todo-app API
//	@description	using a task
//	@version		1.0
//	@host			127.0.0.1:8080
//	@BasePath		/api/v1
func NewRouter(handler *gin.Engine, t usecase.Task) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	h := handler.Group("/api/v1")
	{
		newTaskRoutes(h, t)
	}
}
