package main

import (
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/handler"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/service"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/infrastructure/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/config"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/repository"
	"github.com/gin-contrib/cors"
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
	})
	if err != nil {
		slog.Error("cannot connect to database", "err", err)
		os.Exit(1)
	}

	openaiClient := openai.NewClient(cfg.OpenAI.APIKey)

	interviewerRepository := repository.NewInterviewerRepository(db)
	interviewRepository := repository.NewInterviewRepository(db)
	interviewMessageRepository := repository.NewInterviewMessageRepository(db)
	interviewResultRepository := repository.NewInterviewResultRepository(db)

	openaiInterviewService := service.NewAppOpenAIInterviewService(&cfg.OpenAI, openaiClient)
	interviewerService := service.NewAppInterviewerService(interviewerRepository)
	interviewService := service.NewAppInterviewService(interviewRepository, openaiInterviewService)
	wsInterviewService := service.NewAppWSInterviewService(&service.NewAppWSInterviewServiceConfig{
		InterviewRepository:        interviewRepository,
		InterviewMessageRepository: interviewMessageRepository,
		InterviewResultRepository:  interviewResultRepository,
		OpenAIInterviewService:     openaiInterviewService,
	})

	interviewerHandler := handler.NewAppInterviewerHandler(interviewerService)
	interviewHandler := handler.NewAppInterviewHandler(interviewService, interviewerService)
	wsInterviewHandler := handler.NewAppWSInterviewHandler(wsInterviewService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return origin == cfg.WebClient.Url
		},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))
	apiGroup := router.Group("/api")
	interviewerGroup := apiGroup.Group("/interviewer")
	{
		interviewerGroup.GET("/list", interviewerHandler.GetList)
	}
	interviewGroup := apiGroup.Group("/interview")
	{
		interviewGroup.GET("/list", interviewHandler.GetAll)
		interviewGroup.GET("/:id", interviewHandler.GetByID)
		interviewGroup.POST("", interviewHandler.Create)
	}
	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("/interview", wsInterviewHandler.RunInterview)
	}

	if err := router.Run(fmt.Sprintf(":%d", cfg.AppProxy.Port)); err != nil {
		slog.Error("failed to start backoffice-proxy", "err", err)
		os.Exit(1)
	}
}
