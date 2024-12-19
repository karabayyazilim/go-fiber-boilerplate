package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Database() *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", Env().DbHost, Env().DbUser, Env().DbPassword, Env().DbName, Env().DbPort, Env().DbSSLMode, Env().DbTimeZone)
		dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			panic("failed to connect database")
		}

		db = dbInstance
	})

	return db
}
