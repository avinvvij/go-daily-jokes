package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/avinvvij/go-daily-jokes/controllers"
)

func AddAppSessionRoutes(routerGroup *gin.RouterGroup) {

	appSessionRouter := routerGroup.Group("/appSession")

	//creating a new app session
	appSessionRouter.POST("/" , controllers.NewAppSession )

}
