package model

import (
	"database/sql"
	"qiqiChat/util"
	"time"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// DB 数据库连接结构体
var DB *sql.DB

// InitDB 数据库初始化
func InitDB(dsn string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		util.Err.Fatalln("Failed to open database: ", err.Error())
	}

	//db.SetConnMaxLifetime(1000)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Second * 3000)
	db.SetMaxOpenConns(100)
	//fmt.Println("====>")
	if err = db.Ping(); err != nil {
		util.Err.Fatalln("Failed to connect database: ", err.Error())
	}
	DB = db
	//CreateUserTable()
}
