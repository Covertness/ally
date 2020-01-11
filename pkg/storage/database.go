package storage

import (
	"github.com/jinzhu/gorm"

	// import mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var ins *gorm.DB

// InitDB open the database connection
func InitDB(dialect string, args ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, args...)
	if err != nil {
		return nil, err
	}

	// db.LogMode(true)

	ins = db
	return ins, nil
}

// GetDB get the database instance
func GetDB() *gorm.DB {
	return ins
}
