package http_delivery

import (
	"log/slog"

	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/Ozenkol/rbk-go-final/internal/delivery/http/handlers"
	http_middleware "github.com/Ozenkol/rbk-go-final/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			CRM API
// @version		1.0
// @description	Test API
// @host			localhost:8081
// @BasePath		/
// @schemes		http
func registerRoutes(engine *gin.Engine, log *slog.Logger, deps http_deps.Dependencies) {
	// serve spec
	engine.GET("/swagger.json", func(c *gin.Context) {
		c.File("./api/swagger.json")
	})

	// serve UI locally - no mixed content issue
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/swagger.json"),
	))


	engine.GET("/ping", handlers.Ping)

	engine.GET("/health", func(c *gin.Context) {
		log.Info("Health check")
		c.JSON(200, gin.H{"status": "ok"})
	})

	userHandler := handlers.NewUserHandler(&deps, log)
	authHandler := handlers.NewAuthHandler(&deps, log)
	documentHandler := handlers.NewDocumentHandler(&deps, log)
	noteHandler := handlers.NewNoteHandler(&deps, log)
	offerHandler := handlers.NewOfferHandler(&deps, log)
	taskHandler := handlers.NewTaskHandler(&deps, log)

	authMiddleware := http_middleware.NewAuthMiddleware(&deps)

	v1 := engine.Group("/api/v1")
	{
		// Users
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
		}

		// Auth
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.RegisterUser)
		}

		// Offers
		offers := v1.Group("/offers")
		{
			offers.Use(authMiddleware.MiddlewareFunc())
			offers.POST("", offerHandler.CreateOffer)
			offers.GET("/:id", offerHandler.GetOffer)
			offers.PUT("/:id", offerHandler.UpdateOffer)
			offers.DELETE("/:id", offerHandler.DeleteOffer)
		}

		// Notes
		notes := v1.Group("/notes")
		{
			notes.Use(authMiddleware.MiddlewareFunc())
			notes.POST("", noteHandler.CreateNote)
			notes.GET("/:id", noteHandler.GetNote)
			notes.PUT("/:id", noteHandler.UpdateNote)
			notes.DELETE("/:id", noteHandler.DeleteNote)
		}

		// Documents
		documents := v1.Group("/documents")
		{
			documents.Use(authMiddleware.MiddlewareFunc())
			documents.POST("", documentHandler.CreateDocument)
			documents.GET("/:id", documentHandler.GetDocument)
			documents.PUT("/:id", documentHandler.UpdateDocument)
			documents.DELETE("/:id", documentHandler.DeleteDocument)
		}
		
		// Tasks
		tasks := v1.Group("/tasks")
		{
			tasks.Use(authMiddleware.MiddlewareFunc())
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("/:id", taskHandler.GetByID)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}

		// _ = v1 // remove once first route is registered
		
	}
}
