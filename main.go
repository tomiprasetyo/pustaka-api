package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", RootHandler)
	router.GET("/books/:id", BooksHandler)
	router.GET("/books/:id/:title", BooksHandler)
	router.GET("/query", QueryHandler)
	router.POST("/books", PostBooksHandler)

	router.Run()
}

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":   "Tomi Prasetyo",
		"status": "OK",
	})
}

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")

	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")

	ctx.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

type BookInput struct {
	Title        string
	Price        int
	SeriesNumber string `json:"series_number"`
}

func PostBooksHandler(ctx *gin.Context) {
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":         bookInput.Title,
		"price":         bookInput.Price,
		"series_number": bookInput.SeriesNumber,
	})
}
