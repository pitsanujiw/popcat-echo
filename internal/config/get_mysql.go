package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetMySQL() *sql.DB {
	mysql, err := sql.Open("mysql", Get(EnvMysqlDSN))
	if err != nil {
		panic(err)
	}
	mysql.SetConnMaxLifetime(time.Minute * 3)
	mysql.SetMaxOpenConns(10)
	mysql.SetMaxIdleConns(10)
	return mysql
}
