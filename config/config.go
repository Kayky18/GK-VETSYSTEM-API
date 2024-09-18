package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// init config
	db, err = InitializeSQlite()
	if err != nil {
		return fmt.Errorf("ERRO TO INITIALIZE SQLITE: %w", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize Logger
	logger = NewLogger(p)
	return logger
}
