package main

import (
	"fmt"

	routes "github.com/avinvvij/go-daily-jokes/routes"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/api/v1")
	routes.AddDailyJokesRoutes(v1)
	err := mgm.SetDefaultConfig(nil, "daily_joke", options.Client().ApplyURI("mongodb+srv://AVINVIJ:U6WgZRJ7AcgrqHV7@cluster0.kpcci.mongodb.net/daily_joke"))
	if err != nil {
		fmt.Println("Unable to connect to database")
	} else {
		fmt.Println("Starting the server at port 8080")
		server.Run()
	}
}
