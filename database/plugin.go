package database

import (
	"github.com/allentom/harukap/datasource"
	"gorm.io/gorm"
)

var DefaultDatabasePlugin = datasource.Plugin{
	OnConnected: func(db *gorm.DB) {
		db.AutoMigrate(&Program{})
	},
}
