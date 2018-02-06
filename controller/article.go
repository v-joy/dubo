package controller 

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

var articleList [6]map[string]string
var articleItem map[string]string

func init() {
	for i:=0; i<6; i++ {
		article := make(map[string]string)
		article["id"] = strconv.Itoa(i)
		article["image"] = "http://neisou.baidu.com/images/headportrait/chenlei17/0_112.jpg"
		article["title"] = "题目" + strconv.Itoa(i)
		article["url"] = "https://mbd.baidu.com/article/" + strconv.Itoa(i)
		articleList[i] = article 
	}
	
	articleItem = make(map[string]string)
	articleItem["id"] = "111"
	articleItem["title"] = "标题"
	articleItem["content"] = "内容"
	articleItem["pageView"] = "100"
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world", "errno": 0})
}

func ArticleItem(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": articleItem, "errno": 0})
}

func ArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": articleList, "errno": 0})
}

/*
func Summery(c *gin.Context) {
	db := getDb();

	// 个人得分
	rows, err := db.Query("SELECT cid, sum(score) as score FROM vote where vid=" + strconv.Itoa(vid) + " and cid not in ('正方','反方') group by cid")
	userResult := map[string]int{}
	var score int
	var cid string
	for rows.Next() {
		rows.Columns()
		err = rows.Scan(&cid, &score)
		checkErr(err)
		userResult[cid] = score
	}
	rows.Close()

	// 团队得分
	rows, err = db.Query("SELECT count(id) as total FROM vote where vid=" + strconv.Itoa(vid) + " and cid  in ('正方','反方')")
	var total int
	for rows.Next() {
		rows.Columns()
		err = rows.Scan(&total)
		checkErr(err)
	}
	rows.Close()
	rows, err = db.Query("SELECT cid, sum(score) as score FROM vote where vid=" + strconv.Itoa(vid) + " and cid in ('正方','反方') group by cid")
	teamResult := map[string]int{}
	for rows.Next() {
		rows.Columns()
		err = rows.Scan(&cid, &score)
		checkErr(err)
		teamResult[cid] = score * 10 / total
	}
	rows.Close()
	rows, err = db.Query("SELECT cid,uid, sum(score) as score FROM vote where vid=" + strconv.Itoa(vid) + " and cid in ('正方','反方') group by cid,uid order by cid")
	teamDetail := []map[string]string{}
	var uid,cscore string
	for rows.Next() {
		rows.Columns()
		err = rows.Scan(&cid, &uid, &cscore)
		checkErr(err)
		teamDetail = append(teamDetail, map[string]string{
			"cid":cid,
			"uid":uid,
			"cscore":cscore,
		})
	}
	rows.Close()
	c.JSON(http.StatusOK, gin.H{"userResult": userResult, "teamResult": teamResult,"teamDetail":teamDetail, "success": true})
}
*/
