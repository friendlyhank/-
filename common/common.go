package common

import (
	"miaosha/common/db"
)

func init(){
	//初始化数据库
	db.Init()

	//初始化kafka
	//kafka.Init()
}
