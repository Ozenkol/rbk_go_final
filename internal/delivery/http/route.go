package http_delivery

import (
	"fmt"
	"log/slog"

	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/Ozenkol/rbk-go-final/internal/delivery/http/handlers"
	http_middleware "github.com/Ozenkol/rbk-go-final/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func registerRoutes(engine *gin.Engine, log *slog.Logger, deps http_deps.Dependencies) {
	// serve spec
	engine.GET("/swagger.json", func(c *gin.Context) {
		c.File("./api/swagger.json")
	})

	// serve UI locally - no mixed content issue
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/swagger.json"),
	))

	engine.GET("/ping", func(c *gin.Context) {
		log.Info("Ping received")
		c.JSON(200, gin.H{"message": "pong"})
	})

	engine.GET("/health", func(c *gin.Context) {
		log.Info("Health check")
		c.JSON(200, gin.H{"status": "ok"})
	})

	engine.GET("/debug", func(c *gin.Context) {
		var headers []string
		log.Info("Debug endpoint called")
		for k, v := range c.Request.Header {
			headers = append(headers, fmt.Sprintf("%s: %s", k, v))
		}
		c.JSON(200, gin.H{"ok": true, "headers": headers})
	})

	userHandler := handlers.NewUserHandler(&deps, log)
	authHandler := handlers.NewAuthHandler(&deps, log)
	documentHandler := handlers.NewDocumentHandler(&deps, log)
	noteHandler := handlers.NewNoteHandler(&deps, log)
	offerHandler := handlers.NewOfferHandler(&deps, log)
	taskHandler := handlers.NewTaskHandler(&deps, log)
	clientHandler := handlers.NewClientHandler(&deps, log)
	analyticHandler := handlers.NewAnalyticHandler(&deps, log)
	commHandler := handlers.NewCommunicationHandler(&deps, log)
	companyHandler := handlers.NewCompanyHandler(&deps, log)
	contractHandler := handlers.NewContractHandler(&deps, log)
	meetingHandler := handlers.NewMeetingHandler(&deps, log)
	notificationHandler := handlers.NewNotificationHandler(&deps, log)
	fileHandler := handlers.NewFileHandler(&deps, log)
	invoiceHandler := handlers.NewInvoiceHandler(&deps, log)
	productHandler := handlers.NewProductHandler(&deps, log)
	settingHandler := handlers.NewSettingHandler(&deps, log)
	tagHandler := handlers.NewTagHandler(&deps, log)

	authMiddleware := http_middleware.NewAuthMiddleware(&deps, log)

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
			auth.POST("/login", authHandler.LoginUser)
		}

		v1.Use(authMiddleware.MiddlewareFunc())

		// Offers
		offers := v1.Group("/offers")
		{
			offers.POST("", offerHandler.CreateOffer)
			offers.GET("/:id", offerHandler.GetOffer)
			offers.PUT("/:id", offerHandler.UpdateOffer)
			offers.DELETE("/:id", offerHandler.DeleteOffer)
		}

		// Notes
		notes := v1.Group("/notes")
		{
			notes.POST("", noteHandler.CreateNote)
			notes.GET("/:id", noteHandler.GetNote)
			notes.PUT("/:id", noteHandler.UpdateNote)
			notes.DELETE("/:id", noteHandler.DeleteNote)
		}

		// Documents
		documents := v1.Group("/documents")
		{
			documents.POST("", documentHandler.CreateDocument)
			documents.GET("/:id", documentHandler.GetDocument)
			documents.PUT("/:id", documentHandler.UpdateDocument)
			documents.DELETE("/:id", documentHandler.DeleteDocument)
		}

		// Tasks
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("/:id", taskHandler.GetByID)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}

		// Clients
		clients := v1.Group("/clients")
		{
			clients.POST("", clientHandler.CreateClient)
			clients.GET("", clientHandler.ListClients)
			clients.GET("/:id", clientHandler.GetClient)
			clients.PUT("/:id", clientHandler.UpdateClient)
			clients.DELETE("/:id", clientHandler.DeleteClient)
		}

		// Analytics
		analytics := v1.Group("/analytics")
		{
			analytics.POST("", analyticHandler.Create)
			analytics.GET("", analyticHandler.List)
			analytics.GET("/:id", analyticHandler.GetByID)
			analytics.PUT("/:id", analyticHandler.Update)
			analytics.DELETE("/:id", analyticHandler.Delete)
		}

		// Communications
		communications := v1.Group("/communications")
		{
			communications.POST("", commHandler.Create)
			communications.GET("", commHandler.List)
			communications.GET("/:id", commHandler.GetByID)
			communications.PUT("/:id", commHandler.Update)
			communications.DELETE("/:id", commHandler.Delete)
		}

		// Companies
		companies := v1.Group("/companies")
		{
			companies.POST("", companyHandler.Create)
			companies.GET("", companyHandler.List)
			companies.GET("/:id", companyHandler.GetByID)
			companies.PUT("/:id", companyHandler.Update)
			companies.DELETE("/:id", companyHandler.Delete)
		}

		// Contracts
		contracts := v1.Group("/contracts")
		{
			contracts.POST("", contractHandler.Create)
			contracts.GET("", contractHandler.List)
			contracts.GET("/:id", contractHandler.GetByID)
			contracts.PUT("/:id", contractHandler.Update)
			contracts.DELETE("/:id", contractHandler.Delete)
		}

		// Meetings
		meetings := v1.Group("/meetings")
		{
			meetings.POST("", meetingHandler.Create)
			meetings.GET("", meetingHandler.List)
			meetings.GET("/:id", meetingHandler.GetByID)
			meetings.PUT("/:id", meetingHandler.Update)
			meetings.DELETE("/:id", meetingHandler.Delete)
		}

		// Notifications
		notifications := v1.Group("/notifications")
		{
			notifications.POST("", notificationHandler.Create)
			notifications.GET("", notificationHandler.List)
			notifications.GET("/:id", notificationHandler.GetByID)
			notifications.PUT("/:id", notificationHandler.Update)
			notifications.DELETE("/:id", notificationHandler.Delete)
		}

		// Files
		files := v1.Group("/files")
		{
			files.POST("", fileHandler.Create)
			files.GET("", fileHandler.List)
			files.GET("/:id", fileHandler.GetByID)
			files.PUT("/:id", fileHandler.Update)
			files.DELETE("/:id", fileHandler.Delete)
		}

		// Invoices
		invoices := v1.Group("/invoices")
		{
			invoices.POST("", invoiceHandler.Create)
			invoices.GET("", invoiceHandler.List)
			invoices.GET("/:id", invoiceHandler.GetByID)
			invoices.PUT("/:id", invoiceHandler.Update)
			invoices.DELETE("/:id", invoiceHandler.Delete)
		}

		// Products
		products := v1.Group("/products")
		{
			products.POST("", productHandler.Create)
			products.GET("", productHandler.List)
			products.GET("/:id", productHandler.GetByID)
			products.PUT("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
		}

		// Settings
		settings := v1.Group("/settings")
		{
			settings.POST("", settingHandler.Create)
			settings.GET("", settingHandler.List)
			settings.GET("/:id", settingHandler.GetByID)
			settings.PUT("/:id", settingHandler.Update)
			settings.DELETE("/:id", settingHandler.Delete)
		}

		// Tags
		tags := v1.Group("/tags")
		{
			tags.POST("", tagHandler.Create)
			tags.GET("", tagHandler.List)
			tags.GET("/:id", tagHandler.GetByID)
			tags.PUT("/:id", tagHandler.Update)
			tags.DELETE("/:id", tagHandler.Delete)
		}
	}
}
