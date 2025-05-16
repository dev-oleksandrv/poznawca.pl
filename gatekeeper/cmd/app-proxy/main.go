package main

import (
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/config"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/handler"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/service"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", "err", err)
		os.Exit(1)
	}

	isProductionMode := cfg.AppProxy.Env == "production"
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

	openaiClient := openai.NewClient(cfg.OpenAI.APIKey)

	interviewRepository := repository.NewInterviewRepository(db)
	interviewerRepository := repository.NewInterviewerRepository(db)

	interviewService := service.NewInterviewService(openaiClient, interviewRepository, interviewerRepository)

	interviewHandler := handler.NewInterviewHandler(interviewService)

	router := gin.Default()

	apiGroup := router.Group("/api")

	interviewGroup := apiGroup.Group("/interview")
	{
		interviewGroup.GET("/:id", interviewHandler.GetByID)
		interviewGroup.POST("/", interviewHandler.Create)
	}

	if err := router.Run(fmt.Sprintf(":%d", cfg.AppProxy.Port)); err != nil {
		slog.Error("failed to start server", "err", err)
		os.Exit(1)
	}
}
