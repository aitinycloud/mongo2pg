package pg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)



var db *gorm.DB

var insightcfgdb *gorm.DB

var mysqlinsightdb *gorm.DB

var mysqlinsightcfgdb *gorm.DB

//Setup models Setup init.
func Setup() {
	var err error
	dbType := setting.PgType
	dbName := setting.PgDBName
	user := setting.PgUser
	password := setting.PgPassword
	host := setting.PgHost
	connStr := "postgres://%s:%s@%s/%s?sslmode=disable"
	connStr = fmt.Sprintf(connStr, user, password, host, dbName)
	log.Println(connStr)
	db, err = gorm.Open(dbType, connStr)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(4)
	db.DB().SetMaxOpenConns(16)
	db.DB().SetConnMaxLifetime(0)
}

func GetDB() *sql.DB {
	return db.DB()
}

func GetGormDB() *gorm.DB {
	return db
}

//CloseDB Models Close.
func CloseDB() {
	db.Close()
}
