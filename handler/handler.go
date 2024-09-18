package handler

import (
	"github.com/Kayky18/GK_API/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHanler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
