package routes

import (
	controllers "github.com/avinvvij/go-daily-jokes/controllers"
	middlewares "github.com/avinvvij/go-daily-jokes/middlewares"
	"github.com/gin-gonic/gin"
)

func AddDailyJokesRoutes(routeGroup *gin.RouterGroup) {
	dailyJokes := routeGroup.Group("/dailyjokes")

	dailyJokes.GET("/token", controllers.GenerateAuthToken)

	dailyJokes.Use(middlewares.VerifyJWT)

	dailyJokes.GET("/:id", controllers.GetJokeById)

	dailyJokes.PUT("/publish/:id", controllers.PublishJoke )

	dailyJokes.PUT("/unpublish/:id" , controllers.UnPublishJoke)

	dailyJokes.DELETE("/delete/:id" , controllers.DeleteJoke )

	dailyJokes.POST("/", controllers.NewDailyJoke)

	dailyJokes.GET("/", controllers.GetAllDailyJokes)

}
