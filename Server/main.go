package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"mrunalini.com/controller"
)

var err error

func init() {
	controller.Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/BlogSystem")

	if err != nil {
		log.Fatal("DB Connection Failed to Open")
	}
}

func main() {

	r := gin.Default()
	r.POST("/articles", controller.CreateArticle)
	r.GET("/articles/:id", controller.GetArticleById)
	r.GET("/articles", controller.GetAllArticles)
	r.Run(":8080")
}
