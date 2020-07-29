package data

import (
	"elephant/config"
	"elephant/log"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	// db drivers
	_ "github.com/go-sql-driver/mysql"           // mysql dialect
	_ "github.com/jinzhu/gorm/dialects/postgres" //psql dialect
	_ "github.com/lib/pq"
)

var (
	l   = log.GetLogger()
	err error
	db  *gorm.DB

	cg = config.GetConfig()
)

// Auto Initiate db connection
func init() {

	// Set environment
	env := cg.GetString("app.environment")

	dbDriver := cg.GetString("database.driver")
	dbHost := cg.GetString("database.host")
	dbPort := cg.GetString("database.port")
	dbName := cg.GetString("database.dbname")
	dbUser := cg.GetString("database.user")
	dbPass := cg.GetString("database.pass")

	switch driver := dbDriver; driver {
	case "postgres":
		_uRL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			dbUser, dbPass, dbHost, dbName,
		)

		db, err = gorm.Open(dbDriver, _uRL)
	case "mysql":
		db, err = gorm.Open(dbDriver,
			dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	}

	db.DB().SetMaxOpenConns(25)
	db.DB().SetMaxIdleConns(25)
	db.DB().SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		l.Errorf("cannot connect to db : %v", err)
		panic(err)
	}

	// When not in production, set Gin to "release" mode
	if env != "production" {
		db.LogMode(true)
	}
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}
