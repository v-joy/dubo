package controller 

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"dubo/model"
	"strconv"
)


func init() {
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world", "errno": 0})
}

func ArticleItem(c *gin.Context) {
	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	article := model.ArticleItem(id);
	c.JSON(http.StatusOK, gin.H{"data": article, "errno": 0})
}

func ArticleList(c *gin.Context) {
	lastidstr := c.Query("previd")
	lastid, _ := strconv.Atoi(lastidstr)
	pnstr := c.Query("pn")
	pn, _ := strconv.Atoi(pnstr)
	rnstr := c.Query("rn")
	rn, _ := strconv.Atoi(rnstr)
	categoryidstr := c.Query("categoryid")
	categoryid, _ := strconv.Atoi(categoryidstr)
	articles := model.ArticleList(pn,rn,categoryid,lastid)
	c.JSON(http.StatusOK, gin.H{"data": articles, "errno": 0})
}

