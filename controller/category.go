package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"dubo/model"
)

func CategoryList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": model.CategoryList(), "errno": 0})
}
