package routes

import (
	controllers "github.com/avinvvij/go-daily-jokes/controllers"
	"github.com/gin-gonic/gin"
)

func AddDailyJokesRoutes(routeGroup *gin.RouterGroup) {
	dailyJokes := routeGroup.Group("/dailyjokes")

	dailyJokes.GET("/token", controllers.GenerateAuthToken)

	dailyJokes.GET("/:id", controllers.GetJokeById)

	dailyJokes.PUT("/publish/:id", controllers.PublishJoke )

	dailyJokes.PUT("/unpublish/:id" , controllers.UnPublishJoke)

	dailyJokes.POST("/", controllers.NewDailyJoke)

	dailyJokes.GET("/", controllers.GetAllDailyJokes)

}
