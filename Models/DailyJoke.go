package models

import "github.com/kamva/mgm/v3"

type DailyJoke struct {
	mgm.DefaultModel `bson:",inline"`
	JokeTitle        string `json:"JokeTitle"`
	JokeDescription  string `json:"JokeDescription"`
	JokeImage        string `json:"JokeImage"`
	JokeOwner        string `json:"JokeOwner"`
}
