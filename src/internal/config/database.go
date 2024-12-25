package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

var (
	db    *gorm.DB
	mutex sync.Mutex
	once  sync.Once
)

func Database() *gorm.DB {
	mutex.Lock()
	defer mutex.Unlock()

	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			Env().DbHost, Env().DbUser, Env().DbPassword, Env().DbName, Env().DbPort, Env().DbSSLMode, Env().DbTimeZone)

		dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("failed to connect to the database: %v", err)
		}

		db = dbInstance

		sqlDB, err := dbInstance.DB()
		if err != nil {
			log.Fatalf("failed to get sql DB instance: %v", err)
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)
	})

	return db
}
