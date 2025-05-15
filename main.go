package main

import (
	"cqupt_hub/dao/mysql"
	"cqupt_hub/routes"
	"cqupt_hub/setting"
)

func main() {
	setting.Init()
	mysql.Init(setting.Conf.MysqlConfig)
	routes.SetRouter()
}
