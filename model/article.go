package model 

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"log"
)

type Article struct {
	Id int
	Title string
	Image string
	Content string 
	Source string
	Read_count int
	post_time  int
}

func init() {
	orm.RegisterModel(new(Article))
	orm.RegisterDataBase("default", "mysql", "root:joy139139@tcp(127.0.0.1:3306)/dubo?charset=utf8", 30)
}

func ArticleItem(id int) Article {
	o := orm.NewOrm()
	article := Article{Id:1}
	_ = o.Read(&article)
	log.Println(article)
	return article
}

func ArticleList(pn,rn,categoryid,lastid int) []Article {
	var articles []Article;
	sql := "select id,title,image,content,source,read_count,post_time from article where categoryid=? and id < ? limit ?,?" 	
	o := orm.NewOrm()
	_, err := o.Raw(sql,categoryid,lastid,rn*(pn-1),rn).QueryRows(&articles);
	log.Println(err,categoryid,lastid,pn,rn)
	return articles 
}
