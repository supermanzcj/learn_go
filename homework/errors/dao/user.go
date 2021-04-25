package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

const (
	USERNAME = "root"
	PASSWORD = "Superzc123456"
	HOST      = "127.0.0.1"
	PORT      = "3306"
	DATABASE  = "dbTest"
	CHARSET   = "utf8"
)

/*
 * DB初始化
 */
func init() {
	dbCfg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USERNAME, PASSWORD, HOST, PORT, DATABASE, CHARSET)
	fmt.Printf("dbCfg：%s \n", dbCfg)
	Database, err := sql.Open("mysql", dbCfg)
	if err != nil {
		fmt.Println(err)
		panic("数据源配置不正确: " + err.Error())
	}

	err = Database.Ping()
	if err != nil {
		fmt.Println(err)
		panic("数据源连接失败: " + err.Error())
	}

	Db = Database
}

// 用户表结构体
type User struct {
	id int `db:"id"`
	username string  `db:"username"`
	sex int `db:"sex"`
	mobile string `db:"mobile"`
	addTime string `db:"addTime"`
}

func GetUserList() {
	users := make(map[interface{}]interface{})

	rows, err := Db.Query("SELECT id, username, sex, mobile, addTime FROM tbUser LIMIT ?", 10)
	if err != nil {
		fmt.Println(err)
	}

	var user User

	for rows.Next(){
		rows.Scan(&user.id, &user.username, &user.sex, &user.mobile, &user.addTime)
		users[user.id] = user
	}

	fmt.Println(users)
}