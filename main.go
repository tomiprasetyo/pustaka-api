package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", RootHandler)
	v1.GET("/books/:id", BooksHandler)
	v1.GET("/books/:id/:title", BooksHandler)
	v1.GET("/query", QueryHandler)
	v1.POST("/books", PostBooksHandler)

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
	Title        string `json:"title" binding:"required"`
	Price        int    `json:"price" binding:"required,number"`
	SeriesNumber string `json:"series_number"`
}

func PostBooksHandler(ctx *gin.Context) {
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":         bookInput.Title,
		"price":         bookInput.Price,
		"series_number": bookInput.SeriesNumber,
	})
}
