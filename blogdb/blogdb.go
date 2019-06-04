package blogdb

import (
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)
type Blog struct {
	gorm.Model
	Title string
	Text  string
}

func DbInit() {
	db, err := gorm.Open("sqlite3", "article.sqlite3")
	if err != nil {
		panic("データベース開けない（dbInit）")
	}
	db.AutoMigrate(&Blog{})
	defer db.Close()
}

func DbInsert(title string, text string) {
	db, err := gorm.Open("sqlite3", "article.sqlite3")
	if err != nil {
		panic("データベース開けない（dbInsert）")
	}
	db.Create(&Blog{Title: title, Text: text})
	defer db.Close()
}

func DbGetAll() []Blog {
	db, err := gorm.Open("sqlite3", "article.sqlite3")
	if err != nil {
		panic("データベース開けない（dbGetAll）")
	}
	var blogs []Blog
	db.Order("created_at desc").Find(&blogs)
	db.Close()
	return blogs
}

func DbGetOne(id int) Blog {
	db, err := gorm.Open("sqlite3", "article.sqlite3")
	if err != nil {
		panic("データベース開けない(dbGetOne())")
	}
	var blog Blog
	db.First(&blog, id)
	db.Close()
	return blog
}

func DbUpdate(id int, title string, text string) {
	db, err := gorm.Open("sqlite3", "article.sqlite3")
	if err != nil {
		panic("データベース開けない（dbUpdate）")
	}
	var blog Blog
	db.First(&blog, id)
	blog.Title = title
	blog.Text = text
	db.Save(&blog)
	db.Close()
}

func DbDelete(id int) {
	db, err := gorm.Open("sqlite3", "article.sqlite3")
	if err != nil {
		panic("データベース開けない（dbDelete）")
	}
	var blog Blog
	db.First(&blog, id)
	db.Delete(&blog)
	db.Close()
}
