package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
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
	fmt.Println("init")
	dbCfg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USERNAME, PASSWORD, HOST, PORT, DATABASE, CHARSET)
	fmt.Printf("dbCfg：%s \n", dbCfg)
	Database, err := sql.Open("mysql", dbCfg)
	if err != nil {
		fmt.Println(err)
		panic("数据源配置不正确: " + err.Error())
	}

	Db = Database

	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	if err = Db.Ping(); nil != err {
		fmt.Println(err)
		panic("数据库链接失败: " + err.Error())
	}
}

// 用户表结构体
type User struct {
	id int `db:"id"`
	username string  `db:"username"`
	sex int `db:"sex"`
	mobile string `db:"mobile"`
	addTime string `db:"addTime"`
}

func main() {
	//dbCfg := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USERNAME, PASSWORD, HOST, PORT, DATABASE, CHARSET)
	//fmt.Printf("dbCfg：%s \n", dbCfg)
	//Db, err := sql.Open("mysql", dbCfg)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	panic("数据源配置不正确: " + err.Error())
	//}
	//defer Db.Close()
	fmt.Printf("Db：%v \n", Db)

	//users := make([]User, 0)

	var (
		id int
		username string
	)

	rows, err := Db.Query("SELECT id, username FROM tbUser LIMIT ?", 10)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &username)
		if err != nil {
			fmt.Println(err)
		}
		log.Println(id, username)
	}

	//var user User
	//
	//for rows.Next(){
	//	rows.Scan(&user.id, &user.username, &user.sex, &user.mobile)
	//	users = append(users,user)
	//}
	//
	//fmt.Println(users)
}