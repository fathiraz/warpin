package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SqliteHandler() *gorm.DB {
	// set db name from environment
	var dbName = fmt.Sprintf("%s.db", GetEnv().Sqlite.Database.Name)

	// open our db. this db will keep use without open anymore
	// as a db connection pooling
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}

	// set db log
	db.LogMode(true)

	return db
}
