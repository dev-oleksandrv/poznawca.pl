package database

import (
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type PGQLDatabase struct {
	*gorm.DB
}

type NewPGQLDatabaseConfig struct {
	ConnStr           string
	EnableDebug       bool
	EnableAutoMigrate bool
	MaxIdleConns      int
	MaxOpenConns      int
	ConnMaxLifetime   time.Duration
	ConnMaxIdleTime   time.Duration
}

func NewPGQLDatabase(cfg NewPGQLDatabaseConfig) (*PGQLDatabase, error) {
	gormConfig := &gorm.Config{}

	if cfg.EnableDebug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(cfg.ConnStr), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying SQL database: %w", err)
	}

	if cfg.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}
	if cfg.ConnMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	}

	if err := setupUUIDExtension(db); err != nil {
		return nil, fmt.Errorf("failed to create UUID extension: %w", err)
	}

	if cfg.EnableAutoMigrate {
		if err := autoMigrate(db); err != nil {
			return nil, fmt.Errorf("failed to auto-migrate: %w", err)
		}
	}

	return &PGQLDatabase{DB: db}, nil
}

func setupUUIDExtension(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Interviewer{},
		&model.Interview{},
		&model.InterviewMessage{},
		&model.InterviewResult{},
	)
}
