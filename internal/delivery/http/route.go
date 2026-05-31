package http_delivery

import (
	"log/slog"

	"github.com/Ozenkol/rbk-go-final/internal/delivery/http/handlers" // Force parse
	http_types "github.com/Ozenkol/rbk-go-final/internal/delivery/http/types"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title CRM API
// @version 1.0
// @description Test API
// @host localhost:8081
// @BasePath /
// @schemes http
func registerRoutes(engine *gin.Engine, log *slog.Logger, deps http_types.Dependencies) {
    engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	engine.GET("/health", func(c *gin.Context) {
		log.Info("Health check")
		c.JSON(200, gin.H{"status": "ok"})
	})

	userHandler := handlers.NewUserHandler(&deps)

	v1 := engine.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
		}
		_ = v1 // remove once first route is registered
	}
}
