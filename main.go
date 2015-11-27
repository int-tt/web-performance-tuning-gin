package main

import (
	"database/sql"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	//"github.com/sonots/go-sql_metrics"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MessageLine struct {
	Text     string
	CreateAt string
}

func main() {
	db, errr := sql.Open("mysql", "demouser:demopass@unix(/var/lib/mysql/mysql.sock)/groupwork?charset=utf8")
	if errr != nil {
		panic(errr.Error())
	}
	db.SetMaxIdleConns(2048)
	//db := sql_metrics.WrapDB("demouser:demopass@unix(/var/lib/mysql/mysql.sock)/groupwork?charset=utf8", _db)
	//sql_metrics.Verbose = true
	//sql_metrics.Print(1)
	defer db.Close()
	router := gin.New()
	router.Use(gin.Recovery())
	//router := gin.Default()
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
		if rows, err := db.Query("select (select count(id) from messages where user_id = ?) messages,(select count(id) from  follows where user_id = ?) follow, (select count(id) from  follows where follow_user_id =?) follower,(SELECT name FROM  users WHERE id = ?) name", id, id, id, id); err != nil {
		} else {
			for rows.Next() {
				if err = rows.Scan(&messages, &follow, &follower, &name); err != nil {
				}
			}
		}
		if rows, err := db.Query("select message,created_at from messages where user_id = ? order by created_at desc limit 10", id); err != nil {
		} else {
			for rows.Next() {
				var row MessageLine
				if err = rows.Scan(&row.Text, &row.CreateAt); err != nil {
				}
				messagesLine = append(messagesLine, row)
			}
		}
		c.HTML(http.StatusOK, "part2_all.tpl", gin.H{
			"Message":      messages,
			"Follow":       follow,
			"Follower":     follower,
			"User":         name,
			"MessagesLine": messagesLine,
		})
	})

	router.POST("/exercise/part3", func(c *gin.Context) {
		id := rand.Intn(1000006) + 1
		title := c.PostForm("title")
		message := c.PostForm("message")
		//fmt.Println(title,message)
		_, _ = db.Query("insert into messages values(null,?,?,?,now(),now())", id, title, message+"by "+strconv.Itoa(id))
		c.Redirect(http.StatusMovedPermanently, "/exercise/part1")
	})

	router.GET("/exercise/part4", func(c *gin.Context) {
		string := "チューニングバトル"
		var messagesLine []MessageLine
		if rows, err := db.Query("select message,created_at from messages where title = ? order by created_at desc limit 10", string); err != nil {
		} else {
			for rows.Next() {
				var row MessageLine
				if err = rows.Scan(&row.Text, &row.CreateAt); err != nil {
				}
				messagesLine = append(messagesLine, row)
			}
		}
		c.HTML(http.StatusOK, "part2_all.tpl", gin.H{
			"MessagesLine": messagesLine,
		})
	})

	router.GET("/exercise/part5", func(c *gin.Context) {
		c.HTML(http.StatusOK, "part5.tpl", gin.H{})
	})
	router.RunUnix("main.sock")
	//router.Run("127.0.0.1:8888")
}
