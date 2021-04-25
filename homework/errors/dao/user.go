package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

/*
 * 查询用户列表
 */
func GetUserList(num int) {
	users := make(map[interface{}]interface{})

	rows, err := Db.Query("SELECT id, username, sex, mobile, addTime FROM tbUser LIMIT ?", num)
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

/*
 * 根据用户ID查询用户信息
 */
func GetUserInfo(id int) {
	var user User

	err := Db.QueryRow("SELECT id, username, sex, mobile, addTime FROM tbUser WHERE id = ?", id).Scan(&user.id, &user.username, &user.sex, &user.mobile, &user.addTime)
	if err != nil {
		if err == sql.ErrNoRows {
			// 空结果，并不是一个错误，不应该Wrap这个error抛给上层。
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println(user)
}