package models

import "github.com/kamva/mgm/v3"

type AppSession struct {
	mgm.DefaultModel `bson:",inline"`
	UUID             string `json:"uuid" validate:"required,min=16"`
	FcmToken         string `json:"fcmToken" validate:"required"`
}
