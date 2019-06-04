package main

import (
	"github.com/gin-gonic/gin"
	"strconv"

	"gin-blog/blogdb"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	blogdb.DbInit()

	// index
	r.GET("/", func(ctx *gin.Context) {
		blogs := blogdb.DbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"blogs": blogs,
		})
	})
	
	//create
	r.POST("/new", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		text := ctx.PostForm("text")
		blogdb.DbInsert(title, text)
		ctx.Redirect(302, "/")
	})

	//Detail
	r.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		blog := blogdb.DbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"blog": blog})
	})

	//Update
	r.GET("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		blog := blogdb.DbGetOne(id)
		ctx.HTML(200, "update.html", gin.H{"blog": blog})
	})

	r.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		title := ctx.PostForm("title")
		text := ctx.PostForm("text")
		blogdb.DbUpdate(id, title, text)
		ctx.Redirect(302, "/")
	})

	//排除確認
	r.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		blog := blogdb.DbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"blog": blog})
	})

	//Delete
	r.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		blogdb.DbDelete(id)
		ctx.Redirect(302,"/")
	})

	r.Run()
}