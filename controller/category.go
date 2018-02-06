package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"dubo/model"
)

var categories [6]map[string]string

func init() {
	for i := 0; i < 6; i++ {
	        cate := make(map[string]string)
		cate["id"] = strconv.Itoa(i)
		cate["name"] = "分类" + strconv.Itoa(i)
		cate["readnum"] = "666"
		cate["image"] = "http://neisou.baidu.com/images/headportrait/chenlei17/0_112.jpg"
		categories[i] = cate
	}
	_ = categories;
}

func CategoryList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": model.CategotyList(), "errno": 0})
}
