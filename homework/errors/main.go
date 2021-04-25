package main

import (
	"errors/dao"
)

func main() {
	//查询用户列表
	dao.GetUserList(10)

	//查询用户信息
	dao.GetUserInfo(2)
}