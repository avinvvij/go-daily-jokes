package controllers

import  (
  "io/ioutil"
  "encoding/json"

  "github.com/gin-gonic/gin"
  "github.com/avinvvij/go-daily-jokes/models"
  services "github.com/avinvvij/go-daily-jokes/Services"
  uuid "github.com/satori/go.uuid"
  "github.com/go-playground/validator/v10"
)

func NewAppSession(c *gin.Context){
  jsonData, err := ioutil.ReadAll(c.Request.Body )
  if err != nil{
    c.AbortWithStatusJSON(500, gin.H{
      "error": err.Error(),
      "message": "Server error occurred",
    })
  }else{
    var appSession models.AppSession
    _uuid := uuid.NewV4().String()
    appSession.UUID = _uuid
    json.Unmarshal(jsonData, &appSession)
    validationError := validator.New().Struct(appSession)
    if validationError != nil{
      //there are validation errors
      c.AbortWithStatusJSON(400 , gin.H{
        "message": "Invalid form data",
      })
    }else{
      appSession, err := services.NewAppSession(appSession)
      if err != nil{
        c.AbortWithStatusJSON(500, gin.H{
          "error" : err.Error(),
          "message": "Server error occurred",
        })
        return
      }
      c.JSON(201, appSession)
      return
    }
  }
}
