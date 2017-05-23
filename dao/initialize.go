package dao

import (
	"database/sql"
	"photos/config"
)

// Db ddd
var Db *sql.DB

// Initialize Initializes
func Initialize() {
	linkString := config.Conf.Usename + ":" + config.Conf.Password + "@/" + config.Conf.Databasename
	Db, _ = sql.Open(config.Conf.Database, linkString)
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(15)
}
