package connections

import (
	"elephant/config"
	"elephant/log"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql dialect
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //psql dialect
	_ "github.com/lib/pq"
)

var db *gorm.DB

var (
	l   = log.GetLogger()
	err error
)

// Auto Initiate db connection
func init() {

	cg := config.GetConfig()

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

		// enable db to handle GROUP BY
		_, err := db.DB().Exec(`SET GLOBAL sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''))`)

		if err != nil {
			l.Errorf("cannot replace @@sql_mode to 'ONLY_FULL_GROUP_BY' : %v", err)
		}
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
