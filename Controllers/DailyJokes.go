package controllers

import (
	"encoding/json"
	"io/ioutil"

	services "github.com/avinvvij/go-daily-jokes/Services"
	"github.com/avinvvij/go-daily-jokes/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewDailyJoke(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Invalid request body",
		})
	}
	var dailyJoke models.DailyJoke
	json.Unmarshal(jsonData, &dailyJoke)
	validationError := validator.New().Struct(dailyJoke)
	if validationError != nil {
		errorResult := ""
		for _, e := range validationError.(validator.ValidationErrors) {
			errorResult += e.Error() + "; "
		}
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Invalid request body",
			"errors":  errorResult,
		})
	} else {
		createdDailyJoke, err := services.NewDailyJoke(dailyJoke)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": "Server error occurred",
			})
		} else {
			c.JSON(201, createdDailyJoke)
		}
	}
}

func GetAllDailyJokes(c *gin.Context) {
	dailyJokes, err := services.GetAllJokes()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Server error occurred",
		})
	} else {
		c.JSON(200, dailyJokes)
	}
}

func GetJokeById(c *gin.Context) {
	dailyJoke, err := services.GetJokeById(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, dailyJoke)
	}
}

func DeleteJoke(c *gin.Context) {

}
