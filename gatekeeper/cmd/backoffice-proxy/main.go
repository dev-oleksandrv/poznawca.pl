package main

import (
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/backoffice-proxy/handler"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/backoffice-proxy/service"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/infrastructure/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/config"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/repository"
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

	isProductionMode := cfg.BackofficeProxy.Env == "production"
	if isProductionMode {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.NewPGQLDatabase(database.NewPGQLDatabaseConfig{
		ConnStr:           cfg.Postgres.Url,
		EnableAutoMigrate: cfg.Postgres.AutoMigrateEnabled,
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

	if err := router.Run(fmt.Sprintf(":%d", cfg.BackofficeProxy.Port)); err != nil {
		slog.Error("failed to start backoffice-proxy", "err", err)
		os.Exit(1)
	}
}
