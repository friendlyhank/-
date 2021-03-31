package common

import (
	"miaosha/common/db"
	"miaosha/common/kafka"
)

func init(){
	//初始化数据库
	db.Init()

	//初始化kafka
	kafka.Init()
}
