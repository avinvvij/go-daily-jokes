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

func DeleteJoke(id string) (bool, error){
	dailyJoke, err := GetJokeById(id)
	if err != nil{
		return false, err
	}
	deleteErr := mgm.Coll(&dailyJoke).Delete(&dailyJoke)
	if deleteErr != nil{
		return false, deleteErr
	}
	return true, nil
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

func PublishJoke(id string) (bool, error) {
	dailyJoke, err := GetJokeById(id)
	if err != nil {
		return false, err
	}
	dailyJoke.Published = true
	updateErr := mgm.Coll(&dailyJoke).Update(&dailyJoke)
	if updateErr != nil {
		return false, updateErr
	}
	return true, nil
}

func UnPublishJoke(id string) (bool ,error){
	dailyJoke, err := GetJokeById(id)
	if(err != nil){
		return false, err
	}
	dailyJoke.Published = false;
	updateErr := mgm.Coll(&dailyJoke).Update(&dailyJoke)
	if updateErr != nil {
		return false, updateErr
	}
	return true, nil
}
