package routes

import (
	"github.com/pmohanj/golang-jwt-project/middleware"

	controller "github.com/pmohanj/golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.Authenticate())
	incommingRoutes.GET("/users", controller.GetUsers())
	incommingRoutes.GET("/users/:user_id", controller.GetUser())

}
