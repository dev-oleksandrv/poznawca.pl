package main

import (
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/config"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/handler"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/service"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/ws"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/repository"
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
		EnableDebug:       isProductionMode,
	})
	if err != nil {
		slog.Error("cannot connect to database", "err", err)
		os.Exit(1)
	}

	redisDb, err := database.NewRedisDatabase(database.NewRedisDatabaseConfig{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		slog.Error("cannot connect to redis database", "err", err)
		os.Exit(1)
	}

	openaiClient := openai.NewClient(cfg.OpenAI.APIKey)

	interviewRepository := repository.NewInterviewRepository(db)
	interviewerRepository := repository.NewInterviewerRepository(db)
	interviewMessageRepository := repository.NewInterviewMessageRepository(db)
	interviewResultRepository := repository.NewInterviewResultRepository(db)

	interviewCacheRepository := repository.NewInterviewCacheRepository(redisDb)
	fmt.Println(interviewCacheRepository)

	interviewAIService := service.NewInterviewAIService(
		openaiClient,
		&cfg.OpenAI,
	)
	interviewService := service.NewInterviewService(
		openaiClient,
		&cfg.OpenAI,
		interviewAIService,
		interviewRepository,
		interviewerRepository,
		interviewMessageRepository,
		interviewResultRepository,
	)

	interviewHandler := handler.NewInterviewHandler(interviewService)

	wsHandler := ws.NewHandler(interviewService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return origin == cfg.WebClient.Url
		},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("/interview", wsHandler.RunInterview)
	}

	apiGroup := router.Group("/api")

	interviewGroup := apiGroup.Group("/interview")
	{
		interviewGroup.GET("/:id", interviewHandler.GetByID)
		interviewGroup.POST("", interviewHandler.Create)
	}

	if err := router.Run(fmt.Sprintf(":%d", cfg.AppProxy.Port)); err != nil {
		slog.Error("failed to start server", "err", err)
		os.Exit(1)
	}
}
