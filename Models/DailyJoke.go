package models

import "github.com/kamva/mgm/v3"

type DailyJoke struct {
	mgm.DefaultModel `bson:",inline"`
	JokeTitle        string `json:"JokeTitle" validate:"required,min=5"`
	JokeDescription  string `json:"JokeDescription" validate:"required,min=5"`
	JokeImage        string `json:"JokeImage"`
	JokeOwner        string `json:"JokeOwner" validate:"required,min=3"`
	Published        bool `json:"Published" validate:"default:false"`
}
