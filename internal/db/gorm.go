// internal/db/gorm.go
package db

import (
	"fmt"
	"loan-admin/internal/config"
	logger1 "loan-admin/internal/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/plugin/dbresolver"
)

func NewGormDB(cfg *config.Config, log *logger1.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Master.User,
		cfg.Database.Master.Password,
		cfg.Database.Master.Host,
		cfg.Database.Master.Port,
		cfg.Database.Master.Name,
	)

	dsn1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Slave.User,
		cfg.Database.Slave.Password,
		cfg.Database.Slave.Host,
		cfg.Database.Slave.Port,
		cfg.Database.Slave.Name,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	if cfg.Env == "dev" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	//配置读写分离
	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsn)},
		Replicas: []gorm.Dialector{mysql.Open(dsn1)},
		Policy:   dbresolver.RandomPolicy{},
	}))

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.Database.Master.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.Master.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.Master.ConnMaxLifetime)

	return db, nil
}
