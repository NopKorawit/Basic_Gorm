package handler

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DB() (*gorm.DB, error) {
	dsn := "server=localhost\\SQLEXPRESS;Database=GormTest;praseTime=true"
	dial := sqlserver.Open(dsn)
	return gorm.Open(dial, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
}
