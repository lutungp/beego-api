package database

import "github.com/beego/beego/v2/client/orm"

func Init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@/goecho?charset=utf8")
}
