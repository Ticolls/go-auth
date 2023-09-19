package handler

import (
	"github.com/Ticolls/go-auth/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
