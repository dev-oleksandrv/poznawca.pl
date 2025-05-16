package main

import (
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/config"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/backoffice-proxy/handler"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/backoffice-proxy/service"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/repository"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", "err", err)
		os.Exit(1)
	}

	isProductionMode := cfg.AdminProxy.Env == "production"
	if isProductionMode {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.NewPGQLDatabase(database.NewPGQLDatabaseConfig{
		ConnStr:           cfg.Postgres.Url,
		EnableAutoMigrate: cfg.Postgres.AutoMigrateEnabled,
		EnableDebug:       isProductionMode,
	})
	if err != nil {
		slog.Error("cannot connect to database", "err", err)
		os.Exit(1)
	}

	interviewerRepository := repository.NewInterviewerRepository(db)
	interviewerService := service.NewInterviewerService(interviewerRepository)
	interviewerHandler := handler.NewInterviewerHandler(interviewerService)

	router := gin.Default()

	apiGroup := router.Group("/api")

	interviewerGroup := apiGroup.Group("/interviewer")
	{
		interviewerGroup.GET("/list", interviewerHandler.GetList)
		interviewerGroup.GET("/:id", interviewerHandler.GetByID)
		interviewerGroup.POST("/", interviewerHandler.Create)
		interviewerGroup.PUT("/:id", interviewerHandler.Update)
		interviewerGroup.DELETE("/:id", interviewerHandler.Delete)
	}

	if err := router.Run(fmt.Sprintf(":%d", cfg.AdminProxy.Port)); err != nil {
		slog.Error("failed to start server", "err", err)
		os.Exit(1)
	}
}
