package config

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Connection string
	Host       string
	Port       int
	DBName     string
	Username   string
	Password   string
	Charset    string
	Loc        string
	DB         *gorm.DB
}

const (
	MYSQL      = "mysql"
	POSTGRESQL = "postgres"
	SQLITE3    = "sqlite3"
	SQLSERVER  = "mssql"
)

func (d *Database) Configure(app *iris.Application) {
	var connectionString string

	switch d.Connection {
	case MYSQL:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s", d.Username, d.Password, d.Host, d.Port, d.DBName, d.Charset, d.Loc)
	case POSTGRESQL:
		connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", d.Host, d.Port, d.Username, d.DBName, d.Password)
	case SQLITE3:
		connectionString = "/tmp/gorm.db"
	case SQLSERVER:
		connectionString = fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", d.Username, d.Password, d.Host, d.Port, d.DBName)
	default:
		app.Logger().Warn("Database connection does not support")
		return
	}

	db, err := gorm.Open(d.Connection, connectionString)
	if err != nil {
		app.Logger().Warn("Database failed to connect")
		app.Logger().Warn(err)
		return
	}
	d.DB = db
}