package model

import (
	"qiqiChat/util"
	"strings"
)

//执行数据迁移

// func migration() {
// 	// 自动迁移模式
// 	DB.AutoMigrate(&User{})
// }

const (
	tablename = "users"
	field1    = "id INT(10) PRIMARY KEY,"
	field2    = "created_at TIMESTAMP,"
	field3    = "deleted_at TIMESTAMP,"
	field4    = "user_name VARCHAR(255),"
	field5    = "password_digest VARCHAR(255),"
	field6    = "nick_name VARCHAR(255),"
	field7    = "status VARCHAR(255),"
	field8    = "avatar VARCHAR(1000)" //最后一个字段不要加逗号
)

// CreateUserTable 创建用户表
func CreateUserTable() {
	statement := strings.Join([]string{"CREATE TABLE IF NOT EXISTS ", tablename, "(", field1, field2, field3, field4, field5, field6, field7, field8, ")", ";"}, "")
	stmt, err := DB.Prepare(statement)
	if err != nil {
		util.Err.Fatalln("Failed to prepare sql statement: ", err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		util.Warn.Println("Failed to create user table: ", err.Error())
	} else {
		util.Inform.Println("Succeed to create user table")
	}
}
