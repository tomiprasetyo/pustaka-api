package main

import (
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

	// membuat data

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

	// mengambil data

	// var book []book.Book
	var book book.Book

	// db.First(&book) // mengambil data urutan pertama
	// db.Last(&book) // mengambil data urutan terakhir
	// db.First(&book, 2) // mengambil data berdasarkan primary key
	// db.Find(&book) // mengambil semua data objek slice
	// db.Where("title = ?", "Clean Code").Find() // mengambil data berdasarkan kondisi tertentu
	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		panic(err)
	}

	// mengambil data slice
	// for _, b := range book {
	// 	fmt.Println("Title : ", b.Title)
	// 	fmt.Printf("Book Object %v\n", b)
	// }

	// mengupdate data

	book.Title = "Clean Code 2022 Edition"
	err = db.Save(&book).Error
	if err != nil {
		panic(err)
	}

	// menghapus data

	// db.Delete(&book).Error
	//

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
