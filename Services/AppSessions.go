package services

import (
	models "github.com/avinvvij/go-daily-jokes/models"
	"github.com/kamva/mgm/v3"
)

func NewAppSession(appSession models.AppSession)(models.AppSession , error){
    err := mgm.Coll(&appSession).Create(&appSession)
    return appSession, err
}
