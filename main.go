package main

import (
	"fmt"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&book.Book{})

	// book := book.Book{}
	// book.Title = "Refactoring: Improving the Design of Existing Code"
	// book.Description = "is a book written by Martin Fowler. This book improves your legacy code`s design to enhance software maintainability and make current code easier to understand."
	// book.Price = 110000
	// book.Discount = 10000
	// book.Rating = 5

	// err = db.Create(&book).Error
	// if err != nil {
	// 	panic(err)
	// }

	var book []book.Book

	err = db.Debug().Find(&book).Error
	if err != nil {
		panic(err)
	}

	for _, b := range book {
		fmt.Println("Title : ", b.Title)
		fmt.Printf("Book Object %v\n", b)
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
