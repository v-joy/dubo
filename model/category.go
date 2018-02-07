package model 

import (
	// "github.com/patrickmn/go-cache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Category struct{
	Id int
	Name string
	Subname string
	Image string
} 

func init() {
	orm.RegisterModel(new(Category))
	orm.RegisterDataBase("default", "mysql", "root:joy139139@tcp(127.0.0.1:3306)/dubo?charset=utf8", 30)
}

func CategoryList() []Category {
	var categories []Category
	sql := "select id,name,subname,image from category"
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&categories) 

	return categories
}
