package main

import (
	"github.com/gin-gonic/gin"
	"dubo/controller"
)

func main() {
	prefix := ""

	app := gin.Default()
	app.StaticFile("/favicon.ico", "./static/favicon.ico")
	app.Use(gin.Recovery())
	// app.Use(gin.Logger())

	// app.Static("/static", "./static")
	// app.LoadHTMLGlob("templates/*")

	app.GET("/", controller.Index)
	app.GET(prefix+"/categories", controller.CategoryList)
	app.GET(prefix+"/articles", controller.ArticleList)
	app.GET(prefix+"/article", controller.ArticleItem)

	app.Run("123.56.88.120:8885")
}
