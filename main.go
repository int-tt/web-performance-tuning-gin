package main

import (
	"database/sql"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MessageLine struct {
	Text     string
	CreateAt string
}

func main() {
	db, err := sql.Open("mysql", "demouser:demopass@tcp(127.0.0.1:3306)/groupwork?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := gin.Default()
	router.Static("/static/", "./static/")
	router.LoadHTMLGlob("view/*")

	router.GET("/exercise/part1", func(c *gin.Context) {
		var messageLine [1000]string
		for i := 0; i < 1000; i++ {
			messageLine[i] = "Sunrise" + strconv.Itoa(time.Now().Year()) + "　チューニングバトル！誰が栄冠の1位になるのか？0.001秒を削る熱いバトル！！！誰が？誰が？誰が？誰が栄冠の1位に！！！！！！！！！！！"
		}

		c.HTML(http.StatusOK, "part1.tpl", gin.H{
			"Message": messageLine,
		})
	})

	router.GET("/exercise/part2", func(c *gin.Context) {
		var messages string
		var follow string
		var follower string
		var name string
		var messagesLine []MessageLine
		id := rand.Intn(99999) + 1
		rows, _ := db.Query("select (select count(id) from messages where user_id = ?) messages,(select count(id) from  follows where user_id = ?) follow, (select count(id) from  follows where follow_user_id =?) follower,(SELECT name FROM  users WHERE id = ?) name", id, id, id, id)
		for rows.Next() {
			_ = rows.Scan(&messages, &follow, &follower, &name)
		}
		rows, _ = db.Query("select message,created_at from messages where id = ? order by created_at desc limit 10", id)
		for rows.Next() {
			var row MessageLine
			_ = rows.Scan(&row.Text, &row.CreateAt)
			messagesLine = append(messagesLine, row)
		}
		c.HTML(http.StatusOK, "part2_all.tpl", gin.H{
			"Message":      messages,
			"Follow":       follow,
			"Follower":     follower,
			"User":         name,
			"MessagesLine": messagesLine,
		})
	})
	//router.POST("/exercise/part3", Part3)
	//router.GET("/exercise/part4", Part4)
	//router.GET("/exercise/part5", Part5)

	//router.LoadHTMLGlob("view/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//router.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tpl", gin.H{
	//		"Website": "helloworld",
	//		"Email":   "Email",
	//	})
	//})

	router.Run(":8080")
}

//Part1
func Part1(c *gin.Context) {
	var messageLine [1000]string
	for i := 0; i < 1000; i++ {
		messageLine[i] = "Sunrise" + strconv.Itoa(time.Now().Year()) + "　チューニングバトル！誰が栄冠の1位になるのか？0.001秒を削る熱いバトル！！！誰が？誰が？誰が？誰が栄冠の1位に！！！！！！！！！！！"
	}

	c.HTML(http.StatusOK, "part1.tpl", gin.H{
		"Message": messageLine,
	})
}

//Part2
func Part2(c *gin.Context) {

}

//Part3
func Part3(c *gin.Context) {

}

//Part4
func Part4(c *gin.Context) {

}

//Part5
func Part5(c *gin.Context) {

}
