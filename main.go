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

	bookRepository := book.NewRepository(db)

	// books, err := bookRepository.FindAll()
	// book, err := bookRepository.FindByID(2)
	book := book.Book{
		Title:       "Enterprise Integration Patterns",
		Description: "This book offers an invaluable catalog of various pattern suggestions with real-world solutions that help you design effective messaging solutions for your enterprise.",
		Price:       95000,
		Rating:      4,
		Discount:    0,
	}

	newBook, err := bookRepository.Create(book)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success create new database", newBook)

	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
