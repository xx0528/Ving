package service

import (
	"database/sql"
	"os"
	"time"

	"server/model"

	"github.com/88250/gulu"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

var db *gorm.DB
var useSQLite bool

// ConnectDB connects to the database.
func ConnectDB() {
	var err error
	useSQLite = false
	if "" != model.Conf.SQLite {
		db, err = gorm.Open("sqlite3", model.Conf.SQLite)
		useSQLite = true
	} else if "" != model.Conf.MySQL {
		db, err = gorm.Open("mysql", model.Conf.MySQL)
	} else {
		logger.Fatal("please specify database")
	}
	if nil != err {
		logger.Fatalf("opens database failed: " + err.Error())
	}
	if useSQLite {
		logger.Debug("used [SQLite] as underlying database")
	} else {
		logger.Debug("used [MySQL] as underlying database")
	}

	if err = db.AutoMigrate(model.Models...).Error; nil != err {
		logger.Fatal("auto migrate tables failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.LogMode(model.Conf.ShowSQL)
}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}

// DBStat returns database statistics.
func DBStat() sql.DBStats {
	return db.DB().Stats()
}

// Database returns the underlying database name.
func Database() string {
	if useSQLite {
		return "SQLite"
	}

	return "MySQL"
}
