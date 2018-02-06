package model 

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id int
	Name string
	Readnum int
	image string
}

func init() {
	orm.RegisterModel(new(Article))
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8", 30)
}

func ArticleItem(id int) {
}

func ArticleList() {
}
