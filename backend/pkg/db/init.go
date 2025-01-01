package db

import (
	"log/slog"

	"github.com/glebarez/sqlite"
	"github.com/ragulmathawa/go-react-auth-app/pkg/utils"
	"gorm.io/gorm"
)

func InitSqlite(appConfig utils.AppConfig) {
	slog.Info(
		"initializing sqlite database",
		"db_file_path", appConfig.DBFilePath,
	)
	db, err := gorm.Open(sqlite.Open(appConfig.DBFilePath), &gorm.Config{})

	if err != nil {
		slog.Error(
			"error initializing sqlite database",
			"db_file_path", appConfig.DBFilePath,
		)
	}
	Conn = db
}

func CloseDB() {
	db, err := Conn.DB()
	if err != nil {
		slog.Error(
			"Error closing DB Connection",
			"err", err,
		)
		return
	}
	err = db.Close()
	if err != nil {
		slog.Error(
			"Error closing DB Connection",
			"err", err,
		)
		return
	}
}
