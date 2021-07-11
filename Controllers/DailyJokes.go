package controllers

import (
	"fmt"

	services "github.com/avinvvij/go-daily-jokes/Services"
	models "github.com/avinvvij/go-daily-jokes/models"
	"github.com/gin-gonic/gin"
)

func NewDailyJoke(c *gin.Context) {
	dailyJoke, err := services.NewDailyJoke(models.DailyJoke{
		JokeTitle:       "Hello Devs",
		JokeDescription: "Why to java developers wear specs? Because they can't C",
		JokeOwner:       "Avin",
	})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Server error occurred",
		})
	} else {
		c.JSON(201, dailyJoke)
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
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, dailyJoke)
	}
}

func DeleteJoke(c *gin.Context) {

}
