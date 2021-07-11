package services

import (
	models "github.com/avinvvij/go-daily-jokes/models"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func NewDailyJoke(joke models.DailyJoke) (models.DailyJoke, error) {
	err := mgm.Coll(&joke).Create(&joke)
	if err != nil {
		return models.DailyJoke{}, err
	}
	return joke, nil
}

func GetAllJokes() ([]models.DailyJoke, error) {
	dailyJokes := []models.DailyJoke{}
	err := mgm.Coll(&models.DailyJoke{}).SimpleFind(&dailyJokes, bson.M{})
	if err != nil {
		return []models.DailyJoke{}, err
	} else {
		return dailyJokes, nil
	}
}

func GetJokeById(id string) (models.DailyJoke, error) {
	dailyJoke := models.DailyJoke{}
	err := mgm.Coll(&models.DailyJoke{}).FindByID(id, &dailyJoke)
	if err != nil {
		return models.DailyJoke{}, err
	} else {
		return dailyJoke, nil
	}
}
