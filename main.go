package main

import (
	"fmt"
	"os"

	routes "github.com/avinvvij/go-daily-jokes/routes"
	utils "github.com/avinvvij/go-daily-jokes/utils"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	utils.SetupEnvironment()
	server := gin.Default()
	v1 := server.Group("/api/v1")
	routes.AddDailyJokesRoutes(v1)
	err := mgm.SetDefaultConfig(nil, "daily_joke", options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		fmt.Println("Unable to connect to database")
	} else {
		fmt.Println("Starting the server at port 3001")
		server.Run(":3001")
	}
}
